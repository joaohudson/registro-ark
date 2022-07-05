package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

func CreateLocomotion(db *sql.DB, locomotion models.CategoryRegistryRequest) error {
	rows, err := db.Query("INSERT INTO locomotion(name_locomotion) VALUES($1);", locomotion.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func ListAllLocomotions(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query("SELECT id_locomotion, name_locomotion FROM locomotion;")
	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	result := []models.Category{}
	var locomotion models.Category

	for rows.Next() {
		err2 := rows.Scan(&locomotion.Id, &locomotion.Name)
		if err2 != nil {
			return []models.Category{}, err2
		}
		result = append(result, locomotion)
	}

	return result, nil
}

func ExistsLocomotionById(db *sql.DB, id uint64) (bool, error) {
	return existsCategoryById(db, "locomotion", id)
}
