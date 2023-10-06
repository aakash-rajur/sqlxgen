select
c.id as "id",
c.name as "name",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date
      )
    )
    from movies_companies mc
    inner join movies m on mc.movie_id = m.id
    where true
    and mc.company_id = c.id
    order by m.release_date desc
  ),
  '[]'
) as "movies"
from companies c
where c.id = :id; -- :id type: bigint
