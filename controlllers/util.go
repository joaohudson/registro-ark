package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joaohudson/registro-ark/util"
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
