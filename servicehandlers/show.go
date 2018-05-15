package servicehandlers

import (
	"bookmyshow/dao"
	"bookmyshow/utils"
	"net/http"
)

// ShowHandler : Restful ShowHandler Method
type ShowHandler struct {
}

func (p ShowHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := methodRouter(p, w, r)
	response.(SrvcRes).RenderResponse(w)
}

// Get : Restful Get Method
func (p ShowHandler) Get(r *http.Request) SrvcRes {
	q := r.URL.Query()
	movieID := q.Get("movie_id")
	multiplexID := q.Get("multiplex_id")

	if movieID == "" && multiplexID == "" {
		return SimpleBadRequest("Bad Request")
	} else if movieID != "" {
		movieID := utils.StringToInteger(movieID)
		shows := dao.GetShowsByMovieID(movieID)
		return Response200OK(shows)

	} else if multiplexID != "" {
		multiplexID := utils.StringToInteger(multiplexID)
		shows := dao.GetShowsByMultiplexID(multiplexID)
		return Response200OK(shows)
	} else {
		return SimpleBadRequest("Bad Request")
	}
}

// Put : Restful Put Method
func (p ShowHandler) Put(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}

// Post : Restful Post Handler
func (p ShowHandler) Post(r *http.Request) SrvcRes {
	return ResponseNotImplemented()
}
