package db

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/joaohudson/registro-ark/models"
	"github.com/joaohudson/registro-ark/util"
)

func CreateDino(db *sql.DB, dino models.DinoRegistryRequest) *util.AppError {

	if existsDino(db, dino.Name) {
		return util.ThrowAppError("Já existe um dino com esse nome!")
	}

	if !existsCategory(db, "locomotion", dino.LocomotionId) {
		return util.ThrowAppError("Informe uma locomoção válida!")
	}

	if !existsCategory(db, "region", dino.RegionId) {
		return util.ThrowAppError("Informe uma região válida!")
	}

	if !existsCategory(db, "food", dino.FoodId) {
		return util.ThrowAppError("Informe um tipo de alimentação válida!")
	}

	rows, err := db.Query("INSERT INTO dino(name_dino, id_food, id_locomotion, id_region, utility_dino, training_dino) VALUES($1, $2, $3, $4, $5, $6);", dino.Name, dino.FoodId, dino.LocomotionId, dino.RegionId, dino.Utility, dino.Training)
	if err != nil {
		fmt.Println("Erro ao inserir dados no banco: ", err)
		return util.ThrowAppError("Erro interno do servidor, tente novamente mais tarde.")
	}
	defer rows.Close()

	return nil
}

func FindDinoByFilter(db *sql.DB, filter models.DinoFilter) ([]models.Dino, *util.AppError) {
	const query = `SELECT 
	d.id_dino,
	d.name_dino, 
	f.name_food, 
	l.name_locomotion, 
	r.name_region, 
	d.utility_dino, 
	d.training_dino 
	FROM dino d
	INNER JOIN locomotion l ON d.id_locomotion = l.id_locomotion
	INNER JOIN region r ON d.id_region = r.id_region
	INNER JOIN food f ON d.id_food = f.id_food 
	WHERE 
	(d.name_dino = $1 OR $1 = '') AND
	(d.id_region = $2 OR $2 = 0) AND
	(d.id_locomotion = $3 OR $3 = 0) AND
	(d.id_food = $4 OR $4 = 0);`

	rows, err := db.Query(query, filter.Name, filter.RegionId, filter.LocomotionId, filter.FoodId)
	if err != nil {
		fmt.Println("Erro ao buscar dinos por filtro: ", err)
		return []models.Dino{}, util.ThrowAppError("Não foi possível efetuar a busca.")
	}
	defer rows.Close()

	result := []models.Dino{}
	var dino models.Dino

	for rows.Next() {
		err2 := rows.Scan(&dino.Id, &dino.Name, &dino.Food, &dino.Locomotion, &dino.Region, &dino.Utility, &dino.Training)
		if err2 != nil {
			fmt.Println("Erro ao escanear dinos por filtro: ", err2)
			return []models.Dino{}, util.ThrowAppError("Não foi possível efetuar a busca.")
		}
		result = append(result, dino)
	}

	return result, nil
}

func FindDinoById(db *sql.DB, id uint64) (models.Dino, error) {
	const query = `SELECT 
	d.id_dino,
	d.name_dino, 
	f.name_food, 
	l.name_locomotion, 
	r.name_region, 
	d.utility_dino, 
	d.training_dino 
	FROM dino d
	INNER JOIN locomotion l ON d.id_locomotion = l.id_locomotion
	INNER JOIN region r ON d.id_region = r.id_region
	INNER JOIN food f ON d.id_food = f.id_food 
	WHERE d.id_dino = $1;`

	rows, err := db.Query(query, id)
	if err != nil {
		return models.Dino{}, err
	}
	defer rows.Close()

	var dino models.Dino
	if !rows.Next() {
		return models.Dino{}, util.ThrowAppError("Query vazia para id especificado: " + strconv.FormatUint(id, 10))
	}
	err2 := rows.Scan(&dino.Id, &dino.Name, &dino.Food, &dino.Locomotion, &dino.Region, &dino.Utility, &dino.Training)

	if err2 != nil {
		return models.Dino{}, err2
	}

	return dino, nil
}

//Funções auxiliares

func existsCategory(db *sql.DB, categoryName string, id uint64) bool {
	query := "SELECT * FROM " + categoryName + " WHERE id_" + categoryName + " = $1;"
	rows, err := db.Query(query, id)

	if err != nil {
		fmt.Printf("Erro ao verificar categoria de dino %v para id %v: %v\n", categoryName, id, err)
		return false
	}
	defer rows.Close()

	return rows.Next()
}

func existsDino(db *sql.DB, dinoName string) bool {
	rows, err := db.Query("SELECT * FROM dino WHERE name_dino = $1;", dinoName)
	if err != nil {
		fmt.Println("Erro ao verificar dino por nome: ", err)
		return false
	}
	defer rows.Close()

	return rows.Next()
}
