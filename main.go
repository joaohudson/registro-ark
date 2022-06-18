package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Dino struct {
	Id         uint64 `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	Locomotion string `json:"locomotion"`
	Food       string `json:"food"`
	Training   string `json:"training"`
	Utility    string `json:"utility"`
}

type DinoCategoryResponse struct {
	Regions     []string `json:"regions"`
	Locomotions []string `json:"locomotions"`
	Foods       []string `json:"foods"`
}

const DefaultInternalServerErrorMessage = "Erro interno do servidor, por favor tente mais tarde."

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/api/dino", dinos)
	http.HandleFunc("/api/dino/category", dinoCategories)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}

func dinoCategories(response http.ResponseWriter, request *http.Request) {
	result := DinoCategoryResponse{
		Regions:     []string{},
		Locomotions: []string{},
		Foods:       []string{},
	}

	data, err := readData("data.json")
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

func dinos(response http.ResponseWriter, request *http.Request) {

	result := []Dino{}

	data, err := readData("data.json")
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

func readData(filename string) ([]Dino, error) {
	fileReader, err := os.Open(filename)

	if err != nil {
		return []Dino{}, err
	}

	var obj []Dino
	decoder := json.NewDecoder(fileReader)
	err2 := decoder.Decode(&obj)

	if err2 != nil {
		return []Dino{}, err2
	}

	return obj, nil
}
