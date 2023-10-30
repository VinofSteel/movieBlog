package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/movieBlog/internal/environment"
	_"github.com/lib/pq"
)

const yellow = "\x1b[33m"
const green = "\x1b[32m"
const reset = "\x1b[0m"

func DatabaseConnection() string {
	user, password, host, port, databaseName := environment.EnvVariable("PGUSER"), environment.EnvVariable("PGPASSWORD"), environment.EnvVariable("PGHOST"), environment.EnvVariable("PGPORT"), environment.EnvVariable("PGDATABASE")
	connectionString := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, databaseName)

	fmt.Printf(yellow+"Opening connection with database %v on port %v..."+reset+"\n", databaseName, port)	
	db, err := sql.Open("postgres", connectionString)
	fmt.Printf(green + "Connection opened!" + reset + "\n")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	
	return ""
}