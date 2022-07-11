package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

func CreateFood(db *sql.DB, food models.CategoryRegistryRequest) error {
	rows, err := db.Query("INSERT INTO food(name_food) VALUES($1);", food.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DeleteFood(db *sql.DB, id uint64) error {
	rows, err := db.Query("DELETE FROM food WHERE id_food = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func ListAllFoods(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query("SELECT id_food, name_food FROM food;")
	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	result := []models.Category{}
	var food models.Category

	for rows.Next() {
		err2 := rows.Scan(&food.Id, &food.Name)
		if err2 != nil {
			return []models.Category{}, err2
		}
		result = append(result, food)
	}

	return result, nil
}

func ExistsFoodById(db *sql.DB, id uint64) (bool, error) {
	return existsCategoryById(db, "food", id)
}
