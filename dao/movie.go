package dao

import (
	"fmt"
)

// Movie : The Movie Table Schema
type Movie struct {
	ID            int
	Name          string
	DateOfRelease string
	Rating        float32
}

// // MovieBody : The expected Json format for input from user
// type MovieBody struct {
// 	Name          string
// 	DateOfRelease string
// }

// GetAllMovies : Returns all movies form MoviesDB
func GetAllMovies() []Movie {
	var movies []Movie

	rows, err := db.Query("Select * from movies")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var movie Movie

		err := rows.Scan(
			&movie.ID,
			&movie.Name,
			&movie.DateOfRelease,
			&movie.Rating)

		if err != nil {
			panic(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

// GetMovieByID : Returns movie associated with movieID
func GetMovieByID(movieID int) Movie {
	var movie Movie

	query := "Select * from movies where movie_id=$1"

	err := db.QueryRow(query, movieID).Scan(
		&movie.ID,
		&movie.Name,
		&movie.DateOfRelease,
		&movie.Rating)

	if err != nil {
		return Movie{}
	}

	return movie
}

// GetMoviesByName : Returns movie associated with movieName
func GetMoviesByName(name string) Movie {
	var movie Movie

	query := "Select * from movies where name=$1"

	err := db.QueryRow(query, name).Scan(
		&movie.ID,
		&movie.Name,
		&movie.DateOfRelease,
		&movie.Rating)

	if err != nil {
		fmt.Print(err)
		return Movie{}
	}

	return movie
}

// CreateMovie : Inserts Data into Movie database
func CreateMovie(name string, dateOfRelease string) {

	query := `Insert into movies (name, date_of_release) 
			values ($1, $2)`
	db.QueryRow(query, name, dateOfRelease)
}
