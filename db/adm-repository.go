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
		adm(name_adm, password_adm, permission_manager_adm, permission_manager_category, permission_manager_dino, main_adm)
		VALUES($1, $2, false, $3, $4, false);
	`
	rows, err := a.database.Query(query, adm.Name, adm.Password, adm.PermissionManagerCategory, adm.PermissionManagerDino)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (a *AdmRepository) DeleteAdm(id uint64) error {
	txt, err := a.database.Begin()
	if err != nil {
		return err
	}

	_, err = txt.Exec("DELETE FROM login WHERE id_adm = $1", id)
	if err != nil {
		return err
	}
	_, err = txt.Exec("DELETE FROM adm WHERE id_adm = $1", id)
	if err != nil {
		return err
	}

	err = txt.Commit()

	return err
}

func (a *AdmRepository) PutPermissions(permissions models.AdmChangePermissionsRequest) error {
	const query = `UPDATE adm SET
	permission_manager_category = $1,
	permission_manager_dino = $2
	WHERE id_adm = $3;`

	rows, err := a.database.Query(query, permissions.PermissionManagerCategory, permissions.PermissionManagerDino, permissions.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (a *AdmRepository) PutCredentials(idAdm uint64, newName string, newPassword string) error {
	const query = `UPDATE adm SET
	name_adm = $1,
	password_adm = $2
	WHERE id_adm = $3;`

	rows, err := a.database.Query(query, newName, newPassword, idAdm)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (a *AdmRepository) GetAdms() ([]models.Adm, error) {
	rows, err := a.database.Query("SELECT id_adm, name_adm, permission_manager_dino, permission_manager_category, permission_manager_adm, main_adm FROM adm ORDER BY id_adm DESC;")
	if err != nil {
		return []models.Adm{}, err
	}
	defer rows.Close()

	var adm models.Adm
	result := []models.Adm{}

	for rows.Next() {
		err = rows.Scan(&adm.Id, &adm.Name, &adm.PermissionManagerDino, &adm.PermissionManagerCategory, &adm.PermissionManagerAdm, &adm.MainAdm)
		if err != nil {
			return []models.Adm{}, err
		}
		result = append(result, adm)
	}

	return result, nil
}

func (a *AdmRepository) GetAdmById(id uint64) (*models.Adm, error) {
	rows, err := a.database.Query("SELECT id_adm, name_adm, permission_manager_dino, permission_manager_category, permission_manager_adm, main_adm FROM adm WHERE id_adm = $1;", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, nil
	}

	result := &models.Adm{}
	rows.Scan(&result.Id, &result.Name, &result.PermissionManagerDino, &result.PermissionManagerCategory, &result.PermissionManagerAdm, &result.MainAdm)

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
	rows, err := a.database.Query("SELECT id_adm FROM adm WHERE main_adm = true;")
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
