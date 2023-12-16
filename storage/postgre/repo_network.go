package postgre

import (
	"log"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) SelectNetworks() []models.Network {
	var list []models.Network
	str := `
		SELECT n.id, n.ip_address, nc.id, nc.type_connection
		FROM network n
		JOIN network_connection nc ON nc.id = n.type_connection
	`
	rows, err := db.getConnection().Query(str)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		n := models.Network{}
		err := rows.Scan(
			&n.ID,
			&n.IpAddress,
			&n.NetConn.ID,
			&n.NetConn.TypeConnection,
		)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, n)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}

func (db *dbConfig) SelectResumeNetworks() []models.ResumeNetworks {
	var list []models.ResumeNetworks
	str := `
		SELECT campus.id as id_campus,campus.name as campus_name, dl.floor as number_floor,
		nt.ip_address,
		zone.name as zone_name,
		dependency.name as dependency_name,
		network_connection.type_connection
		FROM dependency_location dl
		JOIN network nt ON dl.id_network = nt.id
		JOIN network_connection ON network_connection.id = nt.type_connection
		JOIN zone ON dl.id_zone = zone.id
		JOIN campus ON campus.id = zone.id_campus
		JOIN dependency ON dependency.id_location = dl.id;
	`
	rows, err := db.getConnection().Query(str)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		r := models.ResumeNetworks{}
		err := rows.Scan(
			&r.IdCampus,
			&r.CampusName,
			&r.NumberFloor,
			&r.IpAddress,
			&r.ZoneName,
			&r.DependendyName,
			&r.TypeConnection,
		)
		if err != nil {
			log.Fatal(err)
		}

		list = append(list, r)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}
