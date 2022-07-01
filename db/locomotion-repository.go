package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateLocomotion(db *sql.DB, locomotion models.CategoryRegistryRequest) *util.AppError {
	rows, err := db.Query("INSERT INTO locomotion(name_locomotion) VALUES($1);", locomotion.Name)
	if err != nil {
		return util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
	}
	defer rows.Close()

	return nil
}

func ListAllLocomotions(db *sql.DB) ([]models.Category, *util.AppError) {
	rows, err := db.Query("SELECT id_locomotion, name_locomotion FROM locomotion;")
	if err != nil {
		return []models.Category{}, util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
	}
	defer rows.Close()

	result := []models.Category{}
	var locomotion models.Category

	for rows.Next() {
		err2 := rows.Scan(&locomotion.Id, &locomotion.Name)
		if err2 != nil {
			return []models.Category{}, util.ThrowAppError("Erro interno do servidor, por favor tente novamente mais tarde!")
		}
		result = append(result, locomotion)
	}

	return result, nil
}
