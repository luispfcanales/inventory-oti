package postgre

import (
	"database/sql"
	"log"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) SelectUserWithCredentials(email string, pwd string) (models.User, error) {
	u := models.User{}
	stmt, err := db.getConnection().Prepare(`
		SELECT 
		u.id_person, u.email, u.password, u.active, u.id_role,
		staff.name AS staff,
		p.first_name, p.last_name
		FROM users u
		JOIN staff on u.id_staff = staff.id 
		JOIN person p on u.id_person = p.id_dni
		WHERE 
		u.email = $1 AND u.password = $2
	`)
	if err != nil {
		log.Println(err)
		return u, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(email, pwd)
	err = row.Scan(
		&u.IDPerson,
		&u.Email,
		&u.Password,
		&u.Active,
		&u.IDRole,
		&u.Staff,
		&u.FirstName,
		&u.LastName,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no user with the provided credentials")
			return u, err
		}
		return u, err
	}
	return u, nil
}

func (db *dbConfig) SelectUsers() ([]models.User, error) {
	var users []models.User
	rows, err := db.getConnection().Query(`
		SELECT 
		u.id_person, u.email, u.password, u.active, u.id_role,
		staff.name AS staff,
		p.first_name, p.last_name, p.birthdate
		FROM users u
		JOIN staff on u.id_staff = staff.id 
		JOIN person p on u.id_person = p.id_dni
	`)
	if err != nil {
		log.Println(err)
		return users, err
	}
	defer rows.Close()

	for rows.Next() {
		u := models.User{}
		err = rows.Scan(
			&u.IDPerson,
			&u.Email,
			&u.Password,
			&u.Active,
			&u.IDRole,
			&u.Staff,
			&u.FirstName,
			&u.LastName,
			&u.Birthdate,
		)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return users, nil
}

func (db *dbConfig) InsertUser(_ models.User) error {
	panic("not implemented") // TODO: Implement
}
