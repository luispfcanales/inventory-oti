package postgre

import (
	"database/sql"
	"log"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) DeletePerson(dni int) bool {
	str := `DELETE FROM person WHERE id_dni = $1`
	res, err := db.getConnection().Exec(str, dni)
	if err != nil {
		return false
	}

	_, err = res.RowsAffected()
	if err != nil {
		return false
	}

	return true
}

func (db *dbConfig) InsertPerson(p models.Person) (models.Person, error) {
	str := `
		INSERT INTO person(id_dni, first_name, last_name, birthdate, address, created_at, update_at)
		VALUES ($1, $2, $3, $4, $5, NOW(), NOW())
	`
	_, err := db.getConnection().Exec(
		str,
		p.IDPerson,
		p.FirstName,
		p.LastName,
		//p.Birthdate.Format("2006-01-02"),
		p.Birthdate,
		p.Address,
	)
	if err != nil {
		return models.Person{}, err
	}

	return p, nil
}

func (db *dbConfig) SelectPerson(dni int) (models.Person, error) {
	var p models.Person

	stmt, err := db.getConnection().Prepare(`
		SELECT id_dni, first_name, last_name, birthdate, address
		FROM person
		WHERE id_dni = $1
	`)
	if err != nil {
		log.Println(err)
		return p, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(dni)
	err = row.Scan(
		&p.IDPerson,
		&p.FirstName,
		&p.LastName,
		&p.Birthdate,
		&p.Address,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no person with the provided credentials")
			return p, err
		}
		return p, err
	}

	return p, nil
}

func (db *dbConfig) SelectPersons() []models.Person {
	var list []models.Person

	qstr := "SELECT id_dni, first_name, last_name, birthdate, address FROM person"
	rows, err := db.getConnection().Query(qstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := models.Person{}

		err := rows.Scan(
			&p.IDPerson,
			&p.FirstName,
			&p.LastName,
			&p.Birthdate,
			&p.Address,
		)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, p)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}

func (db *dbConfig) UpdatePerson(p models.Person) (models.Person, error) {
	queryStr := `
		UPDATE person SET
		first_name = $1,
		last_name = $2,
		birthdate = $3,
		address = $4
		WHERE id_dni = $5
	`

	_, err := db.getConnection().Exec(
		queryStr,
		p.FirstName,
		p.LastName,
		p.Birthdate.Format("2006-01-02"),
		p.Address,
		p.IDPerson,
	)
	if err != nil {
		return models.Person{}, err
	}

	return p, nil
}
