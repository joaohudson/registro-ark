package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type RegionService struct {
	regionRepo *db.RegionRepository
	dinoRepo   *db.DinoRepository
}

func NewRegionService(regionRepo *db.RegionRepository, dinoRepo *db.DinoRepository) *RegionService {
	return &RegionService{regionRepo: regionRepo, dinoRepo: dinoRepo}
}

func (r *RegionService) CreateRegion(region models.CategoryRegistryRequest) *util.ApiError {
	err := r.regionRepo.CreateRegion(region)
	if err != nil {
		fmt.Println("Erro ao criar novo tipo de alimentação no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func (r *RegionService) ListAllRegions() ([]models.Category, *util.ApiError) {
	regions, err := r.regionRepo.ListAllRegions()
	if err != nil {
		fmt.Println("Erro ao listar regiões: ", err)
		return []models.Category{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return regions, nil
}

func (r *RegionService) DeleteRegion(id uint64) *util.ApiError {
	existsRegion, err := r.regionRepo.ExistsRegionById(id)
	if err != nil {
		fmt.Println("Erro ao buscar região por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsRegion {
		return util.ThrowApiError("Essa região não existe!", http.StatusNotFound)
	}

	dinosWithRegion, err2 := r.dinoRepo.FindDinoByFilter(models.DinoFilter{
		RegionId: id,
	})
	if err2 != nil {
		fmt.Println("Erro ao analisar uso de região: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if len(dinosWithRegion) > 0 {
		return util.ThrowApiError("Existem dinos usando esta região, sendo assim, ela não pode ser deletada!", http.StatusPreconditionFailed)
	}

	err3 := r.regionRepo.DeleteRegion(id)
	if err3 != nil {
		fmt.Println("Erro ao deletar região: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
