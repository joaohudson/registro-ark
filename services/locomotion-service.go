package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type LocomotionService struct {
	locomotionRepo *db.LocomotionRepository
	dinoRepo       *db.DinoRepository
}

func NewLocomotionService(locomotionRepo *db.LocomotionRepository, dinoRepo *db.DinoRepository) *LocomotionService {
	return &LocomotionService{locomotionRepo: locomotionRepo, dinoRepo: dinoRepo}
}

func (l *LocomotionService) CreateLocomotion(locomotion models.CategoryRegistryRequest) *util.ApiError {
	err := l.locomotionRepo.CreateLocomotion(locomotion)
	if err != nil {
		fmt.Println("Erro ao criar locomoção no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func (l *LocomotionService) ListAllLocomotions() ([]models.Category, *util.ApiError) {
	locomotion, err := l.locomotionRepo.ListAllLocomotions()
	if err != nil {
		fmt.Println("Erro ao listar locomoções: ", err)
		return []models.Category{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return locomotion, nil
}

func (l *LocomotionService) DeleteLocomotion(id uint64) *util.ApiError {
	existsLocomotion, err := l.locomotionRepo.ExistsLocomotionById(id)
	if err != nil {
		fmt.Println("Erro ao buscar locomoção por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsLocomotion {
		return util.ThrowApiError("Essa locomoção não existe!", http.StatusNotFound)
	}

	dinosWithLocomotion, err2 := l.dinoRepo.FindDinoByFilter(models.DinoFilter{
		LocomotionId: id,
	})
	if err2 != nil {
		fmt.Println("Erro ao analisar uso de locomoção: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if len(dinosWithLocomotion) > 0 {
		return util.ThrowApiError("Existem dinos usando esta locomoção, sendo assim, ela não pode ser deletada!", http.StatusPreconditionFailed)
	}

	err3 := l.locomotionRepo.DeleteLocomotion(id)
	if err3 != nil {
		fmt.Println("Erro ao deletar locomoção: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
