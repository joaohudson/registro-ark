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
	dinoService       *service.DinoService
	locomotionService *service.LocomotionService
	regionService     *service.RegionService
	foodService       *service.FoodService
	loginService      *service.LoginService
}

func NewDinoController(
	dinoService *service.DinoService,
	locomotionService *service.LocomotionService,
	regionService *service.RegionService,
	foodService *service.FoodService,
	loginService *service.LoginService) *DinoController {

	return &DinoController{
		dinoService:       dinoService,
		locomotionService: locomotionService,
		regionService:     regionService,
		foodService:       foodService,
		loginService:      loginService,
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

func (c *DinoController) CreateLocomotion(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, c.loginService, PermissionManagerCategory)
	if err != nil {
		sendError(response, err)
		return
	}

	decoder := json.NewDecoder(request.Body)

	var locomotion models.CategoryRegistryRequest
	err2 := decoder.Decode(&locomotion)
	if err2 != nil {
		fmt.Println("Erro ao fazer parse da locomoção: ", err2)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Locomoção inválida!"))
		return
	}

	err3 := c.locomotionService.CreateLocomotion(locomotion)
	if err3 != nil {
		sendError(response, err3)
		return
	}
}

func (c *DinoController) DeleteLocomotion(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, c.loginService, PermissionManagerCategory)
	if err != nil {
		sendError(response, err)
		return
	}

	id, err2 := parseQueryParameterUint64(request, "id")
	if err2 != nil {
		fmt.Println("Erro no parse do id:", err2)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
	}

	err3 := c.locomotionService.DeleteLocomotion(id)
	if err3 != nil {
		sendError(response, err3)
	}
}

func (c *DinoController) CreateRegion(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, c.loginService, PermissionManagerCategory)
	if err != nil {
		sendError(response, err)
		return
	}

	decoder := json.NewDecoder(request.Body)

	var region models.CategoryRegistryRequest
	err2 := decoder.Decode(&region)
	if err2 != nil {
		fmt.Println("Erro ao fazer parse da região: ", err2)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Região inválida!"))
		return
	}

	err3 := c.regionService.CreateRegion(region)
	if err3 != nil {
		sendError(response, err3)
		return
	}
}

func (c *DinoController) DeleteRegion(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, c.loginService, PermissionManagerCategory)
	if err != nil {
		sendError(response, err)
		return
	}

	id, err2 := parseQueryParameterUint64(request, "id")
	if err2 != nil {
		fmt.Println("Erro no parse do id:", err2)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
	}

	err3 := c.regionService.DeleteRegion(id)
	if err3 != nil {
		sendError(response, err3)
	}
}

func (c *DinoController) CreateFood(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, c.loginService, PermissionManagerCategory)
	if err != nil {
		sendError(response, err)
		return
	}

	decoder := json.NewDecoder(request.Body)

	var food models.CategoryRegistryRequest
	err2 := decoder.Decode(&food)
	if err2 != nil {
		fmt.Println("Erro ao fazer parse do tipo de alimentação: ", err2)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Tipo de Alimentação inválido!"))
		return
	}

	err3 := c.foodService.CreateFood(food)
	if err3 != nil {
		sendError(response, err3)
		return
	}
}

func (c *DinoController) DeleteFood(response http.ResponseWriter, request *http.Request) {

	_, err := authenticate(request, c.loginService, PermissionManagerCategory)
	if err != nil {
		sendError(response, err)
		return
	}

	id, err2 := parseQueryParameterUint64(request, "id")
	if err2 != nil {
		fmt.Println("Erro no parse do id:", err2)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
	}

	err3 := c.foodService.DeleteFood(id)
	if err3 != nil {
		sendError(response, err3)
	}
}

func (c *DinoController) DinoCategories(response http.ResponseWriter, request *http.Request) {

	regions, err := c.regionService.ListAllRegions()
	if err != nil {
		sendError(response, err)
		return
	}

	locomotions, err2 := c.locomotionService.ListAllLocomotions()
	if err2 != nil {
		sendError(response, err2)
		return
	}

	foods, err3 := c.foodService.ListAllFoods()
	if err3 != nil {
		sendError(response, err3)
		return
	}

	result := models.DinoCategoryResponse{
		Regions:     regions,
		Locomotions: locomotions,
		Foods:       foods,
	}

	sendJson(response, result)
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

func (c *DinoController) DeleteDino(response http.ResponseWriter, request *http.Request) {

	idAdm, err := authenticate(request, c.loginService, PermissionManagerCategory)
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
