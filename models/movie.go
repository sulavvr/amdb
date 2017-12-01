package models

import (
	"database/sql"
	"log"
)

/**
 * Movie model contains single movie
 */
type Movie struct {
	id          int
	title       string
	tagline     sql.NullString
	overview    sql.NullString
	rating      float32
	votes       int
	releaseDate string
	runtime     int
	language    string
	revenue     int
	popularity  float32
	adult       bool
	budget      string
	poster      sql.NullString
	genre       string
	genreId     int
	country     string
	production  string
	casts       []Cast
}

func (movie Movie) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":          movie.id,
		"title":       movie.title,
		"tagline":     movie.tagline,
		"overview":    movie.overview,
		"rating":      movie.rating,
		"votes":       movie.votes,
		"releaseDate": movie.releaseDate,
		"runtime":     movie.runtime,
		"language":    movie.language,
		"revenue":     movie.revenue,
		"popularity":  movie.popularity,
		"adult":       movie.adult,
		"budget":      movie.budget,
		"poster":      movie.poster,
		"genre":       movie.genre,
		"genreId":     movie.genreId,
		"production":  movie.production,
		"country":     movie.country,
		"casts":       movie.casts,
	}

	return details
}

func (movie Movie) Count(db *sql.DB) int {
	query := `MATCH (m:Movie) RETURN COUNT(*)`
	var count int
	db.QueryRow(query).Scan(&count)

	return count
}

func (movie Movie) All(skip int, limit int, db *sql.DB) []Movie {
	query := `MATCH (m:Movie)
			OPTIONAL MATCH (m)-[:TYPE_OF]->(g) WITH m, g
            RETURN m.id, m.title, m.tagline, m.overview, m.voteAverage, m.voteCount, m.releaseDate,
            m.runtime, m.language, m.revenue, m.popularity, m.adult, m.budget, m.poster, g.name, g.id
            SKIP {0} LIMIT {1}`
	rows, err := db.Query(query, skip, limit)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var movies []Movie

	for rows.Next() {
		m := movie
		rows.Scan(&m.id, &m.title, &m.tagline, &m.overview, &m.rating, &m.votes, &m.releaseDate, &m.runtime, &m.language, &m.revenue, &m.popularity, &m.adult, &m.budget, &m.poster, &m.genre, &m.genreId)
		movies = append(movies, m)
	}

	return movies
}

func (movie Movie) Find(id int, db *sql.DB) Movie {
	query := `MATCH (m:Movie)
			OPTIONAL MATCH (m)-[:TYPE_OF]->(g) WITH m, g
			OPTIONAL MATCH (m)-[:PRODUCED_BY]->(p) WITH m, g, p
			OPTIONAL MATCH (m)-[:PRODUCED_BY_COUNTRY]->(c) WITH m, g, p, c
			WHERE m.id = {0} RETURN m.id, m.title, m.tagline, m.overview, m.voteAverage, m.voteCount, m.releaseDate, m.runtime, m.language, m.revenue, m.popularity, m.adult, m.budget, m.poster, g.name, g.id, p.name, c.id`
	row := db.QueryRow(query, id)

	row.Scan(&movie.id, &movie.title, &movie.tagline, &movie.overview, &movie.rating, &movie.votes, &movie.releaseDate, &movie.runtime, &movie.language, &movie.revenue, &movie.popularity, &movie.adult, &movie.budget, &movie.poster, &movie.genre, &movie.genreId, &movie.production, &movie.country)

	movie.casts = new(Cast).ForMovie(id, db)
	return movie
}
