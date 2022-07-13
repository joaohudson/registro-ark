package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

type LocomotionRepository struct {
	database *sql.DB
}

func NewLocomotionRepository(database *sql.DB) *LocomotionRepository {
	return &LocomotionRepository{database: database}
}

func (l *LocomotionRepository) CreateLocomotion(locomotion models.CategoryRegistryRequest) error {
	rows, err := l.database.Query("INSERT INTO locomotion(name_locomotion) VALUES($1);", locomotion.Name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (l *LocomotionRepository) DeleteLocomotion(id uint64) error {
	rows, err := l.database.Query("DELETE FROM locomotion WHERE id_locomotion = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (l *LocomotionRepository) ListAllLocomotions() ([]models.Category, error) {
	rows, err := l.database.Query("SELECT id_locomotion, name_locomotion FROM locomotion;")
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

func (l *LocomotionRepository) ExistsLocomotionById(id uint64) (bool, error) {
	return existsCategoryById(l.database, "locomotion", id)
}
