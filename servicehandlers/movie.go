package servicehandlers

import (
	"bookmyshow/dao"
	"bookmyshow/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// MovieHandler : Restful Movie Handler
type MovieHandler struct {
}

func (p MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	methodRouter(p, w, r)
}

// Get : Restful Get Method
func (p MovieHandler) Get(r *http.Request) (string, int) {

	q := r.URL.Query()
	movieID := q.Get("movie_id")
	movieName := q.Get("name")

	if movieID != "" && movieName != "" {
		return "Bad Request", 400

	} else if movieID != "" {
		movieID := utils.StringToInteger(movieID)
		movie := dao.GetMovieByID(movieID)
		return utils.StructToJSONString(movie), 200

	} else if movieName != "" {
		movie := dao.GetMoviesByName(movieName)
		return utils.StructToJSONString(movie), 200
	} else {
		movie := dao.GetAllMovies()
		return utils.StructToJSONString(movie), 200
	}
}

// Put : Restful Put Method
func (p MovieHandler) Put(r *http.Request) (string, int) {
	return "PUT Called", 200
}

// Post : Restful Post Method
func (p MovieHandler) Post(r *http.Request) (string, int) {

	var movie dao.Movie
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return "Bad Request", 400
	}

	dao.CreateMovie(movie.Name, movie.DateOfRelease)

	movie = dao.GetMoviesByName(movie.Name)
	fmt.Println(movie)
	fmt.Println(movie.ID, strconv.Itoa(movie.ID))

	return "Movie ID: " + strconv.Itoa(movie.ID), 200
}
