package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/movieBlog/internal/environment"
)

type Movie struct {
	Title       string
	Director    string
}

const yellow = "\x1b[33m"
const green = "\x1b[32m"
const reset = "\x1b[0m"

var ConnectionString string

func createAppTables(db *sql.DB) {
	func() { //create uuid extension
		extensionQuery := "CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";"

		_, err := db.Exec(extensionQuery)
		if err != nil {
			log.Fatal(err)
		}
	}()

	func() { //creating movie table if it does not exist
		query := `
			CREATE TABLE IF NOT EXISTS movies (
				id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
				title VARCHAR NOT NULL,
				director VARCHAR NOT NULL,
				description TEXT,
				deleted BOOLEAN DEFAULT false,
				created_at TIMESTAMP DEFAULT NOW()
			);
		`
	
		_, err := db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func InitializeDB() {
	user, password, host, port, databaseName := environment.EnvVariable("PGUSER"), environment.EnvVariable("PGPASSWORD"), environment.EnvVariable("PGHOST"), environment.EnvVariable("PGPORT"), environment.EnvVariable("PGDATABASE")
	ConnectionString = fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, host, port, databaseName)

	fmt.Printf(yellow+"Opening connection with database %v on port %v..."+reset+"\n", databaseName, port)	
	db, err := sql.Open("postgres", ConnectionString)
	fmt.Printf(green + "Connection opened!" + reset + "\n")

	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	createAppTables(db)
}

func GetMovies(db *sql.DB) ([]Movie, error) {
	query := "SELECT title, director FROM movies WHERE deleted = false;"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []Movie
	for rows.Next() {
		var movie Movie
		if err := rows.Scan(&movie.Title, &movie.Director); err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return movies, nil
}