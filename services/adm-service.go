package service

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateAdm(database *sql.DB, adm models.AdmRegistryRequest) *util.ApiError {

	existsAdm, err := db.ExistsAdmByName(database, adm.Name)
	if err != nil {
		fmt.Println("Erro ao buscar adm por nome: ", err)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}
	if existsAdm {
		return util.ThrowApiError("Já existe um administrador com este nome!", http.StatusPreconditionFailed)
	}

	err2 := db.CreateAdm(database, adm)
	if err2 != nil {
		fmt.Println("Erro ao criar conta de adminstrador: ", err2)
		return util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return nil
}
