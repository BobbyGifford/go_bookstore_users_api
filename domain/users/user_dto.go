package users

import (
	"github.com/bobbygifford/go_bookstore_users_api/utils/errors"
	"strings"
)

// Data transfer object (Model)
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
}

// Validation method on User struct/type
func (user *User) Validate() *errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}

	return nil
}
