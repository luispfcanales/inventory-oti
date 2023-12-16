package postgre

import (
	"database/sql"
	"errors"
	"log"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) UpdateCampus(c models.Campus) (models.Campus, error) {
	queryStr := `
		UPDATE campus SET
		abbreviation = $1,
		name = $2,
		address = $3,
		state = $4
		WHERE id = $5
	`

	r, err := db.getConnection().Exec(
		queryStr,
		c.Abbreviation,
		c.Name,
		c.Address,
		c.State,
		c.ID,
	)
	if err != nil {
		return models.Campus{}, err
	}

	affect, err := r.RowsAffected()
	if err != nil {
		return models.Campus{}, err
	}
	if affect == 0 {
		return models.Campus{}, errors.New("Not found register to update")
	}

	return c, nil
}
func (db *dbConfig) SelectCampus(id string) (models.Campus, error) {
	var p models.Campus

	stmt, err := db.getConnection().Prepare(`
		SELECT abbreviation, name, address, state, id
		FROM campus
		WHERE id = $1
	`)
	if err != nil {
		log.Println(err)
		return p, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&p.Abbreviation,
		&p.Name,
		&p.Address,
		&p.State,
		&p.ID,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[ no campus with the provided credentials ]")
			return p, err
		}
		return p, err
	}

	return p, nil
}

func (db *dbConfig) SelectAllCampus() []models.Campus {
	var list []models.Campus

	qstr := "SELECT abbreviation, name, address, state, id FROM campus"
	rows, err := db.getConnection().Query(qstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		c := models.Campus{}

		err := rows.Scan(
			&c.Abbreviation,
			&c.Name,
			&c.Address,
			&c.State,
			&c.ID,
		)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, c)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}

func (db *dbConfig) DeleteCampus(id string) error {
	str := `delete from campus where id=$1`
	_, err := db.getConnection().Exec(str, id)
	if err != nil {
		return err
	}
	return nil
}

func (db *dbConfig) InsertCampus(c models.Campus) error {
	str := `
		INSERT INTO campus(abbreviation, name, address, state, created_at, update_at,id)
		VALUES ($1, $2, $3, $4, NOW(), NOW(), $5)
	`
	_, err := db.getConnection().Exec(
		str,
		c.Abbreviation,
		c.Name,
		c.Address,
		c.State,
		c.ID,
	)
	if err != nil {
		return err
	}

	return nil
}
