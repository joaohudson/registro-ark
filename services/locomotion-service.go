package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateLocomotion(database *sql.DB, locomotion models.CategoryRegistryRequest) *util.ApiError {
	err := db.CreateLocomotion(database, locomotion)
	if err != nil {
		fmt.Println("Erro ao criar locomoção no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func ListAllLocomotions(database *sql.DB) ([]models.Category, *util.ApiError) {
	locomotion, err := db.ListAllLocomotions(database)
	if err != nil {
		fmt.Println("Erro ao listar locomoções: ", err)
		return []models.Category{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return locomotion, nil
}
