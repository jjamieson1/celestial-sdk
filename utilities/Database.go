package utilities

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/pressly/goose"
	"github.com/revel/revel"
)

func InitDB(database string) (*sql.DB, error) {
	mysqlDsnSchema := ""
	mysqlDsn := ""
	schema := revel.Config.StringDefault("schema", database)
	dbHost := revel.Config.StringDefault("dbHost", "localhost")
	dbUser := revel.Config.StringDefault("dbUser", "root")
	dbPort := revel.Config.StringDefault("dbPort", "3306")
	dbPass := revel.Config.StringDefault("dbPass", "root")

	revel.AppLog.Infof("using database schema: %s", schema)

	if dbUser != "" && dbPass != "" && dbHost != "" && dbPort != "" {
		mysqlDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)
		mysqlDsnSchema = fmt.Sprintf("%s%s", mysqlDsn, schema)
		flags := os.Getenv("MYSQL_FLAGS")
		if flags != "" {
			mysqlDsn = fmt.Sprintf("%s?%s", mysqlDsn, flags)
			mysqlDsnSchema = fmt.Sprintf("%s?%s", mysqlDsnSchema, flags)
		}
	} else {
		revel.AppLog.Error("no database configuration was provided")
	}
	var err error
	DB, err := sql.Open("mysql", mysqlDsnSchema)
	if err != nil {
		revel.AppLog.Errorf("Failed to connect to database with error: %s", err.Error())
	}
	if err := DB.Ping(); err != nil {
		revel.AppLog.Error("Failed to connect to schema, attempting to create it")
		DB, err = sql.Open("mysql", mysqlDsn)
		if err != nil {
			revel.AppLog.Error(err.Error())
		}

		DB.Exec(fmt.Sprintf("create database if not exists %s", schema))
		DB.Close()

		DB, err = sql.Open("mysql", mysqlDsnSchema)
		if err != nil {
			revel.AppLog.Error(err.Error())
		}
		revel.AppLog.Info("DB Connected")
	}

	if err := DB.Ping(); err != nil {
		revel.AppLog.Errorf("unable to ping database with error: %s", err.Error())
	}

	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(2 * time.Minute)

	revel.AppLog.Infof("database connections in use: %v", DB.Stats().InUse)
	revel.AppLog.Infof("database open connections: %v", DB.Stats().OpenConnections)

	err = doMigrations("./migrations", DB)
	if err != nil {
		revel.AppLog.Errorf("Failed to run migrations: %s", err.Error())
	}
	return DB, nil
}

func doMigrations(dir string, db *sql.DB) error {
	goose.SetDialect("mysql")
	return goose.Run("up", db, dir)
}
