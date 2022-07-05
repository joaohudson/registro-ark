package db

import (
	"database/sql"
	"fmt"
)

func existsCategoryById(db *sql.DB, categoryName string, id uint64) bool {
	query := "SELECT * FROM " + categoryName + " WHERE id_" + categoryName + " = $1;"
	rows, err := db.Query(query, id)

	if err != nil {
		fmt.Printf("Erro ao verificar categoria de dino %v para id %v: %v\n", categoryName, id, err)
		return false
	}
	defer rows.Close()

	return rows.Next()
}
