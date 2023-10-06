-- list hyper parameters, get hyper parameter
select
count(*) over (partition by hp.type) as "totalRecordsCount",
hp.value as "value",
hp.friendly_name as "friendlyName"
from hyper_parameters hp
where true
and hp.type = :type -- :type type: text
and (cast(:value as text) is null or hp.value = :value) -- :value type: text
and (cast(:search as text) is null or hp.friendly_name_search @@ to_tsquery(:search)) -- :search type: text
order by hp.friendly_name
limit :limit -- :limit type: int
offset :offset; -- :offset type: int

-- list movies
select
count(*) over () as "totalRecordsCount",
m.id as "id",
m.title as "title",
m.release_date as "releaseDate",
m.status as "status",
m.popularity as "popularity"
from movies m
where true
and (
  false
  or cast(:search as text) is null
  or m.title_search @@ to_tsquery(:search)
  or m.keywords_search @@ to_tsquery(:search)
) -- :search type: text
and (
  false
  or cast(:genre_id as text) is null
  or m.id in (
    select
    g.movie_id
    from movies_genres g
    where true
    and g.genre_id = :genre_id -- :genre_id type: text
    order by g.movie_id
  )
)
order by (case when :sort = 'desc' then m.id end) desc, m.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int

-- get movie
select
m."id" as "id",
m."title" as "title",
m."original_title" as "originalTitle",
m."original_language" as "originalLanguage",
m."overview" as "overview",
m."runtime" as "runtime",
m."release_date" as "releaseDate",
m."tagline" as "tagline",
m."status" as "status",
m."homepage" as "homepage",
m."popularity" as "popularity",
m."vote_average" as "voteAverage",
m."vote_count" as "voteCount",
m."budget" as "budget",
m."revenue" as "revenue",
m."keywords" as "keywords",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', g.genre_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_genres g
    inner join hyper_parameters hp on (
      true
      and hp.type = 'genre'
      and hp.value = g.genre_id
    )
    where true
    and g.movie_id = m.id
  ),
  '[]'
) as "genres",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', c.country_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_countries c
    inner join hyper_parameters hp on (
      true
      and hp.type = 'country'
      and hp.value = c.country_id
    )
    where true
    and c.movie_id = m.id
  ),
  '[]'
) as "countries",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', l.language_id,
        'name', hp.friendly_name
      ) order by hp.friendly_name
    )
    from movies_languages l
    inner join hyper_parameters hp on (
      true
      and hp.type = 'language'
      and hp.value = l.language_id
    )
    where true
    and l.movie_id = m.id
  ),
  '[]'
) as "languages",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.company_id,
        'name', c.name
      ) order by c.name
    )
    from movies_companies mc
    inner join companies c on mc.company_id = c.id
    where true
    and mc.movie_id = m.id
  ),
  '[]'
) as "companies",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', ma.actor_id,
        'name', a.name,
        'character', ma.character,
        'order', ma.cast_order
      ) order by ma.cast_order
    )
    from movies_actors ma
    inner join actors a on ma.actor_id = a.id
    where true
    and ma.movie_id = m.id
  ),
  '[]'
) as "actors",
coalesce(
  (
    select
    jsonb_agg(
      jsonb_build_object(
        'id', mc.crew_id,
        'name', c.name,
        'job', j.friendly_name,
        'department', d.friendly_name
      ) order by j.friendly_name
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
  ),
  '[]'
) as "crews",
1
from movies m
where true
and m.id = :id; -- :id type: bigint

-- list companies
select
count(*) over () as "totalRecordsCount",
c.id as "id",
c.name as "name"
from companies c
where true
and (
  false
  or cast(:search as text) is null
  or c.name_search @@ to_tsquery(:search)
) -- :search type: text
order by (case when :sort = 'desc' then c.id end) desc, c.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int

-- get company
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

-- list actors
select
count(*) over () as "totalRecordsCount",
a.id as "id",
a.name as "name"
from actors a
where true
and (
  false
  or cast(:search as text) is null
  or a.name_search @@ to_tsquery(:search)
) -- :search type: text
order by (case when :sort = 'desc' then a.id end) desc, a.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int

-- get actor
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

-- list crew
select
count(*) over () as "totalRecordsCount",
c.id as "id",
c.name as "name"
from crew c
where true
and (
  false
  or cast(:search as text) is null
  or c.name_search @@ to_tsquery(:search)
) -- :search type: text
order by (case when :sort = 'desc' then c.id end) desc, c.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int

-- get crew
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
