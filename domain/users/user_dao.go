package users

import (
	"fmt"
	"github.com/BobbyGifford/go_bookstore_users_api/utils/errors"
)

// Data access object
// Work with database here

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Ger () *errors.RestErr {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	currentUser := usersDB[user.Id]
	if currentUser!= nil {
		if user.Email == currentUser.Email {
			return errors.NewBadRequestError(fmt.Sprintf("user %s already exists", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	// Save if user is not in DB
	usersDB[user.Id] = user
	return nil
}
