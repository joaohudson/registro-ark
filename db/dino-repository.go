package db

import (
	"database/sql"
	"strconv"

	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateDino(db *sql.DB, dino models.Dino) error {
	_, err := db.Query("INSERT INTO dino(name_dino, food_dino, locomotion_dino, region_dino, utility_dino, training_dino) VALUES($1, $2, $3, $4, $5, $6);", dino.Name, dino.Food, dino.Locomotion, dino.Region, dino.Utility, dino.Training)
	return err
}

func FindDinoById(db *sql.DB, id uint64) (models.Dino, error) {
	rows, err := db.Query("SELECT id_dino, name_dino, food_dino, locomotion_dino, region_dino, utility_dino, training_dino FROM dino WHERE id_dino = $1;", id)

	if err != nil {
		return models.Dino{}, err
	}

	var dino models.Dino
	if !rows.Next() {
		return models.Dino{}, util.ThrowAppError("Query vazia para id especificado: " + strconv.FormatUint(id, 10))
	}
	err2 := rows.Scan(&dino.Id, &dino.Name, &dino.Food, &dino.Locomotion, &dino.Region, &dino.Utility, &dino.Training)

	if err2 != nil {
		return models.Dino{}, err2
	}

	return dino, nil
}

func ListAllDinos(db *sql.DB) ([]models.Dino, error) {
	rows, err := db.Query("SELECT id_dino, name_dino, food_dino, locomotion_dino, region_dino, utility_dino, training_dino FROM dino;")
	var result []models.Dino

	if err != nil {
		return []models.Dino{}, err
	}

	for rows.Next() {
		var dino models.Dino
		err := rows.Scan(&dino.Id, &dino.Name, &dino.Food, &dino.Locomotion, &dino.Region, &dino.Utility, &dino.Training)
		if err != nil {
			return []models.Dino{}, err
		}
		result = append(result, dino)
	}

	return result, nil
}
