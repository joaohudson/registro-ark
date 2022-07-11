package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateFood(database *sql.DB, category models.CategoryRegistryRequest) *util.ApiError {
	err := db.CreateFood(database, category)
	if err != nil {
		fmt.Println("Erro ao criar novo tipo de alimentação no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func ListAllFoods(database *sql.DB) ([]models.Category, *util.ApiError) {
	foods, err := db.ListAllFoods(database)
	if err != nil {
		fmt.Println("Erro ao listar tipos de alimentação: ", err)
		return []models.Category{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return foods, nil
}

func DeleteFood(database *sql.DB, id uint64) *util.ApiError {
	existsFood, err := db.ExistsFoodById(database, id)
	if err != nil {
		fmt.Println("Erro ao buscar alimentação por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsFood {
		return util.ThrowApiError("Esse tipo de alimentação não existe!", http.StatusNotFound)
	}

	dinosWithFood, err2 := db.FindDinoByFilter(database, models.DinoFilter{
		FoodId: id,
	})
	if err2 != nil {
		fmt.Println("erro ao analisar uso de alimentação: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if len(dinosWithFood) > 0 {
		return util.ThrowApiError("Existem dinos usando este tipo de alimentação, sendo assim, ele não pode ser deletado!", http.StatusPreconditionFailed)
	}

	err3 := db.DeleteFood(database, id)
	if err3 != nil {
		fmt.Println("Erro ao deletar alimentação: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
