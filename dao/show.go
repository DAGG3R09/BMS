package dao

import (
	"log"
	"time"
)

// Show : Show Table Schema
type Show struct {
	ID           int
	MovieID      int
	MultiplexID  int
	Showtime     time.Time
	ScreenNumber int
}

// GetAllShows : Returns all Shows form ShowsDB
func GetAllShows() []Show {
	var shows []Show

	rows, err := db.Query("Select * from Shows")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var show Show

		err := rows.Scan(
			&show.ID,
			&show.MovieID,
			&show.MultiplexID,
			&show.ScreenNumber,
			&show.Showtime)

		if err != nil {
			panic(err)
		}
		shows = append(shows, show)
	}

	return shows
}

// GetAllShowsByMovie : Returns all Shows form ShowsDB
// // by movieID
func GetAllShowsByMovie(movieID int) []Show {
	var shows []Show

	rows, err := db.Query(`Select * from Shows 
		where movie_id=$1`, movieID)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var show Show

		err := rows.Scan(
			&show.ID,
			&show.MovieID,
			&show.MultiplexID,
			&show.ScreenNumber,
			&show.Showtime)

		if err != nil {
			panic(err)
		}
		shows = append(shows, show)
	}

	return shows
}

// GetAllShowsByMultiplex : Returns all Shows form ShowsDB
// by multiplexID
func GetAllShowsByMultiplex(multiplexID int) []Show {
	var shows []Show

	rows, err := db.Query(`Select * from Shows where
		 multiplex_id=$1`, multiplexID)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var show Show

		err := rows.Scan(
			&show.ID,
			&show.MovieID,
			&show.MultiplexID,
			&show.ScreenNumber,
			&show.Showtime)

		if err != nil {
			panic(err)
		}
		shows = append(shows, show)
	}

	return shows
}

// GetSingleShow : Returns Show associated with ShowID
func GetSingleShow(showID int) Show {
	var show Show

	query := "Select * from users where Show_id=$1"

	err := db.QueryRow(query, showID).Scan(
		&show.ID,
		&show.MovieID,
		&show.MultiplexID,
		&show.ScreenNumber,
		&show.Showtime)

	if err != nil {
		return Show{}
	}

	return show
}

// CreateShow : Inserts Data into Show database
func CreateShow(movieID int, multiplexID int,
	showtime time.Time, screenNo int) {

	query := `Insert into Shows 
	(show_id, movie_id, multiplex_id, showtime, screen_no) 
	values ($1, $2, $3, $4)`

	err := db.QueryRow(query, movieID, multiplexID, showtime, screenNo)

	if err != nil {
		log.Fatal("Error in inserting Show", err)
	}
}
