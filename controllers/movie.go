package controllers

import (
	// "log"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sulavvr/neogo/models"
)

type Movie struct {
	model models.Movie
}

func (movie Movie) Index(writer http.ResponseWriter, req *http.Request) {
	_url := strings.Split(req.RequestURI, "/?")
	q, _ := url.ParseQuery(_url[len(_url)-1])
	skip := 0

	if q.Get("skip") != "" {
		skip, _ = strconv.Atoi(q["skip"][0])
	}

	limit := 24
	movies := movie.model.All(skip, limit, DB)
	total := movie.model.Count(DB)
	pages := int(math.Ceil(float64(total / limit)))

	data["movies"] = movies
	data["skip"] = skip
	start := (skip / limit) - 1
	if skip == 0 || start == 0 {
		start = 1
	}

	end := start + 10
	if end > pages {
		end = pages
	}

	data["pagination"] = paginate(start, end, total, limit, "movies/")
	err := Templates.ExecuteTemplate(writer, "movies.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (movie Movie) Show(writer http.ResponseWriter, req *http.Request) {
	url := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(url[len(url)-1])

	_movie := movie.model.Find(id, DB) // get movie

	data["movie"] = _movie

	err := Templates.ExecuteTemplate(writer, "movie.html", data)

	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
