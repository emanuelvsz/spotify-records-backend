SELECT a.id AS id,
       a.name AS name,
       a.super_artist_id AS super_artist_id,
       a.description AS description,
       a.founded_at AS founded_at,
       a.terminated_at AS terminated_at,
       g.name AS genre_name,
       g.description AS genre_description,
       g.created_at AS genre_created_at
FROM artist a
INNER JOIN artist_genre ag ON a.id = ag.artist_id
INNER JOIN genre g ON g.id = ag.genre_id
WHERE a.id = 'cfd8b073-a303-4886-836a-f249e88be9bc';
