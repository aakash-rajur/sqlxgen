-- dump
--
drop table if exists t_movies;
--
create table if not exists t_movies(
  id bigint not null,
  title text not null default '',
  original_title text not null default '',
  original_language text not null default '',
  overview text default '',
  runtime float default 0,
  release_date date not null default '1970-01-01',
  status text not null default '',
  tagline text default '',
  homepage text default '',
  popularity float not null default 0,
  vote_average float not null default 0,
  vote_count bigint not null default 0,
  budget double precision not null default 0,
  revenue double precision not null default 0,

  genre jsonb not null default '[]',
  keywords jsonb not null default '[]',
  production_companies jsonb not null default '[]',
  production_countries jsonb not null default '[]',
  spoken_languages jsonb not null default '[]',
  primary key(id)
);
--
drop table if exists t_movies_credits;
--
create table if not exists t_movies_credits(
  movie_id bigint not null,
  title text not null default '',
  casting jsonb not null default '[]',
  crew jsonb not null default '[]'
);

-- create schema
drop extension if exists btree_gin cascade;

create extension if not exists btree_gin;

-- create functions
create or replace function concat_ws_i(text, variadic text[])
returns text
language internal immutable parallel safe as 'text_concat_ws';

create or replace function array_to_string_i(text[], text)
 returns text
language internal immutable parallel safe strict as 'array_to_text';

-- hyper parameter
drop table if exists hyper_parameters;

create table if not exists hyper_parameters (
  type text not null,
  value text not null default '',
  friendly_name text not null default '',
  friendly_name_search tsvector generated always as (to_tsvector('english', friendly_name)) stored,
  primary key(type, value)
);

insert into hyper_parameters(
  type,
  value,
  friendly_name
)
(
with
g as (
select
distinct on(g.id)
g.id as genre_id,
g.name as genre_name
from t_movies m,
  jsonb_to_recordset(m.genre) as g(id bigint, name text)
),
pc as (
select
distinct on(p.iso_3166_1)
p.iso_3166_1 as production_country_id,
p.name as production_country_name
from t_movies m,
  jsonb_to_recordset(m.production_countries) as p(iso_3166_1 text, name text)
),
l as (
select
distinct on(l.iso_639_1)
l.iso_639_1 as language_id,
(case when l.name = '' then l.iso_639_1 else l.name end) as language_name
from t_movies m,
  jsonb_to_recordset(m.spoken_languages) as l(iso_639_1 text, name text)
),
j as (
select
distinct on(c.job)
lower(trim(regexp_replace(c.job, ' ', '_', 'g'))) as job_id,
c.job as job_name
from t_movies_credits m,
  jsonb_to_recordset(m.crew) as c(id bigint, name text, job text, department text)
),
d as (
select
distinct on(c.department)
lower(trim(regexp_replace(c.department, ' ', '_', 'g'))) as department_id,
c.department as department_name
from t_movies_credits m,
  jsonb_to_recordset(m.crew) as c(id bigint, name text, job text, department text)
)
(
select
'genre' as type,
lower(regexp_replace(g.genre_name, ' ', '_', 'g')) as value,
g.genre_name as friendly_name
from g
group by g.genre_id, g.genre_name
) union all (
select
'country' as type,
lower(pc.production_country_id) as value,
pc.production_country_name as friendly_name
from pc
) union all (
select
'language' as type,
lower(l.language_id) as value,
l.language_name as friendly_name
from l
) union all (
select
'job' as type,
j.job_id as value,
j.job_name as friendly_name
from j
) union all (
select
'department' as type,
d.department_id as value,
d.department_name as friendly_name
from d
)
)
on conflict do nothing;

create index hyper_parameters_search_idx on hyper_parameters using gin(type, friendly_name_search);

-- create movies
drop table if exists movies;

create table if not exists movies (
  id serial not null,
  title text not null default '',
  original_title text not null default '',
  original_language text not null default '',
  overview text not null default '',
  runtime int not null default 0,
  release_date date not null default '1970-01-01',
  tagline text not null default '',
  status text not null default '',
  homepage text not null default '',
  popularity float not null default 0,
  vote_average float not null default 0,
  vote_count int not null default 0,
  budget bigint not null default 0,
  revenue bigint not null default 0,
  keywords text[] not null default '{}',
  title_search tsvector generated always as (to_tsvector('english', concat_ws_i(' ', title, original_title))) stored,
  keywords_search tsvector generated always as (to_tsvector('english', array_to_string_i(keywords, ' '))) stored,
  primary key(id)
);

insert into movies(
  id,
  title,
  original_title,
  original_language,
  overview,
  runtime,
  release_date,
  tagline,
  status,
  homepage,
  popularity,
  vote_average,
  vote_count,
  budget,
  revenue
)
select
id,
title,
original_title,
original_language,
overview,
runtime,
release_date,
tagline,
status,
homepage,
popularity,
vote_average,
vote_count,
budget,
revenue
from t_movies;

with k as (
  select
  m.id as id,
  array_agg(keywords.name order by keywords.name) as keyword_name
  from t_movies m,
    jsonb_to_recordset(m.keywords) as keywords(id bigint, name text)
  group by m.id
)
update movies
set keywords = k.keyword_name
from k
where movies.id = k.id;

select setval('movies_id_seq', (select max(id) from movies));

create index movies_runtime_idx on movies(runtime);

create index movies_language_idx on movies(original_language);

create index movies_popularity_idx on movies(popularity);

create index movies_budget_idx on movies(budget);

create index movies_revenue_idx on movies(revenue);

create index movies_title_search_idx on movies using gin(title_search);

create index movies_keywords_search_idx on movies using gin(keywords_search);


-- create genre associations
drop table if exists movies_genres;

