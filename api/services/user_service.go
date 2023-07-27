package services

import (
	"errors"
	"go_server/api/helpers"
	"go_server/api/models"
	"go_server/api/repositories"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		userRepo: repositories.NewUserRepository(db),
	}
}

func (s *UserService) CreateUser(user *models.User) helpers.APIError {
	existingUser, err := s.userRepo.FindByEmail(strings.ToLower(user.Email))
	if err != nil {
		return helpers.ErrorResponse(err)
	}

	if existingUser != nil {
		return helpers.ErrorResponse(errors.New("E-mail já cadastrado"))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return helpers.ErrorResponse(err)
	}
	user.Password = string(hashedPassword)

	err = s.userRepo.Create(user)
	if err != nil {
		return helpers.ErrorResponse(err)
	}

	return helpers.APIError{Error: false, Message: ""}
}

func (s *UserService) UpdateUserName(id uint, name string) helpers.APIError {
	existingUser, err := s.userRepo.FindByID(id)
	if err != nil {
		return helpers.ErrorResponse(err)
	}

	if existingUser == nil {
		return helpers.ErrorResponse(errors.New("Usuário não encontrado"))
	}

	name = strings.ToLower(name)

	nameTaken, err := s.userRepo.IsNameTakenByAnotherUser(id, name)
	if err != nil {
		return helpers.ErrorResponse(err)
	}

	if nameTaken {
		return helpers.ErrorResponse(errors.New("Nome já foi cadastrado"))
	}

	err = s.userRepo.UpdateName(id, name)
	if err != nil {
		return helpers.ErrorResponse(err)
	}

	return helpers.APIError{Error: false, Message: ""}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}
