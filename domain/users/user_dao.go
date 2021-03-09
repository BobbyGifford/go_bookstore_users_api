package users

import (
	"fmt"
	"github.com/BobbyGifford/go_bookstore_users_api/utils/errors"
	"github.com/bobbygifford/go_bookstore_users_api/datasources/mysql/users_db"
	"github.com/bobbygifford/go_bookstore_users_api/utils/date_utils"
	"strings"
)

// Data access object
// Database logic

var (
	usersDB = make(map[int64]*User)
)

const (
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?);"
	indexUniqueEmail = "users.users_email_uindex"
)

func (user *User) Get() *errors.RestErr {
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
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, insertErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)
	if insertErr != nil {
		if strings.Contains(insertErr.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", insertErr.Error()))
	}

	// Same as above (Worse performance and no query validation)
	// result, err := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, idErr := insertResult.LastInsertId()
	if idErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", idErr.Error()))
	}

	user.Id = userId

	return nil
}
