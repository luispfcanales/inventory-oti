package postgre

import (
	"database/sql"
	"log"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) SelectStaff() ([]models.Staff, error) {
	var staff []models.Staff
	rows, err := db.getConnection().Query(`
		SELECT 
		id,name
		FROM staff
	`)
	if err != nil {
		log.Println(err)
		return staff, err
	}
	defer rows.Close()

	for rows.Next() {
		u := models.Staff{}
		err = rows.Scan(
			&u.ID,
			&u.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		staff = append(staff, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return staff, nil
}
func (db *dbConfig) SelectRole() ([]models.Role, error) {
	var role []models.Role
	rows, err := db.getConnection().Query(`
		SELECT 
		id,name
		FROM role
	`)
	if err != nil {
		log.Println(err)
		return role, err
	}
	defer rows.Close()

	for rows.Next() {
		u := models.Role{}
		err = rows.Scan(
			&u.ID,
			&u.Name,
		)
		if err != nil {
			log.Fatal(err)
		}
		role = append(role, u)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return role, nil
}

func (db *dbConfig) SelectUserWithCredentials(email string, pwd string) (models.User, error) {
	u := models.User{}
	stmt, err := db.getConnection().Prepare(`
		SELECT 
		u.id_person, u.email, u.password, u.active, role.name,
		staff.name AS staff,
		p.first_name, p.last_name
		FROM users u
		JOIN staff on u.id_staff = staff.id 
		JOIN person p on u.id_person = p.id_dni
		JOIN role on role.id = u.id_role
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
		&u.RoleName,
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
		u.id_person, u.email, u.password, u.active, role.name,
		staff.name AS staff,
		p.first_name, p.last_name, p.birthdate
		FROM users u
		JOIN staff on u.id_staff = staff.id 
		JOIN person p on u.id_person = p.id_dni
		JOIN role on role.id = u.id_role
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
			&u.RoleName,
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

func (db *dbConfig) InsertUser(u *models.User) error {
	str := `
		INSERT INTO users(
			id, email,password,active,id_person,id_role,
			update_at,
			id_dependency, id_staff
		)
		VALUES ($1, $2, $3, $4, $5, $6, NOW(), $7, $8)
	`
	_, err := db.getConnection().Exec(
		str,
		u.Key,
		u.Email,
		"12345678",
		u.Active,
		u.IDPerson,
		u.RoleName,
		u.IDDependency,
		u.Staff,
	)
	if err != nil {
		return err
	}

	return nil
}
