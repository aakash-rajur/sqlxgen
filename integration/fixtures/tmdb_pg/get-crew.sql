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
        'releaseDate', m.release_date,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by m.release_date desc
    )
    from movies_crew mc
    inner join movies m on mc.movie_id = m.id
    inner join hyper_parameters d on (
      true
      and d.type = 'department'
      and d.value = mc.department_id
    )
    inner join hyper_parameters j on (
      true
      and j.type = 'job'
      and j.value = mc.job_id
    )
    where true
    and mc.crew_id = c.id
  ),
  '[]'
) as "movies"
from crew c
where c.id = :id; -- :id type: bigint
