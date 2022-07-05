package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateDino(database *sql.DB, dino models.DinoRegistryRequest) *util.ApiError {

	err := db.CreateDino(database, dino)

	if err != nil {
		return util.ThrowApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}

func DeleteDino(database *sql.DB, id uint64) *util.ApiError {
	err := db.DeleteDino(database, id)
	if err != nil {
		return util.ThrowApiError(err.Error(), http.StatusBadRequest)
	}

	return nil
}

func FindDinoById(database *sql.DB, id uint64) (models.Dino, *util.ApiError) {
	dino, err := db.FindDinoById(database, id)
	if err != nil {
		fmt.Println("Erro ao recuperar dino por id: ", err.Error())
		return models.Dino{}, util.ThrowApiError("Dino não encontrado!", http.StatusBadRequest)
	}

	return dino, nil
}

func FindDinoByFilter(database *sql.DB, filter models.DinoFilter) ([]models.Dino, *util.ApiError) {
	dinos, err := db.FindDinoByFilter(database, filter)
	if err != nil {
		return []models.Dino{}, util.ThrowApiError(err.Error(), http.StatusBadRequest)
	}

	return dinos, nil
}
