package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateRegion(db *sql.DB, region models.CategoryRegistryRequest) *util.AppError {
	rows, err := db.Query("INSERT INTO region(name_region) VALUES($1);", region.Name)
	if err != nil {
		return util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
	}
	defer rows.Close()

	return nil
}

func ListAllRegions(db *sql.DB) ([]models.Category, *util.AppError) {
	rows, err := db.Query("SELECT id_region, name_region FROM region;")
	if err != nil {
		return []models.Category{}, util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
	}
	defer rows.Close()

	result := []models.Category{}
	var region models.Category

	for rows.Next() {
		err2 := rows.Scan(&region.Id, &region.Name)
		if err2 != nil {
			return []models.Category{}, util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
		}
		result = append(result, region)
	}

	return result, nil
}
