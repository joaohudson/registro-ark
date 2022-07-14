package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	service "github.com/joaohudson/registro-ark/services"
	"github.com/joaohudson/registro-ark/util"
)

type PermissionType uint

const (
	PermissionManagerNone     PermissionType = 0
	PermissionManagerAdm                     = 1
	PermissionManagerCategory                = 2
	PermissionManagerDino                    = 3
)

func parseQueryParameterUint64(request *http.Request, parameterName string) (uint64, error) {
	valueStr := request.URL.Query().Get(parameterName)
	var value uint64
	var err error
	if valueStr == "" {
		value, err = 0, nil
	} else {
		value, err = strconv.ParseUint(valueStr, 10, 64)
	}

	return value, err
}

func sendError(response http.ResponseWriter, err *util.ApiError) {
	response.WriteHeader(err.StatusCode)
	response.Write([]byte(err.Message))
}

func sendJson(response http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(response)
	response.Header().Add("Content-Type", "application/json; charset=utf-8")
	err := encoder.Encode(data)

	if err != nil {
		fmt.Println("Erro no parse json: ", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(util.DefaultInternalServerError))
		return
	}
}

func authenticate(request *http.Request, loginService *service.LoginService, permission PermissionType) (uint64, *util.ApiError) {
	token := request.Header.Get("Authorization")
	id, err := loginService.GetIdByToken(token)
	if err != nil {
		return 0, err
	}

	var err2 *util.ApiError

	switch permission {
	case PermissionManagerNone:
		err2 = nil
	case PermissionManagerAdm:
		err2 = loginService.CheckPermissionManagerAdm(id)
	case PermissionManagerDino:
		err2 = loginService.CheckPermissionManagerDino(id)
	case PermissionManagerCategory:
		err2 = loginService.CheckPermissionManagerCategory(id)
	default:
		fmt.Println("Tipo de permissão não identificada na autenticação: ", permission)
		err2 = util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError)
	}

	return id, err2
}
