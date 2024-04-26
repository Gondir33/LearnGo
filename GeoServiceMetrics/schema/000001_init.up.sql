CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE IF NOT EXISTS search_history (
    id SERIAL PRIMARY KEY,
    search_request VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS addresses (
    id SERIAL PRIMARY KEY,
    lat VARCHAR(100),
    lon VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS history_search_address (
    id SERIAL PRIMARY KEY,
    id_search_history INT NOT NULL REFERENCES search_history (id),
    id_address INT NOT NULL REFERENCES addresses (id),
    CONSTRAINT id_address_unique UNIQUE (id_search_history, id_address)
);


--INSERT INTO search_history (search_request) VALUES ($1) RETURNING id
--INSERT INTO addresses (lat, lon) VALUES ($1,$2) RETURNING id
--INSERT INTO history_search_address (id_search_history, id_address) VALUES ($1,$2)


--SELECT id FROM search_history WHERE similarity(search_request, $1) > 0.7 ORDER BY similarity(search_request, $1) DESC
--SELECT id_address FROM history_search_address WHERE id_search_history = $1
--SELECT (lat, lon) FROM addresses WHERE id = $1