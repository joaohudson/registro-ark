package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateFood(db *sql.DB, food models.CategoryRegistryRequest) *util.AppError {
	rows, err := db.Query("INSERT INTO food(name_food) VALUES($1);", food.Name)
	if err != nil {
		return util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
	}
	defer rows.Close()

	return nil
}

func ListAllFoods(db *sql.DB) ([]models.Category, *util.AppError) {
	rows, err := db.Query("SELECT id_food, name_food FROM food;")
	if err != nil {
		return []models.Category{}, util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
	}
	defer rows.Close()

	result := []models.Category{}
	var food models.Category

	for rows.Next() {
		err2 := rows.Scan(&food.Id, &food.Name)
		if err2 != nil {
			return []models.Category{}, util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
		}
		result = append(result, food)
	}

	return result, nil
}
