package postgre

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once
var cnn *sql.DB

type dbConfig struct {
	host     string
	port     int
	user     string
	password string
	dbname   string
}

func NewPostgreStorage() *dbConfig {
	db := &dbConfig{}
	loadEnv(db)
	return db
}

func (db *dbConfig) getConnection() *sql.DB {
	str := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		db.host, db.port, db.user, db.password, db.dbname,
	)
	once.Do(func() {
		cn, err := sql.Open("postgres", str)
		if err != nil {
			log.Fatal(err)
		}
		cnn = cn
	})

	return cnn
}

func loadEnv(db *dbConfig) {
	db.host = "localhost"
	db.user = "postgres"
	db.password = "luis"
	db.dbname = "oti"
	db.port = 80
}
