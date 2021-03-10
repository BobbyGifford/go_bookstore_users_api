package users

import (
	"fmt"
	"github.com/bobbygifford/go_bookstore_users_api/datasources/mysql/users_db"
	"github.com/bobbygifford/go_bookstore_users_api/utils/date_utils"
	"github.com/bobbygifford/go_bookstore_users_api/utils/errors"
	"strings"
)

// Data access object
// Database logic

const (
	queryInsertUser  = "INSERT INTO users(first_name, last_name, email, password, date_created) VALUES(?, ?, ?, ?, ?);"
	indexUniqueEmail = "users.users_email_uindex"
	queryGetUser     = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?;"
	queryUpdateUser  = "UPDATE users SET first_name=?, last_name=?, email=? WHERE id=?;"
)

func (user *User) Get() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryGetUser)
	if err != nil {
		//logger.Error("error when trying to prepare get user statement", err)
		//return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
		return errors.NewInternalServerError("error")

	}
	defer stmt.Close()

	result := stmt.QueryRow(user.Id)

	if getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated); getErr != nil {
		//logger.Error("error when trying to get user by id", getErr)
		//return rest_errors.NewInternalServerError("error when tying to get user", errors.New("database error"))
		return errors.NewInternalServerError("error")
	}
	return nil
}

func (user *User) Save() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}

	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	insertResult, insertErr := stmt.Exec(user.FirstName, user.LastName, user.Email, user.Password, user.DateCreated)
	if insertErr != nil {
		if strings.Contains(insertErr.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", insertErr.Error()))
	}

	// Same as above (Worse performance and no query validation)
	// insertResult, insertErr := users_db.Client.Exec(queryInsertUser, user.FirstName, user.LastName, user.Email, user.DateCreated)

	userId, idErr := insertResult.LastInsertId()
	if idErr != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user %s", idErr.Error()))
	}

	user.Id = userId

	return nil
}

func (user *User) Update() *errors.RestErr {
	stmt, err := users_db.Client.Prepare(queryUpdateUser)
	if err != nil {
		return errors.NewInternalServerError("error when tying to update user")
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.FirstName, user.LastName, user.Email, user.Id)
	if err != nil {
		return errors.NewInternalServerError("error when tying to update user")
	}
	return nil
}
