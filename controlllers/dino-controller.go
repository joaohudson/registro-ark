package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
	"github.com/joaohudson/registro-ark/util"
)

type DinoController struct {
	dinoService  *service.DinoService
	loginService *service.LoginService
}

func NewDinoController(
	dinoService *service.DinoService,
	loginService *service.LoginService) *DinoController {

	return &DinoController{
		dinoService:  dinoService,
		loginService: loginService,
	}
}

func (c *DinoController) FindDinoById(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		fmt.Println("Erro no parse do id: ", err)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	data, err2 := c.dinoService.FindDinoById(id)
	if err2 != nil {
		sendError(response, err2)
		return
	}

	sendJson(response, data)
}

func parseDinoFilter(request *http.Request) (models.DinoFilter, *util.ApiError) {
	region, err := parseQueryParameterUint64(request, "region")
	if err != nil {
		fmt.Println("Erro no parse do parâmetro region: ", err)
		return models.DinoFilter{}, util.ThrowApiError("Região inválida!", http.StatusBadRequest)
	}

	locomotion, err2 := parseQueryParameterUint64(request, "locomotion")
	if err2 != nil {
		fmt.Println("Erro no parse do parâmetro locomotion: ", err2)
		return models.DinoFilter{}, util.ThrowApiError("Locomoção inválida!", http.StatusBadRequest)
	}

	food, err3 := parseQueryParameterUint64(request, "food")
	if err3 != nil {
		fmt.Println("Erro no parse do parâmetro food: ", err3)
		return models.DinoFilter{}, util.ThrowApiError("Tipo de alimentação inválido!", http.StatusBadRequest)
	}

	name := request.URL.Query().Get("name")

	return models.DinoFilter{
		RegionId:     region,
		LocomotionId: locomotion,
		FoodId:       food,
		Name:         name,
	}, nil
}

func (c *DinoController) FindDinoByFilter(response http.ResponseWriter, request *http.Request) {

	filter, err := parseDinoFilter(request)
	if err != nil {
		sendError(response, err)
		return
	}

	dinos, err2 := c.dinoService.FindDinoByFilter(filter)
	if err2 != nil {
		sendError(response, err2)
		return
	}

	sendJson(response, dinos)
}

func (c *DinoController) FindDinoByFilterForAdm(response http.ResponseWriter, request *http.Request) {

	idAdm, err := authenticate(request, c.loginService, PermissionManagerDino)
	if err != nil {
		sendError(response, err)
		return
	}

	filter, err := parseDinoFilter(request)
	if err != nil {
		sendError(response, err)
		return
	}

	dinos, err2 := c.dinoService.FindDinoByFilterForAdm(idAdm, filter)
	if err2 != nil {
		sendError(response, err2)
		return
	}

	sendJson(response, dinos)
}

func (c *DinoController) CreateDino(response http.ResponseWriter, request *http.Request) {

	idAdm, err := authenticate(request, c.loginService, PermissionManagerDino)
	if err != nil {
		sendError(response, err)
		return
	}

	var dino models.DinoRegistryRequest
	decoder := json.NewDecoder(request.Body)
	err2 := decoder.Decode(&dino)
	if err2 != nil {
		fmt.Println("Erro na criação de dino: ", err2)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Informações do dino inválidas!"))
		return
	}

	err3 := c.dinoService.CreateDino(idAdm, dino)
	if err3 != nil {
		sendError(response, err3)
		return
	}
}

func (c *DinoController) PutDino(response http.ResponseWriter, request *http.Request) {

	idAdm, err := authenticate(request, c.loginService, PermissionManagerDino)
	if err != nil {
		sendError(response, err)
		return
	}
	idDino, err2 := parseQueryParameterUint64(request, "id")
	if err2 != nil {
		fmt.Println("Erro no parse do id do dino: ", err2)
		sendError(response, util.ThrowApiError("Esse dino não existe!", http.StatusBadRequest))
		return
	}

	var dino models.DinoRegistryRequest
	decoder := json.NewDecoder(request.Body)
	err2 = decoder.Decode(&dino)
	if err2 != nil {
		fmt.Println("Erro na alteração de dino: ", err2)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Informações do dino inválidas!"))
		return
	}

	err3 := c.dinoService.PutDino(idDino, idAdm, dino)
	if err3 != nil {
		sendError(response, err3)
		return
	}
}

func (c *DinoController) DeleteDino(response http.ResponseWriter, request *http.Request) {

	idAdm, err := authenticate(request, c.loginService, PermissionManagerDino)
	if err != nil {
		sendError(response, err)
		return
	}

	id, err2 := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)
	if err2 != nil {
		fmt.Println("Erro no parse do id: ", err2)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	err3 := c.dinoService.DeleteDino(idAdm, id)

	if err3 != nil {
		sendError(response, err3)
		return
	}
}
