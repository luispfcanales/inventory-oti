package postgre

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("[ Error loading .env file]")
	}
	db.host = os.Getenv("HOST_DB")
	db.user = os.Getenv("USER_DB")
	db.password = os.Getenv("PASSWORD_DB")
	db.dbname = os.Getenv("DBNAME_DB")

	port, err := strconv.Atoi(os.Getenv("PORT_DB"))
	if err != nil {
		log.Fatal("[ Error loading PORT_DB .env file]")
	}
	db.port = port
}
