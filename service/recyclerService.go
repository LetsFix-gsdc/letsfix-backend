package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type RecyclerService interface {
	SaveRecycler(models.Recycler) models.Recycler
	//Update(models.User)
	//Delete(models.User)
	FindAllRecyclers() []models.Recycler
	FindByRecyclerId(uint) models.Recycler
}

type recyclerService struct {
	recyclerRepository repository.RecyclerRepository
}

func NewRecyclerService(repo repository.RecyclerRepository) RecyclerService {
	return &recyclerService{
		recyclerRepository: repo,
	} 
}

func (service *recyclerService) SaveRecycler(r models.Recycler) models.Recycler {
	service.recyclerRepository.SaveRecycler(r)
	return r
}

func (service *recyclerService) FindAllRecyclers() []models.Recycler {
	return service.recyclerRepository.FindAllRecyclers()
}

func (service *recyclerService) FindByRecyclerId(recycler_id uint) models.Recycler {
	return service.recyclerRepository.FindByRecyclerId(recycler_id)
}