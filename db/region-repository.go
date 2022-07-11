package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

func CreateRegion(db *sql.DB, region models.CategoryRegistryRequest) error {
	rows, err := db.Query("INSERT INTO region(name_region) VALUES($1);", region.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func DeleteRegion(db *sql.DB, id uint64) error {
	rows, err := db.Query("DELETE FROM region WHERE id_region = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func ListAllRegions(db *sql.DB) ([]models.Category, error) {
	rows, err := db.Query("SELECT id_region, name_region FROM region;")
	if err != nil {
		return []models.Category{}, err
	}
	defer rows.Close()

	result := []models.Category{}
	var region models.Category

	for rows.Next() {
		err2 := rows.Scan(&region.Id, &region.Name)
		if err2 != nil {
			return []models.Category{}, err2
		}
		result = append(result, region)
	}

	return result, nil
}

func ExistsRegionById(db *sql.DB, id uint64) (bool, error) {
	return existsCategoryById(db, "region", id)
}
