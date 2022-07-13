package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

type FoodRepository struct {
	database *sql.DB
}

func NewFoodRepository(database *sql.DB) *FoodRepository {
	return &FoodRepository{database: database}
}

func (f *FoodRepository) CreateFood(food models.CategoryRegistryRequest) error {
	rows, err := f.database.Query("INSERT INTO food(name_food) VALUES($1);", food.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (f *FoodRepository) DeleteFood(id uint64) error {
	rows, err := f.database.Query("DELETE FROM food WHERE id_food = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (f *FoodRepository) ListAllFoods() ([]models.Category, error) {
	rows, err := f.database.Query("SELECT id_food, name_food FROM food;")
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

func (f *FoodRepository) ExistsFoodById(id uint64) (bool, error) {
	return existsCategoryById(f.database, "food", id)
}
