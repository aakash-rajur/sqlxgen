select
a."id" as "id",
a."name" as "name",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', ma.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'character', ma.character
      ) order by m.release_date desc
    )
    from movies_actors ma
    inner join movies m on ma.movie_id = m.id
    where true
    and ma.actor_id = a.id
  ),
  '[]'
) as "movies"
from actors a
where a.id = :id; -- :id type: bigint
