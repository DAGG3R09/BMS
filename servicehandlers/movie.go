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
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p MovieHandler) Get(r *http.Request) SrvcRes {

	q := r.URL.Query()
	movieID := q.Get("movie_id")
	movieName := q.Get("name")

	if movieID != "" && movieName != "" {
		return SimpleBadRequest("Bad Request")

	} else if movieID != "" {
		movieID := utils.StringToInteger(movieID)
		movie := dao.GetMovieByID(movieID)
		return Simple200OK(utils.StructToJSONString(movie))

	} else if movieName != "" {
		movie := dao.GetMoviesByName(movieName)
		return Simple200OK(utils.StructToJSONString(movie))
	} else {
		movie := dao.GetAllMovies()
		return Simple200OK(utils.StructToJSONString(movie))
	}
}

// Put : Restful Put Method
func (p MovieHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Method
func (p MovieHandler) Post(r *http.Request) SrvcRes {

	var movie dao.Movie
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&movie)

	if err != nil {
		// log.Fatal(err)
		fmt.Println(err)
		return SimpleBadRequest("Bad Request")
	}

	dao.CreateMovie(movie.Name, movie.DateOfRelease)

	movie = dao.GetMoviesByName(movie.Name)
	fmt.Println(movie)
	fmt.Println(movie.ID, strconv.Itoa(movie.ID))

	return Simple200OK("Movie ID: " + strconv.Itoa(movie.ID))
}
