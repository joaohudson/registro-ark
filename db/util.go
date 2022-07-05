package db

import (
	"database/sql"
)

func existsCategoryById(db *sql.DB, categoryName string, id uint64) (bool, error) {
	query := "SELECT * FROM " + categoryName + " WHERE id_" + categoryName + " = $1;"
	rows, err := db.Query(query, id)

	if err != nil {

		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
