package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Dino struct {
	Name   string
	Owner  string
	Action string
}

const DefaultInternalServerErrorMessage = "Erro interno do servidor, por favor tente mais tarde."

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/dino", dinos)
	http.HandleFunc("/owner", owners)
	http.HandleFunc("/dino/name", dinoNames)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Erro ao iniciar servidor: ", err)
	}
}

func dinoNames(response http.ResponseWriter, request *http.Request) {
	result := []string{}

	data, err := readData("data.json")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
		fmt.Println("Erro ao ler arquivo json: ", err)
		return
	}

	set := make(map[string]bool)
	for i := range data {
		set[data[i].Name] = true
	}

	for k := range set {
		result = append(result, k)
	}

	sendJson(response, result)
}

func owners(response http.ResponseWriter, request *http.Request) {
	result := []string{}

	data, err := readData("data.json")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(DefaultInternalServerErrorMessage))
		fmt.Println("Erro ao ler arquivo json: ", err)
		return
	}

	set := make(map[string]bool)
	for i := range data {
		set[data[i].Owner] = true
	}

	for k := range set {
		result = append(result, k)
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

	owner := request.URL.Query().Get("owner")
	name := request.URL.Query().Get("name")

	for i := range data {
		ownerMatch := owner == "" || owner == data[i].Owner
		nameMatch := name == "" || name == data[i].Name

		if ownerMatch && nameMatch {
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
