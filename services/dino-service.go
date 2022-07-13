package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type DinoService struct {
	dinoRepo       *db.DinoRepository
	locomotionRepo *db.LocomotionRepository
	regionRepo     *db.RegionRepository
	foodRepo       *db.FoodRepository
}

func NewDinoService(
	dinoRepo *db.DinoRepository,
	locomotionRepo *db.LocomotionRepository,
	regionRepo *db.RegionRepository,
	foodRepo *db.FoodRepository) *DinoService {

	return &DinoService{
		dinoRepo:       dinoRepo,
		locomotionRepo: locomotionRepo,
		regionRepo:     regionRepo,
		foodRepo:       foodRepo}
}

func (s *DinoService) CreateDino(dino models.DinoRegistryRequest) *util.ApiError {

	nameLen := len(dino.Name)
	if nameLen > util.MaxDinoName {
		message := fmt.Sprintf("Nome do dino muito longo!\nO tamanho máximo permitido é de %v caracteres.", util.MaxDinoName)
		return util.ThrowApiError(message, http.StatusBadRequest)
	} else if nameLen == 0 {
		return util.ThrowApiError("Informe o nome do dino!", http.StatusBadRequest)
	}

	utilityLen := len(dino.Utility)
	if utilityLen > util.MaxDinoUtility {
		message := fmt.Sprintf("Utilidade do dino muito longa!\nO tamanho máximo permitido é de %v caracteres.", util.MaxDinoUtility)
		return util.ThrowApiError(message, http.StatusBadRequest)
	} else if utilityLen == 0 {
		return util.ThrowApiError("Informe a utilidade do dino!", http.StatusBadRequest)
	}

	if len(dino.Training) == 0 {
		return util.ThrowApiError("Informe a descrição do dino!", http.StatusBadRequest)
	}

	existsDino, err := s.dinoRepo.ExistsDinoByName(dino.Name)
	if err != nil {
		fmt.Println("Erro ao verificar dino por nome: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if existsDino {
		return util.ThrowApiError("Já existe um dino com esse nome!", http.StatusPreconditionFailed)
	}

	existsLocomotion, err2 := s.locomotionRepo.ExistsLocomotionById(dino.LocomotionId)
	if err2 != nil {
		fmt.Printf("Erro ao verificar locomoção para id %v: %v\n", dino.LocomotionId, err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsLocomotion {
		return util.ThrowApiError("Informe uma locomoção válida!", http.StatusBadRequest)
	}

	existsRegion, err3 := s.regionRepo.ExistsRegionById(dino.RegionId)
	if err3 != nil {
		fmt.Printf("Erro ao verificar região para id %v: %v\n", dino.RegionId, err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsRegion {
		return util.ThrowApiError("Informe uma região válida!", http.StatusBadRequest)
	}

	existsFood, err4 := s.foodRepo.ExistsFoodById(dino.FoodId)
	if err4 != nil {
		fmt.Printf("Erro ao verificar alimentação para id %v: %v\n", dino.RegionId, err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsFood {
		return util.ThrowApiError("Informe um tipo de alimentação válida!", http.StatusBadRequest)
	}

	err5 := s.dinoRepo.CreateDino(dino)

	if err5 != nil {
		fmt.Println("Erro ao inserir dados no banco: ", err5)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusBadRequest)
	}

	return nil
}

func (s *DinoService) DeleteDino(id uint64) *util.ApiError {
	exists, err := s.dinoRepo.ExistsDinoById(id)
	if err != nil {
		fmt.Println("Erro ao verificar dino por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if !exists {
		return util.ThrowApiError("Esse dino não existe!", http.StatusNotFound)
	}

	err2 := s.dinoRepo.DeleteDino(id)
	if err2 != nil {
		fmt.Println("Erro ao deletar dino: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func (s *DinoService) FindDinoById(id uint64) (models.Dino, *util.ApiError) {
	dino, err := s.dinoRepo.FindDinoById(id)
	if err != nil {
		fmt.Println("Erro ao recuperar dino por id: ", err.Error())
		return models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if dino == nil {
		return models.Dino{}, util.ThrowApiError("Dino não encontrado!", http.StatusBadRequest)
	}

	return *dino, nil
}

func (s *DinoService) FindDinoByFilter(filter models.DinoFilter) ([]models.Dino, *util.ApiError) {
	dinos, err := s.dinoRepo.FindDinoByFilter(filter)
	if err != nil {
		fmt.Println("Erro ao buscar dinos por filtro: ", err)
		return []models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return dinos, nil
}
