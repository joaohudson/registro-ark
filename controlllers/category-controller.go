package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
	"github.com/joaohudson/registro-ark/util"
)

type CategoryController struct {
	loginService      *service.LoginService
	locomotionService *service.LocomotionService
	regionService     *service.RegionService
	foodService       *service.FoodService
}

func NewCategoryController(
	loginService *service.LoginService,
	locomotionService *service.LocomotionService,
	regionService *service.RegionService,
	foodService *service.FoodService) *CategoryController {

	return &CategoryController{
		loginService:      loginService,
		locomotionService: locomotionService,
		regionService:     regionService,
		foodService:       foodService,
	}
}

func (c *CategoryController) CreateLocomotion(response http.ResponseWriter, request *http.Request) {

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

func (c *CategoryController) DeleteLocomotion(response http.ResponseWriter, request *http.Request) {

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

func (c *CategoryController) CreateRegion(response http.ResponseWriter, request *http.Request) {

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

func (c *CategoryController) DeleteRegion(response http.ResponseWriter, request *http.Request) {

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

func (c *CategoryController) CreateFood(response http.ResponseWriter, request *http.Request) {

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

func (c *CategoryController) DeleteFood(response http.ResponseWriter, request *http.Request) {

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

func (c *CategoryController) DinoCategories(response http.ResponseWriter, request *http.Request) {

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
