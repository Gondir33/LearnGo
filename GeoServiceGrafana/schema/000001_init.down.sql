DROP TABLE IF EXISTS history_search_address CASCADE;
DROP TABLE IF EXISTS search_history CASCADE;
DROP TABLE IF EXISTS addresses CASCADE;

--INSERT INTO search_history (search_request) VALUES ($1) RETURNING id
--INSERT INTO addresses (lat, lon) VALUES ($1,$2) RETURNING id
--INSERT INTO history_search_address (id_search_history, id_address) VALUES ($1,$2)


--SELECT id FROM search_history WHERE similarity(search_request, $1) > 0.7 ORDER BY similarity(search_request, $1) DESC
--SELECT id_address FROM history_search_address WHERE id_search_history = $1
--SELECT (lat, lon) FROM addresses WHERE id = $1