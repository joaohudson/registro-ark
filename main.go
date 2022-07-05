package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
)

const DefaultInternalServerErrorMessage = "Erro interno do servidor, por favor tente mais tarde."

var database *sql.DB

func main() {

	database = db.CreatePostgresDatabase()
	defer database.Close()

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)

	router.Post("/api/dino", createDino)
	router.Delete("/api/dino", deleteDino)
	router.Get("/api/dino", dino)
	router.Get("/api/dinos", dinos)
	router.Get("/api/dino/categories", dinoCategories)
	router.Post("/api/dino/category/food", createFood)
	router.Post("/api/dino/category/locomotion", createLocomotion)
	router.Post("/api/dino/category/region", createRegion)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}

func createLocomotion(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var locomotion models.CategoryRegistryRequest
	err := decoder.Decode(&locomotion)
	if err != nil {
		fmt.Println("Erro ao fazer parse da locomoção: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Locomoção inválida!"))
		return
	}

	err2 := db.CreateLocomotion(database, locomotion)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
}

func createRegion(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var region models.CategoryRegistryRequest
	err := decoder.Decode(&region)
	if err != nil {
		fmt.Println("Erro ao fazer parse da região: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Região inválida!"))
		return
	}

	err2 := db.CreateRegion(database, region)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
}

func createFood(response http.ResponseWriter, request *http.Request) {

	decoder := json.NewDecoder(request.Body)

	var food models.CategoryRegistryRequest
	err := decoder.Decode(&food)
	if err != nil {
		fmt.Println("Erro ao fazer parse do tipo de alimentação: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Tipo de Alimentação inválido!"))
		return
	}

	err2 := db.CreateFood(database, food)
	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err.Error()))
		return
	}
}

func dinoCategories(response http.ResponseWriter, request *http.Request) {

	regions, err := db.ListAllRegions(database)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err.Error()))
		return
	}

	locomotions, err2 := db.ListAllLocomotions(database)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err2.Error()))
		return
	}

	foods, err3 := db.ListAllFoods(database)
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

func createDino(response http.ResponseWriter, request *http.Request) {
	var dino models.DinoRegistryRequest
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&dino)
	if err != nil {
		fmt.Println("Erro na criação de dino: ", err)
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte("Informações do dino inválidas!"))
		return
	}

	err2 := db.CreateDino(database, dino)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(err2.Error()))
		return
	}
}

func deleteDino(response http.ResponseWriter, request *http.Request) {
	id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		fmt.Println("Erro no parse do id: ", err)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	err2 := db.DeleteDino(database, id)

	if err2 != nil {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(err2.Error()))
		return
	}
}

func dino(response http.ResponseWriter, request *http.Request) {

	id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)
	if err != nil {
		fmt.Println("Erro no parse do id: ", err)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	data, err := db.FindDinoById(database, id)
	if err != nil {
		fmt.Println("Erro ao recuperar dino: ", err)
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	sendJson(response, data)
}

func dinos(response http.ResponseWriter, request *http.Request) {

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

	dinos, searchErr := db.FindDinoByFilter(database, filter)
	if searchErr != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(searchErr.Error()))
		return
	}

	sendJson(response, dinos)
}

//funções auxiliares

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

func sendJson(response http.ResponseWriter, data interface{}) {
	encoder := json.NewEncoder(response)
	response.Header().Add("Content-Type", "application/json; charset=utf-8")
	err := encoder.Encode(data)

	if err != nil {
		fmt.Println("Erro no parse json: ", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
		return
	}
}
