package service

import (
	"gsdc/letsfix/models"
	"gsdc/letsfix/repository"

	//"google.golang.org/grpc/channelz/service"
)

type BrandService interface {
	SaveBrand(models.Brand) models.Brand
	//Update(models.User)
	//Delete(models.User)
	FindAllBrands() []models.Brand
	FindByBrandId(uint) models.Brand
}

type brandService struct {
	brandRepository repository.BrandRepository
}

func NewBrandService(repo repository.BrandRepository) BrandService {
	return &brandService{
		brandRepository: repo,
	} 
}

func (service *brandService) SaveBrand(b models.Brand) models.Brand {
	service.brandRepository.SaveBrand(b)
	return b
}

func (service *brandService) FindAllBrands() []models.Brand {
	return service.brandRepository.FindAllBrands()
}

func (service *brandService) FindByBrandId(brand_id uint) models.Brand {
	return service.brandRepository.FindByBrandId(brand_id)
}