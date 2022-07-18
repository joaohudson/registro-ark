package db

import (
	"database/sql"
	"fmt"

	"github.com/joaohudson/registro-ark/util"

	_ "github.com/lib/pq"
)

const postgresDriver = "postgres"

func CreatePostgresDatabase() *sql.DB {
	dataSourceName, err := util.GetEnv("DATABASE_URL")
	if err != nil {
		fmt.Println("Erro ao recuperar variável de ambiente: ", err)
		panic(err)
	}

	db, err := sql.Open(postgresDriver, dataSourceName)
	if err != nil {
		fmt.Println("Erro ao inciar banco de dados: ", err)
		panic("Não foi possível iniciar a conexão com o banco de dados")
	}

	return db
}
