package db

import (
	"encoding/json"
	"os"

	"github.com/joaohudson/registro-ark/models"
)

func ReadData(filename string) ([]models.Dino, error) {
	fileReader, err := os.Open(filename)

	if err != nil {
		return []models.Dino{}, err
	}

	var obj []models.Dino
	decoder := json.NewDecoder(fileReader)
	err2 := decoder.Decode(&obj)

	if err2 != nil {
		return []models.Dino{}, err2
	}

	return obj, nil
}
