package services

import (
	"log"

	"github.com/mashingan/smapping"
	"github.com/ragil000/go-restful.git/dto"
	"github.com/ragil000/go-restful.git/entities"
	"github.com/ragil000/go-restful.git/repositories"
)

// UserService is
type UserService interface {
	Update(user dto.UserUpdateDTO) entities.User
	Profile(userID string) entities.User
}

type userService struct {
	userRepository repositories.UserRepository
}

// NewUserService is
func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (service *userService) Update(user dto.UserUpdateDTO) entities.User {
	userToUpdate := entities.User{}
	err := smapping.FillStruct(&userToUpdate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v:", err)
	}
	updatedUser := service.userRepository.UpdateUser(userToUpdate)
	return updatedUser
}

func (service *userService) Profile(userID string) entities.User {
	return service.userRepository.ProfileUser(userID)
}
