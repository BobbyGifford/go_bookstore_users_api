package users

import (
	"fmt"
	"github.com/BobbyGifford/go_bookstore_users_api/utils/errors"
	"github.com/bobbygifford/go_bookstore_users_api/datasources/mysql/users_db"
	"github.com/bobbygifford/go_bookstore_users_api/utils/date"
)

// Data access object
// Database logic

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get () *errors.RestErr {
	if err := users_db.Client.Ping(); err != nil {
		panic(err)
	}

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

	user.DateCreated = date.GetNowString()

	// Save if user is not in DB
	usersDB[user.Id] = user
	return nil
}
