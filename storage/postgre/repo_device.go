package postgre

import (
	"database/sql"
	"log"
	"time"

	"github.com/luispfcanales/inventory-oti/models"
)

func (db *dbConfig) InsertDevice(d models.Device) (models.Device, error) {
	str := `
		INSERT INTO device(patrimonial_code,serial_code,brand,model_name,
			id_state_device,
			id_type_device,
			id_dependency,
			adquisition_date,
			more_info
		)
		VALUES ($1,$2,$3,$4,
			$5,
			$6,
			$7,
			$8,
			$9
		)
	`
	_, err := db.getConnection().Exec(
		str,
		d.PatrimonialCode, d.SerialCode, d.Brand, d.ModelName,
		d.StateDevice,
		d.TypeDevice,
		d.DependencyDevice,
		d.AdquisitionDate,
		d.MoreInfo,
	)
	if err != nil {
		return models.Device{}, err
	}

	return d, nil
}

func (db *dbConfig) SelectAllDevice() []models.Device {
	var list []models.Device
	str := `
		SELECT d.patrimonial_code , d.serial_code , d.brand , d.model_name,
		sd.description as state,
		td.name as name_type_device,
		dependency.abbreviation,
		d.adquisition_date,
		d.more_info 
		FROM device d
		JOIN state_device sd ON sd.id = d.id_state_device
		JOIN type_device td ON td.id = d.id_type_device 
		JOIN dependency ON dependency.id = d.id_dependency
	`
	rows, err := db.getConnection().Query(str)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var date sql.NullTime
		r := models.Device{}

		err := rows.Scan(
			&r.PatrimonialCode,
			&r.SerialCode,
			&r.Brand,
			&r.ModelName,
			&r.StateDevice,
			&r.TypeDevice,
			&r.DependencyDevice,
			&date,
			&r.MoreInfo,
		)
		if err != nil {
			log.Fatal(err)
		}

		value, err := date.Value()
		if err != nil {
			log.Println(err)
		}

		v, ok := value.(time.Time)
		if !ok {
			r.AdquisitionDate = "0000-00-00"
			list = append(list, r)
		} else {
			r.AdquisitionDate = v.Format("2006-01-02")
			list = append(list, r)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return list
}
