package services

import (
	"github.com/BobbyGifford/go_bookstore_users_api/domain/users"
	"github.com/BobbyGifford/go_bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	return &user, nil
}
