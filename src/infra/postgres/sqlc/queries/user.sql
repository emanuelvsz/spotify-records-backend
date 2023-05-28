select a.id,
        a.name,
        g.id,
        g.name
    from artist a 
    inner join artist_genre ag on a.id = ag.artist_id
    inner join genre g on g.id = ag.genre_id;