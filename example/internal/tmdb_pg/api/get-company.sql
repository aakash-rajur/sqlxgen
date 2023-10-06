select
c."id" as "id",
c."name" as "name",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.movie_id,
        'title', m.title,
        'releaseDate', m.release_date
      ) order by m.release_date desc
    )
    from movies_companies mc
    inner join movies m on mc.movie_id = m.id
    where true
    and mc.company_id = c.id
  ),
  '[]'
) as "movies"
from companies c
where c.id = :id; -- :id type: bigint
