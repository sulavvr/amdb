USING PERIODIC COMMIT 1000
LOAD CSV WITH HEADERS FROM "file:///movies_metadata.csv" AS line
CREATE (m:Movie {id: toInteger(line.id), imdbId: toInteger(line.imdb_id), title: line.original_title, adult: toBoolean(line.adult), budget: line.budget, language: line.original_language, overview: line.overview, popularity: line.popularity, runtime: line.runtime, releaseDate: line.release_date, revenue: line.revenue, tagline: line.tagline, voteAverage: line.vote_average, voteCount: line.vote_count, poster: line.poster_path })
    WITH m, line, split(split(apoc.convert.toString(line.genres), "[")[1], "]") AS s
    WITH m, line, apoc.convert.fromJsonMap(s[0]) AS genre
    WITH m,  genre, line, split(split(apoc.convert.toString(line.production_companies), "[")[1], "]") as pcomp
    WITH m,  genre, line, apoc.convert.fromJsonMap(pcomp[0]) AS pcomps
    WITH m,  genre,  pcomps, split(split(apoc.convert.toString(line.production_countries), "[")[1], "]") as pcon
    WITH m,  genre,  pcomps, apoc.convert.fromJsonMap(pcon[0]) AS pcons
FOREACH(ignoreMe IN CASE WHEN apoc.convert.toBoolean(genre) <> false THEN [1] ELSE [] END |
    MERGE (g:Genre {id: genre.id}) ON CREATE SET g.name = genre.name
    MERGE (m)-[:TYPE_OF]->(g)
)
FOREACH(ignoreMe IN CASE WHEN apoc.convert.toBoolean(pcomps) <> false THEN [1] ELSE [] END |
    MERGE (p:Production {id: pcomps.id}) ON CREATE set p.name = pcomps.name
    MERGE (m)-[:PRODUCED_BY]->(p)
)
FOREACH(ignoreMe IN CASE WHEN apoc.convert.toBoolean(pcons) <> false THEN [1] ELSE [] END |
    MERGE (c:Country {id: pcons.iso_3166_1}) ON CREATE set c.name = pcons.name
    MERGE (m)-[:PRODUCED_BY_COUNTRY]->(c)
)

// Update movie node with poster path
USING PERIODIC COMMIT 1000
LOAD CSV WITH HEADERS FROM "file:///movies_metadata.csv" AS line
MATCH (m:Movie {id: toInteger(line.id) }) SET m.poster = line.poster_path;

// Create cast nodes and add relationship to movie
LOAD CSV WITH HEADERS FROM "file:///credits.csv" AS line
WITH line LIMIT 45504
MATCH (m:Movie {id: toInteger(line.id)})
WITH m, apoc.convert.fromJsonList(line.cast) AS casts
FOREACH(ignoreMe IN CASE WHEN apoc.convert.toBoolean(m) <> false THEN [1] ELSE [] END |
    FOREACH(ignoreMe IN CASE WHEN apoc.convert.toBoolean(casts) <> false THEN [1] ELSE [] END |
        FOREACH (cast in casts |
            MERGE (c:Cast{id: toInteger(cast.id)}) ON CREATE
            SET c.name = cast.name,
            c.gender = cast.gender,
            c.profile = cast.profile_path
            MERGE (c)-[:ACTS_IN{character: cast.character, credit_id: cast.credit_id, cast_id: cast.cast_id}]->(m)
        )
    )
)
