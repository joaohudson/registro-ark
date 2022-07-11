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

func DeleteLocomotion(database *sql.DB, id uint64) *util.ApiError {
	existsLocomotion, err := db.ExistsLocomotionById(database, id)
	if err != nil {
		fmt.Println("Erro ao buscar locomoção por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsLocomotion {
		return util.ThrowApiError("Essa locomoção não existe!", http.StatusNotFound)
	}

	dinosWithLocomotion, err2 := db.FindDinoByFilter(database, models.DinoFilter{
		LocomotionId: id,
	})
	if err2 != nil {
		fmt.Println("Erro ao analisar uso de locomoção: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if len(dinosWithLocomotion) > 0 {
		return util.ThrowApiError("Existem dinos usando esta locomoção, sendo assim, ela não pode ser deletada!", http.StatusPreconditionFailed)
	}

	err3 := db.DeleteLocomotion(database, id)
	if err3 != nil {
		fmt.Println("Erro ao deletar locomoção: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
