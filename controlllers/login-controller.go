package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
	"github.com/joaohudson/registro-ark/util"
)

type LoginController struct {
	loginService *service.LoginService
}

func NewLoginController(loginService *service.LoginService) *LoginController {
	return &LoginController{loginService: loginService}
}

func (l *LoginController) Login(response http.ResponseWriter, request *http.Request) {
	decode := json.NewDecoder(request.Body)
	var credentials models.LoginRequest
	err := decode.Decode(&credentials)
	if err != nil {
		fmt.Println("Erro ao fazer parse das credenciais: ", err)
		sendError(response, util.ThrowApiError("Dados do login inválidos!", http.StatusBadRequest))
		return
	}

	token, err2 := l.loginService.Login(credentials)
	if err2 != nil {
		sendError(response, err2)
		return
	}

	response.Write([]byte(token))
}
