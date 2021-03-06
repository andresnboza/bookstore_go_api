package services

import (
	"bookstore_users-api/domain/users"
	"bookstore_users-api/utils/date_utils"
	"bookstore_users-api/utils/errors"
	"bookstore_users-api/utils/crypto_utils"
)

var (
	UsersServices usersServiceInterface = &usersService{}
)

type usersService struct {

}

type usersServiceInterface interface {
	GetUser(user_id int64) (*users.User, *errors.RestErr)
	CreateUser(user users.User) (*users.User, *errors.RestErr)
	UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr)
	DeleteUser(userId int64) *errors.RestErr
	SearchUser(status string) (users.Users, *errors.RestErr)
}

func (s *usersService) GetUser(user_id int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: user_id}

	if err := result.Get(); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *usersService) CreateUser(user users.User) (*users.User, *errors.RestErr) {
	// Validating the fields of the user for creation
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowString()
	user.Status = users.StatusActive
	user.Password = crypto_utils.GetMd5(user.Password) 
	
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *usersService) UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestErr) {
	current, err := UsersServices.GetUser(user.Id)
	if err != nil {
		return nil, err
	}

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if isPartial {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}


func (s *usersService) DeleteUser(userId int64) *errors.RestErr {
	user := &users.User{Id: userId}
	return user.Delete()
}

func (s *usersService) SearchUser(status string) (users.Users, *errors.RestErr) {
	dao := &users.User{}
	return dao.FindByStatus(status)
}