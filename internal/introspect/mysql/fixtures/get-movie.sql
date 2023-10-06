select
m.id as "id",
m.title as "title",
m.original_title as "originalTitle",
m.original_language as "originalLanguage",
m.overview as "overview",
m.runtime as "runtime",
m.release_date as "releaseDate",
m.tagline as "tagline",
m.status as "status",
m.homepage as "homepage",
m.popularity as "popularity",
m.vote_average as "voteAverage",
m.vote_count as "voteCount",
m.budget as "budget",
m.revenue as "revenue",
m.keywords as "keywords",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      )
    )
    from movies_genres g
    inner join hyper_parameters hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where true
    and g.movie_id = m.id
    order by hp.friendly_name
  ),
  cast('[]' as json)
) as "genres",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', c.country_id,
        'name', hp.friendly_name
      )
    )
    from movies_countries c
    inner join hyper_parameters hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where true
    and c.movie_id = m.id
    order by hp.friendly_name
  ),
  cast('[]' as json)
) as "countries",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', l.language_id,
        'name', hp.friendly_name
      )
    )
    from movies_languages l
    inner join hyper_parameters hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where true
    and l.movie_id = m.id
    order by hp.friendly_name
  ),
  cast('[]' as json)
) as "languages",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.company_id,
        'name', c.name
      )
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where true
    and mc.movie_id = m.id
    order by c.name
  ),
  cast('[]' as json)
) as "companies",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.cast,
        'order', ma.cast_order
      )
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where true
    and ma.movie_id = m.id
    order by ma.cast_order
  ),
  cast('[]' as json)
) as "actors",
coalesce(
  (
    select
    json_arrayagg(
      json_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      )
    )
    from movies_crew mc
    inner join crew c on mc.crew_id = c.id
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
    and mc.movie_id = m.id
    order by j.friendly_name
  ),
  cast('[]' as json)
) as "crews",
1
from movies m
where true
and m.id = :id; -- :id type: bigint
