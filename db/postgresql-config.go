package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const postgresDriver = "postgres"

const user = "postgres"

const host = "localhost"

const port = "5432"

const password = "postgresql"

const dbName = "postgres"

const tableName = "dino"

func CreatePostgresDatabase() *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := sql.Open(postgresDriver, dataSourceName)

	if err != nil {
		fmt.Println("Erro ao inciar banco de dados: ", err)
		panic("Não foi possível iniciar a conexão com o banco de dados")
	}

	return db
}
