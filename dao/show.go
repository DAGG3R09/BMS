package dao

import (
	"log"
	"time"
)

// Show : Show Table Schema
type Show struct {
	ID            int
	MovieID       int       `json:"movie_id"`
	MultiplexID   int       `json:"multiplex_id"`
	ShowDate      time.Time `json:"show_date"`
	Showtime      time.Time `json:"show_time"`
	ScreenNumber  int       `json:"screen_no"`
	MovieName     string    `json:"movie_name"`
	MultiplexName string    `json:"multiplex_name"`
}

// GetAllShows : Returns all Shows form ShowsDB
func GetAllShows() []Show {
	var shows []Show

	rows, err := db.Query("Select * from Show")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var show Show

		err := rows.Scan(
			&show.ID,
			&show.MovieID,
			&show.MultiplexID,
			&show.Showtime,
			&show.ShowDate,
			&show.ScreenNumber,
			&show.MovieName,
			&show.MultiplexName)

		if err != nil {
			panic(err)
		}
		show.ShowDate.Format("2006-01-02")
		shows = append(shows, show)
	}

	return shows
}

// GetShowsByMovieID : Returns all Shows form ShowsDB
// // by movieID
func GetShowsByMovieID(movieID int) []Show {
	var shows []Show

	rows, err := db.Query(`Select * from Show
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
			&show.ShowDate,
			&show.Showtime,
			&show.ScreenNumber,
			&show.MovieName,
			&show.MultiplexName)

		if err != nil {
			panic(err)
		}
		show.ShowDate.Format("2006-01-02")
		shows = append(shows, show)
	}

	return shows
}

// GetShowsByMultiplexID : Returns all Shows form ShowsDB
// by multiplexID
func GetShowsByMultiplexID(multiplexID int) []Show {
	var shows []Show

	rows, err := db.Query(`Select * from Show where
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
			&show.Showtime,
			&show.ShowDate,
			&show.ScreenNumber,
			&show.MovieName,
			&show.MultiplexName)

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

	query := "Select * from show where Show_id=$1"

	err := db.QueryRow(query, showID).Scan(
		&show.ID,
		&show.MovieID,
		&show.MultiplexID,
		&show.Showtime,
		&show.ShowDate,
		&show.ScreenNumber,
		&show.MovieName,
		&show.MultiplexName)

	if err != nil {
		return Show{}
	}

	show.ShowDate.Format("2006-01-02")

	return show
}

// CreateShow : Inserts Data into Show database
func CreateShow(movieID int, multiplexID int,
	showtime time.Time, screenNo int) {

	query := `Insert into Show
	(show_id, movie_id, multiplex_id,showdate, showtime, screen_no) 
	values ($1, $2, $3, $4)`

	err := db.QueryRow(query, movieID, multiplexID, showtime, screenNo)

	if err != nil {
		log.Fatal("Error in inserting Show", err)
	}
}
