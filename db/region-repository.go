package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

type RegionRepository struct {
	database *sql.DB
}

func NewRegionRepository(database *sql.DB) *RegionRepository {
	return &RegionRepository{database: database}
}

func (r *RegionRepository) CreateRegion(region models.CategoryRegistryRequest) error {
	rows, err := r.database.Query("INSERT INTO region(name_region) VALUES($1);", region.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *RegionRepository) DeleteRegion(id uint64) error {
	rows, err := r.database.Query("DELETE FROM region WHERE id_region = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *RegionRepository) ListAllRegions() ([]models.Category, error) {
	rows, err := r.database.Query("SELECT id_region, name_region FROM region;")
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

func (r *RegionRepository) ExistsRegionById(id uint64) (bool, error) {
	return existsCategoryById(r.database, "region", id)
}
