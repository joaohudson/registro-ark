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

func (a *AdmRepository) GetAdms() ([]models.Adm, error) {
	rows, err := a.database.Query("SELECT id_adm, name_adm, permission_manager_dino, permission_manager_category, permission_manager_adm FROM adm;")
	if err != nil {
		return []models.Adm{}, err
	}
	defer rows.Close()

	var adm models.Adm
	result := []models.Adm{}

	for rows.Next() {
		err = rows.Scan(&adm.Id, &adm.Name, &adm.PermissionManagerDino, &adm.PermissionManagerCategory, &adm.PermissionManagerAdm)
		if err != nil {
			return []models.Adm{}, err
		}
		result = append(result, adm)
	}

	return result, nil
}

func (a *AdmRepository) GetAdmById(id uint64) (*models.Adm, error) {
	rows, err := a.database.Query("SELECT id_adm, name_adm, permission_manager_dino, permission_manager_category, permission_manager_adm FROM adm WHERE id_adm = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	result := &models.Adm{}
	rows.Scan(&result.Id, &result.Name, &result.PermissionManagerDino, &result.PermissionManagerCategory, &result.PermissionManagerAdm)

	return result, nil
}

func (a *AdmRepository) GetAdmIdByCredentials(name string, password string) (uint64, error) {
	rows, err := a.database.Query("SELECT id_adm FROM adm WHERE name_adm = $1 AND password_adm = $2;", name, password)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, nil
	}

	var id uint64
	rows.Scan(&id)
	return id, nil
}

func (a *AdmRepository) ExistsAdmByName(name string) (bool, error) {
	rows, err := a.database.Query("SELECT * FROM adm WHERE name_adm = $1;", name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (a *AdmRepository) GetMainAdmId() (uint64, error) {
	rows, err := a.database.Query("SELECT id_adm FROM adm WHERE permission_manager_adm = true;")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, nil
	}

	var id uint64
	rows.Scan(&id)
	return id, nil
}