create table if not exists movies_genres (
  movie_id int not null,
  genre_id text not null,
  primary key(movie_id, genre_id)
);

insert into movies_genres(
  movie_id,
  genre_id
)
with g as (
select
m.id as id,
g.id as genre_id,
g.name as genre_name
from t_movies m,
  jsonb_to_recordset(m.genre) as g(id bigint, name text)
)
select
g.id as id,
lower(regexp_replace(g.genre_name, ' ', '_', 'g')) as genre
from g;

create index movies_genres_genre_idx on movies_genres(genre_id);

-- create company
drop table if exists companies;

create table if not exists companies (
  id bigint not null,
  name text not null default '',
  name_search tsvector generated always as (to_tsvector('english', name)) stored,
  primary key(id)
);

insert into companies(
  id,
  name
)
with t as (
select
m.id as id,
p.id as production_company_id,
p.name as production_company_name,
count(p.id) over (partition by p.id) as productions
from t_movies m,
  jsonb_to_recordset(m.production_companies) as p(id bigint, name text)
)
select
distinct on(t.production_company_id)
t.production_company_id as id,
t.production_company_name as name
from t
order by 1;

create index companies_name_search_idx on companies using gin(name_search);

-- create production associations
drop table if exists movies_companies;

create table if not exists movies_companies (
  movie_id int not null,
  company_id bigint not null,
  primary key(movie_id, company_id)
);

insert into movies_companies(
  movie_id,
  company_id
)
with t as (
select
m.id as id,
p.id as production_company_id,
p.name as production_company_name,
count(p.id) over (partition by p.id) as productions
from t_movies m,
  jsonb_to_recordset(m.production_companies) as p(id bigint, name text)
)
select
t.id as movie_id,
t.production_company_id as company_id
from t
order by 1;

create index movies_companies_company_idx on movies_companies(company_id);

-- create country associations
drop table if exists movies_countries;

create table if not exists movies_countries (
  movie_id int not null,
  country_id text not null,
  primary key(movie_id, country_id)
);

insert into movies_countries(
  movie_id,
  country_id
)
with t as (
select
m.id as id,
p.iso_3166_1 as production_country_id,
p.name as production_country_name
from t_movies m,
  jsonb_to_recordset(m.production_countries) as p(iso_3166_1 text, name text)
)
select
t.id as movie_id,
lower(t.production_country_id) as value
from t;

create index movies_countries_country_idx on movies_countries(country_id);

-- create language associations
drop table if exists movies_languages;

create table if not exists movies_languages (
  movie_id int not null,
  language_id text not null,
  primary key(movie_id, language_id)
);

insert into movies_languages(
  movie_id,
  language_id
)
with t as (
select
m.id as id,
l.iso_639_1 as language_id,
l.name as language_name
from t_movies m,
  jsonb_to_recordset(m.spoken_languages) as l(iso_639_1 text, name text)
)
select
t.id as movie_id,
lower(t.language_id) as value
from t;

create index movies_languages_language_idx on movies_languages(language_id);

-- crate actors
drop table if exists actors;

create table if not exists actors (
  id serial not null,
  name text not null default '',
  name_search tsvector generated always as (to_tsvector('english', name)) stored,
  primary key(id)
);

insert into actors(
  id,
  name
)
select
distinct on(c.id)
c.id as cast_id,
c.name as actor_name
from t_movies_credits m,
  jsonb_to_recordset(m.casting) as c(id bigint, name text, character text, "order" int);

select setval('actors_id_seq', (select max(id) from actors));

create index actors_name_search_idx on actors using gin(name_search);

-- create actor associations
drop table if exists movies_actors;

create table if not exists movies_actors (
  movie_id int not null,
  actor_id bigint not null,
  character text not null default '',
  cast_order int not null default 0,
  character_search tsvector generated always as (to_tsvector('english', character)) stored,
  primary key(movie_id, actor_id)
);

insert into movies_actors(
  movie_id,
  actor_id,
  character,
  cast_order
)
select
m.movie_id as movie_id,
c.id as actor_id,
c.character as character,
c."order" as cast_order
from t_movies_credits m,
  jsonb_to_recordset(m.casting) as c(id bigint, name text, character text, "order" int)
on conflict do nothing;

create index movies_actors_actor_idx on movies_actors(actor_id);

create index movies_actors_character_search_idx on movies_actors using gin(character_search);

-- create crew
drop table if exists crew;

create table if not exists crew (
  id serial not null,
  name text not null default '',
  name_search tsvector generated always as (to_tsvector('english', name)) stored,
  primary key(id)
);

insert into crew(
  id,
  name
)
select
distinct on(c.id)
c.id as crew_id,
c.name as crew_name
from t_movies_credits m,
  jsonb_to_recordset(m.crew) as c(id bigint, name text, job text, department text);

select setval('crew_id_seq', (select max(id) from crew));

create index crew_name_search_idx on crew using gin(name_search);

-- create crew associations
drop table if exists movies_crew;

create table if not exists movies_crew (
  movie_id int not null,
  crew_id bigint not null,
  job_id text not null default '',
  department_id text not null default '',
  primary key(movie_id, crew_id)
);

insert into movies_crew(
  movie_id,
  crew_id,
  job_id,
  department_id
)
select
m.movie_id as movie_id,
c.id as crew_id,
lower(trim(regexp_replace(c.job, ' ', '_', 'g'))) as job,
lower(trim(regexp_replace(c.department, ' ', '_', 'g'))) as department
from t_movies_credits m,
  jsonb_to_recordset(m.crew) as c(id bigint, name text, job text, department text)
on conflict do nothing;

create index movies_crew_crew_idx on movies_crew(crew_id);

create index movies_crew_job_idx on movies_crew(job_id, movie_id);

create index movies_crew_department_idx on movies_crew(department_id, movie_id);
