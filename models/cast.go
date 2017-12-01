package models

import (
	"database/sql"
	"log"
)

type Cast struct {
	id        int
	gender    int
	name      string
	profile   string
	character string
	castId    int
	creditId  int
}

func (cast Cast) GetInfo() map[string]interface{} {
	details := map[string]interface{}{
		"id":        cast.id,
		"gender":    cast.gender,
		"name":      cast.name,
		"profile":   cast.profile,
		"character": cast.character,
		"castId":    cast.castId,
		"creditId":  cast.creditId,
	}

	return details
}

func (cast Cast) ForMovie(id int, db *sql.DB) []Cast {
	query := `MATCH (c:Cast) OPTIONAL MATCH (c)-[r:ACTS_IN]->(m:Movie) WITH c, m, r WHERE m.id = {0} RETURN c.id, c.gender, c.name, c.profile, r.character, r.castId, r.creditId`

	rows, err := db.Query(query, id)

	if err != nil {
		log.Fatal(err.Error())
	}
	defer rows.Close()

	var casts []Cast
	for rows.Next() {
		c := cast
		rows.Scan(&c.id, &c.gender, &c.name, &c.profile, &c.character, &c.castId, &c.creditId)
		casts = append(casts, c)
	}

	return casts
}
