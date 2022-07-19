package db

import (
	"database/sql"

	"github.com/joaohudson/registro-ark/models"
)

type LoginRepository struct {
	database *sql.DB
}

func NewLoginRepository(database *sql.DB) *LoginRepository {
	return &LoginRepository{database: database}
}

func (l *LoginRepository) SaveClaims(claims *models.Claims) error {
	existsClaims, err := l.existsClaimsByIdAdm(claims.Id)
	if err != nil {
		return err
	}
	if existsClaims {
		rows, err := l.database.Query("UPDATE login SET dt_login = $1 WHERE id_adm = $2;", claims.DateLogin, claims.Id)
		if err != nil {
			return err
		}
		defer rows.Close()
	} else {
		rows, err := l.database.Query("INSERT INTO login(id_adm, dt_login) VALUES($1, $2);", claims.Id, claims.DateLogin)
		if err != nil {
			return err
		}
		defer rows.Close()
	}

	return nil
}

func (l *LoginRepository) ExistsClaims(claims *models.Claims) (bool, error) {
	rows, err := l.database.Query("SELECT id_adm, dt_login FROM login WHERE id_adm = $1 AND dt_login = $2;", claims.Id, claims.DateLogin)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}

func (l *LoginRepository) DeleteClaims(idAdm uint64) error {
	rows, err := l.database.Query("DELETE FROM login WHERE id_adm = $1", idAdm)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (l *LoginRepository) existsClaimsByIdAdm(idAdm uint64) (bool, error) {
	rows, err := l.database.Query("SELECT id_adm FROM login WHERE id_adm = $1", idAdm)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		return false, nil
	}

	return true, nil
}
