package database

import "database/sql"

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