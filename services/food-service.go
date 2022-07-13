package service

import (
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

type FoodService struct {
	foodRepo *db.FoodRepository
	dinoRepo *db.DinoRepository
}

func NewFoodService(foodRepo *db.FoodRepository, dinoRepo *db.DinoRepository) *FoodService {
	return &FoodService{foodRepo: foodRepo, dinoRepo: dinoRepo}
}

func (f *FoodService) CreateFood(category models.CategoryRegistryRequest) *util.ApiError {
	err := f.foodRepo.CreateFood(category)
	if err != nil {
		fmt.Println("Erro ao criar novo tipo de alimentação no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func (f *FoodService) ListAllFoods() ([]models.Category, *util.ApiError) {
	foods, err := f.foodRepo.ListAllFoods()
	if err != nil {
		fmt.Println("Erro ao listar tipos de alimentação: ", err)
		return []models.Category{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return foods, nil
}

func (f *FoodService) DeleteFood(id uint64) *util.ApiError {
	existsFood, err := f.foodRepo.ExistsFoodById(id)
	if err != nil {
		fmt.Println("Erro ao buscar alimentação por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsFood {
		return util.ThrowApiError("Esse tipo de alimentação não existe!", http.StatusNotFound)
	}

	dinosWithFood, err2 := f.dinoRepo.FindDinoByFilter(models.DinoFilter{
		FoodId: id,
	})
	if err2 != nil {
		fmt.Println("erro ao analisar uso de alimentação: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if len(dinosWithFood) > 0 {
		return util.ThrowApiError("Existem dinos usando este tipo de alimentação, sendo assim, ele não pode ser deletado!", http.StatusPreconditionFailed)
	}

	err3 := f.foodRepo.DeleteFood(id)
	if err3 != nil {
		fmt.Println("Erro ao deletar alimentação: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
