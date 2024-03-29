package db

import (
	"database/sql"
	"errors"
	"time"

	"github.com/joaohudson/registro-ark/models"
)

type DinoRepository struct {
	db *sql.DB
}

func NewDinoRepository(db *sql.DB) *DinoRepository {
	return &DinoRepository{db: db}
}

func (r *DinoRepository) CreateDino(idAdm uint64, dino models.DinoRegistryRequest) (uint64, error) {
	now := time.Now()

	const query = `
	INSERT INTO 
	dino(name_dino, id_food, id_locomotion, id_region, id_adm, dt_creation, utility_dino, training_dino)
	VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id_dino;`

	rows, err := r.db.Query(query, dino.Name, dino.FoodId, dino.LocomotionId, dino.RegionId, idAdm, now, dino.Utility, dino.Training)
	if err != nil {

		return 0, err
	}
	defer rows.Close()

	var idDino uint64

	if !rows.Next() {
		return 0, errors.New("Erro inesperado ao tentar buscar o id durante insert do dino")
	}
	err = rows.Scan(&idDino)
	if err != nil {
		return 0, err
	}

	return idDino, nil
}

func (r *DinoRepository) PutDino(idDino uint64, idAdm uint64, dino models.DinoRegistryRequest) error {
	const query = `
	UPDATE dino 
	SET name_dino = $1, id_food = $2, id_locomotion = $3, id_region = $4, utility_dino = $5, training_dino = $6
	WHERE id_dino = $7;`

	rows, err := r.db.Query(query, dino.Name, dino.FoodId, dino.LocomotionId, dino.RegionId, dino.Utility, dino.Training, idDino)
	if err != nil {

		return err
	}
	defer rows.Close()

	return nil
}

func (r *DinoRepository) FindDinoByFilter(filter models.DinoFilter) ([]models.Dino, error) {
	const query = `SELECT 
	d.id_dino,
	d.id_adm,
	d.name_dino, 
	f.id_food,
	f.name_food,
	l.id_locomotion, 
	l.name_locomotion,
	r.id_region, 
	r.name_region, 
	d.dt_creation,
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
	(d.id_food = $4 OR $4 = 0)
	ORDER BY(d.name_dino) ASC;`

	rows, err := r.db.Query(query, filter.Name, filter.RegionId, filter.LocomotionId, filter.FoodId)
	if err != nil {
		return []models.Dino{}, err
	}
	defer rows.Close()

	result := []models.Dino{}
	var dino models.Dino

	for rows.Next() {
		err2 := rows.Scan(&dino.Id, &dino.IdAdm, &dino.Name, &dino.Food.Id, &dino.Food.Name, &dino.Locomotion.Id, &dino.Locomotion.Name, &dino.Region.Id, &dino.Region.Name, &dino.CreationDate, &dino.Utility, &dino.Training)
		if err2 != nil {
			return []models.Dino{}, err2
		}
		result = append(result, dino)
	}

	return result, nil
}

func (r *DinoRepository) FindDinoByFilterForAdm(idAdm uint64, filter models.DinoFilter) ([]models.Dino, error) {
	const query = `SELECT 
	d.id_dino,
	d.id_adm,
	d.name_dino, 
	f.id_food,
	f.name_food,
	l.id_locomotion, 
	l.name_locomotion,
	r.id_region, 
	r.name_region, 
	d.dt_creation,
	d.utility_dino, 
	d.training_dino 
	FROM dino d
	INNER JOIN locomotion l ON d.id_locomotion = l.id_locomotion
	INNER JOIN region r ON d.id_region = r.id_region
	INNER JOIN food f ON d.id_food = f.id_food
	INNER JOIN adm a ON a.id_adm = d.id_adm 
	WHERE 
	(d.name_dino = $1 OR $1 = '') AND
	(d.id_region = $2 OR $2 = 0) AND
	(d.id_locomotion = $3 OR $3 = 0) AND
	(d.id_food = $4 OR $4 = 0) AND
	(a.id_adm = $5)
	ORDER BY(d.dt_creation) DESC;`

	rows, err := r.db.Query(query, filter.Name, filter.RegionId, filter.LocomotionId, filter.FoodId, idAdm)
	if err != nil {
		return []models.Dino{}, err
	}
	defer rows.Close()

	result := []models.Dino{}
	var dino models.Dino

	for rows.Next() {
		err2 := rows.Scan(&dino.Id, &dino.IdAdm, &dino.Name, &dino.Food.Id, &dino.Food.Name, &dino.Locomotion.Id, &dino.Locomotion.Name, &dino.Region.Id, &dino.Region.Name, &dino.CreationDate, &dino.Utility, &dino.Training)
		if err2 != nil {
			return []models.Dino{}, err2
		}
		result = append(result, dino)
	}

	return result, nil
}

