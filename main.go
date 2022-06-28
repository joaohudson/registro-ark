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

	data, err := db.ListAllDinos(database)
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

func createDino(response http.ResponseWriter, request *http.Request) {
	var dino models.Dino
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
		fmt.Println("Erro ao inserir dino no banco: ", err2)
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
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

	result := []models.Dino{}

	data, err := db.ListAllDinos(database)
	if err != nil {
		fmt.Println("Erro ao listar dinos: ", err)
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
