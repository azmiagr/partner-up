package service

import (
	"intern-bcc/internal/repository"
	"intern-bcc/pkg/bcrypt"
	"intern-bcc/pkg/jwt"
	// "intern-bcc/model"
)

type Service struct {
	User IUserService
}

func NewService(r *repository.Repository, bcrypt bcrypt.Interface, jwtAuth jwt.Interface) *Service {
	return &Service{
		User: NewUserService(r.User, bcrypt, jwtAuth),
	}
}

// func (u *UserService) Login(param model.UserLogin) (model.UserLoginResponse, error) {

// }
