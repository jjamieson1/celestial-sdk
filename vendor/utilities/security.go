package utilities

import (
	"github.com/jjamieson1/eden-sdk/models"
	b "golang.org/x/crypto/bcrypt"
)

func HashAndSalt(password []byte) ([]byte, error) {

	hash, err := b.GenerateFromPassword(password, b.DefaultCost)
	if err != nil {
		return nil, err
	}
	return hash, err
}

func ComparePasswords(storedPassword, password []byte) error {
	err := b.CompareHashAndPassword(storedPassword, password)
	if err != nil {
		return err
	}
	return err
}

func IsAdmin(u models.User) bool {
	for _, v := range u.Roles {
		if v.RoleId == "1" {
			return true
		}
	}
	return false
}
