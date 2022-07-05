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

	existsDino, err := db.ExistsDinoByName(database, dino.Name)
	if err != nil {
		fmt.Println("Erro ao verificar dino por nome: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if existsDino {
		return util.ThrowApiError("Já existe um dino com esse nome!", http.StatusPreconditionFailed)
	}

	existsLocomotion, err2 := db.ExistsLocomotionById(database, dino.LocomotionId)
	if err2 != nil {
		fmt.Printf("Erro ao verificar locomoção para id %v: %v\n", dino.LocomotionId, err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsLocomotion {
		return util.ThrowApiError("Informe uma locomoção válida!", http.StatusBadRequest)
	}

	existsRegion, err3 := db.ExistsRegionById(database, dino.RegionId)
	if err3 != nil {
		fmt.Printf("Erro ao verificar região para id %v: %v\n", dino.RegionId, err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsRegion {
		return util.ThrowApiError("Informe uma região válida!", http.StatusBadRequest)
	}

	existsFood, err4 := db.ExistsFoodById(database, dino.FoodId)
	if err4 != nil {
		fmt.Printf("Erro ao verificar alimentação para id %v: %v\n", dino.RegionId, err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if !existsFood {
		return util.ThrowApiError("Informe um tipo de alimentação válida!", http.StatusBadRequest)
	}

	err5 := db.CreateDino(database, dino)

	if err5 != nil {
		fmt.Println("Erro ao inserir dados no banco: ", err5)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusBadRequest)
	}

	return nil
}

func DeleteDino(database *sql.DB, id uint64) *util.ApiError {
	exists, err := db.ExistsDinoById(database, id)
	if err != nil {
		fmt.Println("Erro ao verificar dino por id: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	if !exists {
		return util.ThrowApiError("Esse dino não existe!", http.StatusNotFound)
	}

	err2 := db.DeleteDino(database, id)
	if err2 != nil {
		fmt.Println("Erro ao deletar dino: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}

func FindDinoById(database *sql.DB, id uint64) (models.Dino, *util.ApiError) {
	dino, err := db.FindDinoById(database, id)
	if err != nil {
		fmt.Println("Erro ao recuperar dino por id: ", err.Error())
		return models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if dino == nil {
		return models.Dino{}, util.ThrowApiError("Dino não encontrado!", http.StatusBadRequest)
	}

	return *dino, nil
}

func FindDinoByFilter(database *sql.DB, filter models.DinoFilter) ([]models.Dino, *util.ApiError) {
	dinos, err := db.FindDinoByFilter(database, filter)
	if err != nil {
		fmt.Println("Erro ao buscar dinos por filtro: ", err)
		return []models.Dino{}, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return dinos, nil
}
