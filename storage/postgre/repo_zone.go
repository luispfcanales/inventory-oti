package postgre

import (
	"database/sql"
	"log"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) SelectAllZone() []models.Zone {
	var list []models.Zone

	qstr := `
		SELECT z.id, z.name, z.floors_number, z.color, campus.name
		FROM zone z
		JOIN campus on campus.id = z.id_campus
	`
	rows, err := db.getConnection().Query(qstr)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		z := models.Zone{}

		err := rows.Scan(
			&z.ID,
			&z.Name,
			&z.FloorNumbers,
			&z.Color,
			&z.NameCampus,
		)
		if err != nil {
			log.Fatal(err)
		}
		list = append(list, z)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return list
}
func (db *dbConfig) SelectZone(id string) (models.Zone, error) {
	var z models.Zone

	stmt, err := db.getConnection().Prepare(`
		SELECT id, name, floors_number, color, id_campus
		FROM zone
		WHERE id = $1
	`)
	if err != nil {
		log.Println(err)
		return z, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(
		&z.ID,
		&z.Name,
		&z.FloorNumbers,
		&z.Color,
		&z.IDCampus,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("[ no zone with the provided credentials ]")
			return z, err
		}
		log.Println(err)
		return z, err
	}

	return z, nil
}
