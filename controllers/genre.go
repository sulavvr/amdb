package controllers

import (
	// "log"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/sulavvr/neogo/models"
)

type Genre struct {
	model models.Genre
}

func (genre Genre) Index(writer http.ResponseWriter, req *http.Request) {
	genres := genre.model.All(DB)

	data["genres"] = genres

	err := Templates.ExecuteTemplate(writer, "genres.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}

func (genre Genre) Show(writer http.ResponseWriter, req *http.Request) {
	path := strings.Split(req.URL.Path, "/")
	id, _ := strconv.Atoi(path[len(path)-1])

	_url := strings.Split(req.RequestURI, "?")
	q, _ := url.ParseQuery(_url[len(_url)-1])
	skip := 0

	if q.Get("skip") != "" {
		skip, _ = strconv.Atoi(q["skip"][0])
	}

	limit := 24
	movies := genre.model.Movies(id, skip, limit, DB)
	total := genre.model.Count(id, DB)
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

	data["pagination"] = paginate(start, end, total, limit, fmt.Sprintf("genre/%d", id))
	err := Templates.ExecuteTemplate(writer, "movies.html", data)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
	}
}
