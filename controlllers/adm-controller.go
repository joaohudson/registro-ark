package controller

import (
	"encoding/json"
	"net/http"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
)

type AdmController struct {
	admService   *service.AdmService
	loginService *service.LoginService
}

func NewAdmController(admService *service.AdmService, loginService *service.LoginService) *AdmController {
	return &AdmController{admService: admService, loginService: loginService}
}

func (a *AdmController) CreateAdm(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, a.loginService, PermissionManagerAdm)
	if err != nil {
		sendError(response, err)
		return
	}

	decode := json.NewDecoder(request.Body)
	var adm models.AdmRegistryRequest
	err2 := decode.Decode(&adm)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Dados da conta inválidos!"))
		return
	}

	err3 := a.admService.CreateAdm(adm)
	if err3 != nil {
		sendError(response, err3)
		return
	}
}
