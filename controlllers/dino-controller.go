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
}

func NewDinoController(
	dinoService *service.DinoService,
	locomotionService *service.LocomotionService,
	regionService *service.RegionService,
	foodService *service.FoodService) *DinoController {

	return &DinoController{
		dinoService:       dinoService,
		locomotionService: locomotionService,
		regionService:     regionService,
		foodService:       foodService,
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

func (c *DinoController) FindDinoByFilter(response http.ResponseWriter, request *http.Request) {

	region, err := parseQueryParameterUint64(request, "region")
	if err != nil {
		fmt.Println("Erro no parse do parâmetro: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Região inválida!"))
		return
	}

	locomotion, err2 := parseQueryParameterUint64(request, "locomotion")
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Locomoção inválida!"))
		return
	}

	food, err3 := parseQueryParameterUint64(request, "food")
	if err3 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Tipo de alimentação inválida!"))
		return
	}

	name := request.URL.Query().Get("name")

	filter := models.DinoFilter{
		RegionId:     region,
		LocomotionId: locomotion,
		FoodId:       food,
		Name:         name,
	}

	dinos, searchErr := c.dinoService.FindDinoByFilter(filter)
	if searchErr != nil {
		sendError(response, searchErr)
		return
	}

	sendJson(response, dinos)
}

func (c *DinoController) CreateLocomotion(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var locomotion models.CategoryRegistryRequest
	err := decoder.Decode(&locomotion)
	if err != nil {
		fmt.Println("Erro ao fazer parse da locomoção: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Locomoção inválida!"))
		return
	}

	err2 := c.locomotionService.CreateLocomotion(locomotion)
	if err2 != nil {
		sendError(response, err2)
		return
	}
}

func (c *DinoController) DeleteLocomotion(response http.ResponseWriter, request *http.Request) {
	id, err := parseQueryParameterUint64(request, "id")
	if err != nil {
		fmt.Println("Erro no parse do id:", err)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
	}

	err2 := c.locomotionService.DeleteLocomotion(id)
	if err2 != nil {
		sendError(response, err2)
	}
}

func (c *DinoController) CreateRegion(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var region models.CategoryRegistryRequest
	err := decoder.Decode(&region)
	if err != nil {
		fmt.Println("Erro ao fazer parse da região: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Região inválida!"))
		return
	}

	err2 := c.regionService.CreateRegion(region)
	if err2 != nil {
		sendError(response, err2)
		return
	}
}

func (c *DinoController) DeleteRegion(response http.ResponseWriter, request *http.Request) {
	id, err := parseQueryParameterUint64(request, "id")
	if err != nil {
		fmt.Println("Erro no parse do id:", err)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
	}

	err2 := c.regionService.DeleteRegion(id)
	if err2 != nil {
		sendError(response, err2)
	}
}

func (c *DinoController) CreateFood(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var food models.CategoryRegistryRequest
	err := decoder.Decode(&food)
	if err != nil {
		fmt.Println("Erro ao fazer parse do tipo de alimentação: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Tipo de Alimentação inválido!"))
		return
	}

	err2 := c.foodService.CreateFood(food)
	if err2 != nil {
		sendError(response, err2)
		return
	}
}

func (c *DinoController) DeleteFood(response http.ResponseWriter, request *http.Request) {
	id, err := parseQueryParameterUint64(request, "id")
	if err != nil {
		fmt.Println("Erro no parse do id:", err)
		sendError(response, util.ThrowApiError(util.DefaultInternalServerError, http.StatusInternalServerError))
	}

	err2 := c.foodService.DeleteFood(id)
	if err2 != nil {
		sendError(response, err2)
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
	var dino models.DinoRegistryRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dino)
	if err != nil {
		fmt.Println("Erro na criação de dino: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Informações do dino inválidas!"))
		return
	}

	err2 := c.dinoService.CreateDino(dino)
	if err2 != nil {
		sendError(response, err2)
		return
	}
}

func (c *DinoController) DeleteDino(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		fmt.Println("Erro no parse do id: ", err)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	err2 := c.dinoService.DeleteDino(id)

	if err2 != nil {
		sendError(response, err2)
		return
	}
}