func (r *DinoRepository) FindDinoByFilterForMainAdm(filter models.DinoFilter) ([]models.Dino, error) {
	const query = `SELECT 
	d.id_dino,
	d.id_adm,
	d.name_dino, 
	f.id_food,
	f.name_food,
	l.id_locomotion, 
	l.name_locomotion,
	r.id_region, 
	r.name_region, 
	d.dt_creation,
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
	(d.id_food = $4 OR $4 = 0)
	ORDER BY(d.dt_creation) DESC;`

	rows, err := r.db.Query(query, filter.Name, filter.RegionId, filter.LocomotionId, filter.FoodId)
	if err != nil {
		return []models.Dino{}, err
	}
	defer rows.Close()

	result := []models.Dino{}
	var dino models.Dino

	for rows.Next() {
		err2 := rows.Scan(&dino.Id, &dino.IdAdm, &dino.Name, &dino.Food.Id, &dino.Food.Name, &dino.Locomotion.Id, &dino.Locomotion.Name, &dino.Region.Id, &dino.Region.Name, &dino.CreationDate, &dino.Utility, &dino.Training)
		if err2 != nil {
			return []models.Dino{}, err2
		}
		result = append(result, dino)
	}

	return result, nil
}

func (r *DinoRepository) FindDinoById(id uint64) (*models.Dino, error) {
	const query = `SELECT 
	d.id_dino,
	d.id_adm,
	d.name_dino, 
	f.id_food,
	f.name_food,
	l.id_locomotion, 
	l.name_locomotion,
	r.id_region, 
	r.name_region, 
	d.dt_creation,
	d.utility_dino, 
	d.training_dino
	FROM dino d
	INNER JOIN locomotion l ON d.id_locomotion = l.id_locomotion
	INNER JOIN region r ON d.id_region = r.id_region
	INNER JOIN food f ON d.id_food = f.id_food 
	WHERE d.id_dino = $1;`

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var dino models.Dino
	if !rows.Next() {
		return nil, nil
	}
	err2 := rows.Scan(&dino.Id, &dino.IdAdm, &dino.Name, &dino.Food.Id, &dino.Food.Name, &dino.Locomotion.Id, &dino.Locomotion.Name, &dino.Region.Id, &dino.Region.Name, &dino.CreationDate, &dino.Utility, &dino.Training)

	if err2 != nil {
		return nil, err2
	}

	return &dino, nil
}

func (r *DinoRepository) DeleteDino(id uint64) error {
	rows, err := r.db.Query("DELETE FROM dino WHERE id_dino = $1;", id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (r *DinoRepository) ExistsDinoById(id uint64) (bool, error) {
	rows, err := r.db.Query("SELECT * FROM dino WHERE id_dino = $1;", id)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (r *DinoRepository) ExistsDinoByName(dinoName string) (bool, error) {
	rows, err := r.db.Query("SELECT * FROM dino WHERE name_dino = $1;", dinoName)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}

func (r *DinoRepository) FindDinoIdByName(name string) (uint64, error) {
	rows, err := r.db.Query("SELECT id_dino FROM dino WHERE name_dino = $1;", name)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, nil
	}

	var id uint64
	err = rows.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
