package db

import "database/sql"

type ImageRepository struct {
	database *sql.DB
}

func NewImageRepository(database *sql.DB) *ImageRepository {
	return &ImageRepository{database: database}
}

func (this *ImageRepository) Insert(name string, data string) error {
	rows, err := this.database.Query("INSERT INTO img(name_img, content_img) VALUES($1, $2);", name, data)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (this *ImageRepository) Get(name string) (string, error) {
	rows, err := this.database.Query("SELECT content_img FROM img WHERE name_img = $1;", name)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	if !rows.Next() {
		return "", nil
	}

	data := ""
	err = rows.Scan(&data)
	if err != nil {
		return "", err
	}

	return data, nil
}

func (this *ImageRepository) Delete(name string) error {
	rows, err := this.database.Query("DELETE FROM img WHERE name_img = $1;", name)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}

func (this *ImageRepository) Put(name string, content string) error {

	existsImage, err := this.existsImageByName(name)
	if err != nil {
		return err
	}

	if existsImage {
		rows, err := this.database.Query("UPDATE img SET content_img = $1 WHERE name_img = $2;", content, name)
		if err != nil {
			return err
		}
		defer rows.Close()
	} else {
		rows, err := this.database.Query("INSERT INTO img(name_img, content_img) VALUES($1, $2);", name, content)
		if err != nil {
			return err
		}
		defer rows.Close()
	}

	return nil
}

func (this *ImageRepository) existsImageByName(name string) (bool, error) {
	rows, err := this.database.Query("SELECT name_img FROM img WHERE name_img = $1;", name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
