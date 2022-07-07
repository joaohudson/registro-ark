package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateRegion(database *sql.DB, region models.CategoryRegistryRequest) *util.ApiError {
	err := db.CreateRegion(database, region)
	if err != nil {
		fmt.Println("Erro ao criar novo tipo de alimentação no banco: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func ListAllRegions(database *sql.DB) ([]models.Category, *util.ApiError) {
	regions, err := db.ListAllRegions(database)
	if err != nil {
		fmt.Println("Erro ao listar regiões: ", err)
		return []models.Category{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return regions, nil
}
