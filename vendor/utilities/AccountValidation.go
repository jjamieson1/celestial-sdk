package utilities

import (
	"fmt"
	"net/mail"
	"regexp"

	"github.com/jjamieson1/eden-sdk/models"
)

func ValidateName(account models.User) error {
	if account.FirstName == "" || account.LastName == "" {
		return fmt.Errorf("firstname and lastname is required")
	}
	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return fmt.Errorf("email is not valid: %s", err.Error())
	}
	return nil
}

func ValidatePhone(number string) error {
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	isValid := re.MatchString(number)
	if !isValid {
		return fmt.Errorf("phone number is not valid")

	}
	return nil
}

func ValidateAddress(address models.Address) error {
	return nil
}

func ValidateRole(role models.Role) error {
	return nil
}

func ValidatePassword(password string) error {
	return nil
}
