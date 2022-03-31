package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"
)

type OwnershipService interface {
	SaveOwnership(models.Ownership) models.Ownership
	//UpdateDevice(models.Device)
	//DeleteDevice(models.Device)
	FindAllOwnerships() []models.Ownership
	FindOwnershipByUserId(string) []models.Ownership
	FindDevicesByUserId(string) []models.Ownership
}

type ownershipService struct {
	ownershipRepository repository.OwnershipRepository
}

func NewOwnershipService(repo repository.OwnershipRepository) OwnershipService {
	return &ownershipService{
		ownershipRepository: repo,
	}
}

func (service *ownershipService) SaveOwnership(ownership models.Ownership) models.Ownership {
	service.ownershipRepository.SaveOwnership(ownership)
	return ownership
}

func (service *ownershipService) FindAllOwnerships() []models.Ownership {
	return service.ownershipRepository.FindAllOwnerships()
}

func (service *ownershipService) FindOwnershipByUserId(user_id string) []models.Ownership {
	return service.ownershipRepository.FindOwnershipByUserId(user_id)
}

func (service *ownershipService) FindDevicesByUserId(user_id string) []models.Ownership {
	return service.ownershipRepository.FindDevicesByUserId(user_id)
}
