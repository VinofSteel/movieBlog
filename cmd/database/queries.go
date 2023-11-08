package database

import (
	"database/sql"
	"fmt"
	"log"
)

func GetMovies(db *sql.DB) ([]Movie, error) {
	fmt.Println("Getting movies in DB...")
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

func InsertMovie(db *sql.DB, movie Movie) string {
	fmt.Printf("Inserting movie %v in DB...\n", movie.Title)

	query := `INSERT INTO movies (title, director)
		VALUES ($1, $2) RETURNING id`

	var id string
	err := db.QueryRow(query, movie.Title, movie.Director).Scan(&id)
	if err != nil {
		log.Fatal(err)
	}

	return id
}