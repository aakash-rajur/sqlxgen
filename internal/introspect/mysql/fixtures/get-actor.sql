select
a.id as "id",
a.name as "name",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', ma.movie_id,
        'title', m.title,
        'releaseDate', m.release_date,
        'character', ma.cast
      )
    )
    from movies_actors ma
    inner join movies m on ma.movie_id = m.id
    where true
    and ma.actor_id = a.id
    order by m.release_date desc
  ),
  cast('[]' as json)
) as "movies"
from actors a
where a.id = :id; -- :id type: bigint
