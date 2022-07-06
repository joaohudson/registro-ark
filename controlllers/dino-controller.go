package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
	service "github.com/joaohudson/registro-ark/services"
)

type DinoController struct {
	database *sql.DB
}

func NewDinoController(database *sql.DB) *DinoController {
	return &DinoController{database: database}
}

func (c *DinoController) FindDinoById(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		fmt.Println("Erro no parse do id: ", err)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	data, err2 := service.FindDinoById(c.database, id)
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

	dinos, searchErr := service.FindDinoByFilter(c.database, filter)
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

	err2 := db.CreateLocomotion(c.database, locomotion)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
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

	err2 := db.CreateRegion(c.database, region)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
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

	err2 := db.CreateFood(c.database, food)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
}

func (c *DinoController) DinoCategories(response http.ResponseWriter, request *http.Request) {

	regions, err := db.ListAllRegions(c.database)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	locomotions, err2 := db.ListAllLocomotions(c.database)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err2.Error()))
		return
	}

	foods, err3 := db.ListAllFoods(c.database)
	if err3 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err3.Error()))
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

	err2 := service.CreateDino(c.database, dino)
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

	err2 := service.DeleteDino(c.database, id)

	if err2 != nil {
		sendError(response, err2)
		return
	}
}
