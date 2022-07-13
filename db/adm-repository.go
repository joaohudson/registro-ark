package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

type AdmRepository struct {
	database *sql.DB
}

func NewAdmRepository(database *sql.DB) *AdmRepository {
	return &AdmRepository{database: database}
}

func (a *AdmRepository) CreateAdm(adm models.AdmRegistryRequest) error {
	const query = `
		INSERT INTO 
		adm(name_adm, password_adm, permission_manager_adm, permission_manager_category, permission_manager_dino)
		VALUES($1, $2, $3, $4, $5);
	`
	rows, err := a.database.Query(query, adm.Name, adm.Password, adm.PermissionManagerAdm, adm.PermissionManagerCategory, adm.PermissionManagerDino)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (a *AdmRepository) ExistsAdmByName(name string) (bool, error) {
	rows, err := a.database.Query("SELECT * FROM adm WHERE name_adm = $1;", name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
