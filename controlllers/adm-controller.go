package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
	"github.com/joaohudson/registro-ark/util"
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

func (a *AdmController) PutAdmPermissions(response http.ResponseWriter, request *http.Request) {
	_, err := authenticate(request, a.loginService, PermissionManagerAdm)
	if err != nil {
		sendError(response, err)
		return
	}

	var permissions models.AdmChangePermissionsRequest
	decoder := json.NewDecoder(request.Body)
	err2 := decoder.Decode(&permissions)
	if err2 != nil {
		fmt.Println("Erro ao fazer parse das permissões: ", err2)
		sendError(response, util.ThrowApiError("Não foi possível alterar as permissões do administrador! Recarregue a página e tente novamente.", http.StatusBadRequest))
		return
	}

	err = a.admService.PutAdmPermissions(permissions)
	if err != nil {
		sendError(response, err)
		return
	}
}

func (a *AdmController) GetAdms(response http.ResponseWriter, request *http.Request) {
	_, err := authenticate(request, a.loginService, PermissionManagerAdm)
	if err != nil {
		sendError(response, err)
		return
	}

	adms, err := a.admService.GetAdms()
	if err != nil {
		sendError(response, err)
		return
	}

	sendJson(response, adms)
}

func (a *AdmController) GetAdm(response http.ResponseWriter, request *http.Request) {
	idAdm, err := authenticate(request, a.loginService, PermissionManagerNone)
	if err != nil {
		sendError(response, err)
		return
	}

	adm, err2 := a.admService.GetAdm(idAdm)
	if err2 != nil {
		sendError(response, err2)
		return
	}

	sendJson(response, adm)
}
