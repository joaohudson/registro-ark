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
