package utilities

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pressly/goose"
)

func GetDBHandle() (db *sql.DB, err error) {
	mysqlDsnSchema := ""
	mysqlDsn := ""
	schema := os.Getenv("DATABASE_NAME")
	dbHost := os.Getenv("DATABASE_HOST")
	dbUser := os.Getenv("DATABASE_USER")
	dbPort := os.Getenv("DATABASE_PORT")
	dbPass := os.Getenv("DATABASE_PASS")

	if dbUser != "" && dbPass != "" && dbHost != "" && dbPort != "" {
		mysqlDsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUser, dbPass, dbHost, dbPort)
		mysqlDsnSchema = fmt.Sprintf("%s%s", mysqlDsn, schema)
		flags := os.Getenv("MYSQL_FLAGS")
		if flags != "" {
			mysqlDsn = fmt.Sprintf("%s?%s", mysqlDsn, flags)
			mysqlDsnSchema = fmt.Sprintf("%s?%s", mysqlDsnSchema, flags)
		}
	} else {
		return nil, errors.New("no database configuration was provided")
	}

	db, err = sql.Open("mysql", mysqlDsnSchema)
	if err := db.Ping(); err != nil {
		fmt.Printf("Failed to connect to schema, attempting to create it")
		db, err = sql.Open("mysql", mysqlDsn)
		if err != nil {
			return nil, err
		}
		_, err := db.Exec(fmt.Sprintf("create database if not exists %s", schema))
		if err != nil {
			return nil, err
		}
		db.Close()
		db, err = sql.Open("mysql", mysqlDsnSchema)
		if err != nil {
			return nil, err
		}
	}

	if err := db.Ping(); err != nil {
		// log.Fatalf("Failed to ping database: %s", err.Error())
		return nil, err
	}

	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(2 * time.Minute)

	err = doMigrations("./migrations", db)
	if err != nil {
		log.Fatalf("Failed to run migrations: %s", err.Error())
	}

	return db, err
}

func doMigrations(dir string, db *sql.DB) error {
	goose.SetDialect("mysql")
	return goose.Run("up", db, dir)
}
