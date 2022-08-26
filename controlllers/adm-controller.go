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

func (a *AdmController) DeleteAdm(response http.ResponseWriter, request *http.Request) {
	_, apiErr := authenticate(request, a.loginService, PermissionManagerAdm)
	if apiErr != nil {
		sendError(response, apiErr)
		return
	}

	idAdm, err := parseQueryParameterUint64(request, "id")
	if err != nil {
		sendError(response, util.ThrowApiError("Administrador não encontrado!", http.StatusBadRequest))
		return
	}

	apiErr = a.admService.DeleteAdm(idAdm)
	if apiErr != nil {
		sendError(response, apiErr)
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

func (a *AdmController) PutAdmCredentials(response http.ResponseWriter, request *http.Request) {
	idAdm, apiErr := authenticate(request, a.loginService, PermissionManagerNone)
	if apiErr != nil {
		sendError(response, apiErr)
		return
	}

	var credentials models.AdmChangeCredentialsRequest
	decode := json.NewDecoder(request.Body)
	err := decode.Decode(&credentials)
	if err != nil {
		fmt.Println("Erro ao fazer parse da mudança de credenciais: ", err)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
		return
	}

	apiErr = a.admService.PutAdmCredentials(idAdm, credentials)
	if apiErr != nil {
		sendError(response, apiErr)
		return
	}

	a.loginService.Logout(idAdm)
	response.WriteHeader(http.StatusUnauthorized)
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
