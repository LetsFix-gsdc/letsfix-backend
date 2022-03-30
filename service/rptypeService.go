package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type RptypeService interface {
	SaveRptype(models.Repair_Type) models.Repair_Type
	FindAllRptypes() []models.Repair_Type
	FindRptypeById(uint) models.Repair_Type
	FindRptypeByRepairer(string) []models.Repair_Type
	FindRptypeByType(string) []models.Repair_Type
	FindRptypeByBrand(string) []models.Repair_Type
	FindRptypeByComponent(string) []models.Repair_Type
}

type rptypeService struct {
	rptypeRepository repository.RptypeRepository
}

func NewRptypeService(repo repository.RptypeRepository) RptypeService {
	return &rptypeService{
		rptypeRepository: repo,
	} 
}

func (service *rptypeService) SaveRptype(r models.Repair_Type) models.Repair_Type {
	service.rptypeRepository.SaveRptype(r)
	return r
}

func (service *rptypeService) FindAllRptypes() []models.Repair_Type {
	return service.rptypeRepository.FindAllRptypes()
}

func (service *rptypeService) FindRptypeById(r_id uint) models.Repair_Type {
	return service.rptypeRepository.FindRptypeById(r_id)
}

func (service *rptypeService) FindRptypeByRepairer(r string) []models.Repair_Type {
	return service.rptypeRepository.FindRptypeByRepairer(r)
}

func (service *rptypeService) FindRptypeByType(t string) []models.Repair_Type {
	return service.rptypeRepository.FindRptypeByType(t)
}

func (service *rptypeService) FindRptypeByBrand(b string) []models.Repair_Type {
	return service.rptypeRepository.FindRptypeByBrand(b)
}

func (service *rptypeService) FindRptypeByComponent(c string) []models.Repair_Type {
	return service.rptypeRepository.FindRptypeByComponent(c)
}