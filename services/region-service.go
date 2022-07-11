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

func DeleteRegion(database *sql.DB, id uint64) *util.ApiError {
	existsRegion, err := db.ExistsRegionById(database, id)
	if err != nil {
		fmt.Println("Erro ao buscar região por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsRegion {
		return util.ThrowApiError("Essa região não existe!", http.StatusNotFound)
	}

	dinosWithRegion, err2 := db.FindDinoByFilter(database, models.DinoFilter{
		RegionId: id,
	})
	if err2 != nil {
		fmt.Println("Erro ao analisar uso de região: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if len(dinosWithRegion) > 0 {
		return util.ThrowApiError("Existem dinos usando esta região, sendo assim, ela não pode ser deletada!", http.StatusPreconditionFailed)
	}

	err3 := db.DeleteRegion(database, id)
	if err3 != nil {
		fmt.Println("Erro ao deletar região: ", err3)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
