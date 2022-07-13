package controller

import (
	"encoding/json"
	"net/http"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
)

type AdmController struct {
	admService *service.AdmService
}

func NewAdmController(admService *service.AdmService) *AdmController {
	return &AdmController{admService: admService}
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

	err2 := a.admService.CreateAdm(adm)
	if err2 != nil {
		sendError(response, err2)
		return
	}
}
