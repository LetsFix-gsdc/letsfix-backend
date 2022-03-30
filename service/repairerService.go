package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type RepairerService interface {
	SaveRepairer(models.Repairer) models.Repairer
	//Update(models.User)
	//Delete(models.User)
	FindAllRepairers() []models.Repairer
	FindByRepairerId(uint) models.Repairer
}

type repairerService struct {
	repairerRepository repository.RepairerRepository
}

func NewRepairerService(repo repository.RepairerRepository) RepairerService {
	return &repairerService{
		repairerRepository: repo,
	} 
}

func (service *repairerService) SaveRepairer(r models.Repairer) models.Repairer {
	service.repairerRepository.SaveRepairer(r)
	return r
}

func (service *repairerService) FindAllRepairers() []models.Repairer {
	return service.repairerRepository.FindAllRepairers()
}

func (service *repairerService) FindByRepairerId(repairer_id uint) models.Repairer {
	return service.repairerRepository.FindByRepairerId(repairer_id)
}