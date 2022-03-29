package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type UserService interface {
	Save(models.User) models.User
	Update(models.User)
	Delete(models.User)
	FindAll() []models.User
	FindByUserId(string) models.User
}

type userService struct {
	userRepository repository.UserRepository
}

func New(repo repository.UserRepository) UserService {
	return &userService{
		userRepository: repo,
	} 
}

func (service *userService) Save(user models.User) models.User {
	service.userRepository.Save(user)
	return user
}

func (service *userService) FindAll() []models.User {
	return service.userRepository.FindAll()
}

func (service *userService) Update(user models.User) {
	service.userRepository.Update(user)
}

func (service *userService) Delete(user models.User) {
	service.userRepository.Delete(user)
}

func (service *userService) FindByUserId(user_id string) models.User {
	return service.userRepository.FindByUserId(user_id)
}