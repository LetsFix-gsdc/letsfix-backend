package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type TypeService interface {
	SaveType(models.Type) models.Type
	//Update(models.User)
	//Delete(models.User)
	FindAllTypes() []models.Type
	FindByTypeId(uint) models.Type
}

type typeService struct {
	typeRepository repository.TypeRepository
}

func NewTypeService(repo repository.TypeRepository) TypeService {
	return &typeService{
		typeRepository: repo,
	} 
}

func (service *typeService) SaveType(t models.Type) models.Type {
	service.typeRepository.SaveType(t)
	return t
}

func (service *typeService) FindAllTypes() []models.Type {
	return service.typeRepository.FindAllTypes()
}

func (service *typeService) FindByTypeId(type_id uint) models.Type {
	return service.typeRepository.FindByTypeId(type_id)
}