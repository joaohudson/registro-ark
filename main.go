package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/joaohudson/registro-ark/db"
	"github.com/joaohudson/registro-ark/models"
)

const DefaultInternalServerErrorMessage = "Erro interno do servidor, por favor tente mais tarde."

func main() {

	router := chi.NewRouter()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/*", fs)

	router.Get("/api/dino", dino)
	router.Get("/api/dinos", dinos)
	router.Get("/api/dino/categories", dinoCategories)

	err := http.ListenAndServe(":8081", router)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}

func dinoCategories(response http.ResponseWriter, request *http.Request) {
	result := models.DinoCategoryResponse{
		Regions:     []string{},
		Locomotions: []string{},
		Foods:       []string{},
	}

	data, err := db.ReadData("data.json")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
		fmt.Println("Erro ao ler arquivo json: ", err)
		return
	}

	regionSet := make(map[string]bool)
	locomotionSet := make(map[string]bool)
	foodSet := make(map[string]bool)
	for i := range data {
		regionSet[data[i].Region] = true
		locomotionSet[data[i].Locomotion] = true
		foodSet[data[i].Food] = true
	}

	for k := range regionSet {
		result.Regions = append(result.Regions, k)
	}

	for k := range locomotionSet {
		result.Locomotions = append(result.Locomotions, k)
	}

	for k := range foodSet {
		result.Foods = append(result.Foods, k)
	}

	sendJson(response, result)
}

func dino(response http.ResponseWriter, request *http.Request) {

	data, err := db.ReadData("data.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo json: ", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
		return
	}

	id, err := strconv.ParseUint(request.URL.Query().Get("id"), 10, 64)

	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("Dino não encontrado!"))
		return
	}

	for i := range data {

		if id == data[i].Id {
			sendJson(response, data[i])
			break
		}
	}
}

func dinos(response http.ResponseWriter, request *http.Request) {

	result := []models.Dino{}

	data, err := db.ReadData("data.json")
	if err != nil {
		fmt.Println("Erro ao ler arquivo json: ", err)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
		return
	}

	region := request.URL.Query().Get("region")
	locomotion := request.URL.Query().Get("locomotion")
	food := request.URL.Query().Get("food")
	name := request.URL.Query().Get("name")

	for i := range data {
		regionMatch := region == "" || region == data[i].Region
		locomotionMatch := locomotion == "" || locomotion == data[i].Locomotion
		foodMatch := food == "" || food == data[i].Food
		nameMatch := name == "" || name == data[i].Name

		if regionMatch && locomotionMatch && foodMatch && nameMatch {
			result = append(result, data[i])
		}
	}

	sendJson(response, result)
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
