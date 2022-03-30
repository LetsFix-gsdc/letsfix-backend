package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type RctypeService interface {
	SaveRctype(models.Recycler_Type) models.Recycler_Type
	//Update(models.User)
	//Delete(models.User)
	FindAllRctypes() []models.Recycler_Type
	FindRctypeById(uint) models.Recycler_Type
	FindRctypeByType(string) []models.Recycler_Type
}

type rctypeService struct {
	rctypeRepository repository.RctypeRepository
}

func NewRctypeService(repo repository.RctypeRepository) RctypeService {
	return &rctypeService{
		rctypeRepository: repo,
	} 
}

func (service *rctypeService) SaveRctype(r models.Recycler_Type) models.Recycler_Type {
	service.rctypeRepository.SaveRctype(r)
	return r
}

func (service *rctypeService) FindAllRctypes() []models.Recycler_Type {
	return service.rctypeRepository.FindAllRctypes()
}

func (service *rctypeService) FindRctypeById(r_id uint) models.Recycler_Type {
	return service.rctypeRepository.FindRctypeById(r_id)
}

func (service *rctypeService) FindRctypeByType(t string) []models.Recycler_Type {
	return service.rctypeRepository.FindRctypeByType(t)
}