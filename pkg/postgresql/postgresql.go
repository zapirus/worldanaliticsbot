package postgresql

import (
	"Projects/WordAnalytics/pkg/logger"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "telebot_db"
)

type Data struct {
	Id        int
	Url       string
	Info      string
	CreatedAt string
}

func Connect() (*sql.DB, error) {
	log := logger.GetLogger()

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Errorf("Opening postgres failed: %e", err)
	}

	if err = db.Ping(); err != nil {
		log.Errorf("Opening postgres failed: %e", err)
	}

	log.Info("connection successful")

	return db, nil
}

func Insert(url string, objects []byte, db *sql.DB) {
	log := logger.GetLogger()

	sqlstatement := `INSERT INTO sites (url , words)
                   VALUES ($1, $2)`

	_, err := db.Exec(sqlstatement, url, objects)
	if err != nil {
		log.Errorf("Can't insert to table sites: %e", err)
	}
	log.Info("Inserting is successful")
}

func Select(db *sql.DB) []byte {
	log := logger.GetLogger()

	rows, err := db.Query("SELECT words FROM sites;")
	if err != nil {
		log.Fatalf("Can't select from table words: %e", err)
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Errorf("failed to close rows: %e", err)
		}
	}(rows)

	for rows.Next() {
		var title []byte
		if err = rows.Scan(&title); err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(title))
		return title
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}
