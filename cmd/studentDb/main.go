package main

import (
	"database/sql"
	
	"log"

	"github.com/briMus-88/deer/pkg/studentDb"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	var dbConnection *sql.DB
	connStr := `host=localhost port=5432 user=postgres password=2811 dbname=studentGolangTraining sslmode=disable`

	dbConnection, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConnection.Ping()
	if err != nil {
		log.Fatal(err)
	}

	store := studentDb.NewStudentStore(dbConnection)

//	store.DisplayIntro()

	store.MainMenuDisplay()

	defer dbConnection.Close()
}
