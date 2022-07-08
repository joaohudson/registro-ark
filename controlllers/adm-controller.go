package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
)

type AdmController struct {
	database *sql.DB
}

func NewAdmController(database *sql.DB) *AdmController {
	return &AdmController{database: database}
}

func (a *AdmController) CreateAdm(response http.ResponseWriter, request *http.Request) {
	decode := json.NewDecoder(request.Body)
	var adm models.AdmRegistryRequest
	err := decode.Decode(&adm)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Dados da conta inválidos!"))
		return
	}

	err2 := service.CreateAdm(a.database, adm)
	if err2 != nil {
		sendError(response, err2)
		return
	}
}
