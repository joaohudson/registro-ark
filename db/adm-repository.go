package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

func CreateAdm(db *sql.DB, adm models.AdmRegistryRequest) error {
	const query = `
		INSERT INTO 
		adm(name_adm, password_adm, permission_manager_adm, permission_manager_category, permission_manager_dino)
		VALUES($1, $2, $3, $4, $5);
	`
	rows, err := db.Query(query, adm.Name, adm.Password, adm.PermissionManagerAdm, adm.PermissionManagerCategory, adm.PermissionManagerDino)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func ExistsAdmByName(db *sql.DB, name string) (bool, error) {
	rows, err := db.Query("SELECT * FROM adm WHERE name_adm = $1;", name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
