package models

import (
	"database/sql"
	"log"
)

type Genre struct {
	id   int
	name string
}

func (genre Genre) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":   genre.id,
		"name": genre.name,
	}

	return details
}

func (genre Genre) All(db *sql.DB) []Genre {
	query := `MATCH (g:Genre) RETURN g.id, g.name ORDER BY g.name`

	rows, err := db.Query(query)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var genres []Genre

	for rows.Next() {
		rows.Scan(&genre.id, &genre.name)
		genres = append(genres, genre)
	}

	return genres
}

func (genre Genre) Movies(id, skip, limit int, db *sql.DB) []Movie {
	query := `MATCH (g:Genre)<-[:TYPE_OF]-(m) WHERE g.id = {0} RETURN m.id, m.title, m.tagline, m.overview, m.voteAverage, m.voteCount, m.releaseDate, m.runtime, m.language, m.revenue, m.popularity, m.adult, m.budget, m.poster SKIP {1} LIMIT {2}`

	rows, err := db.Query(query, id, skip, limit)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		m := Movie{}
		rows.Scan(&m.id, &m.title, &m.tagline, &m.overview, &m.rating, &m.votes, &m.releaseDate, &m.runtime, &m.language, &m.revenue, &m.popularity, &m.adult, &m.budget, &m.poster)
		movies = append(movies, m)
	}

	return movies
}

func (genre Genre) Count(id int, db *sql.DB) int {
	query := `MATCH (g:Genre)<-[:TYPE_OF]-(m) WHERE g.id = {0} RETURN COUNT(*)`
	var count int
	db.QueryRow(query, id).Scan(&count)

	return count
}
