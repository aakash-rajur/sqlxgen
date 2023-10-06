-- hyper parameter
drop table if exists hyper_parameters;

create table if not exists hyper_parameters (
  type text not null,
  value text not null,
  friendly_name text not null,
  primary key (type(350), value(350))
);

-- create movies
drop table if exists movies;

create table if not exists movies (
  id serial not null,
  title text not null,
  original_title text not null,
  original_language text not null,
  overview text not null,
  runtime int not null default 0,
  release_date date not null default '1970-01-01',
  tagline text not null,
  status text not null,
  homepage text not null,
  popularity float not null default 0,
  vote_average float not null default 0,
  vote_count int not null default 0,
  budget bigint not null default 0,
  revenue bigint not null default 0,
  keywords text not null,
  primary key(id),
  index (runtime),
  index (release_date),
  index (popularity),
  index(budget),
  index(revenue),
  index(title(750)),
  index(keywords(750))
);

-- create genre associations
drop table if exists movies_genres;

create table if not exists movies_genres (
  movie_id bigint not null,
  genre_id text not null,
  primary key(movie_id, genre_id(750))
);

-- create company
drop table if exists companies;

create table if not exists companies (
  id bigint not null,
  name text not null,
  primary key(id)
);

-- create production associations
drop table if exists movies_companies;

create table if not exists movies_companies (
  movie_id bigint not null,
  company_id bigint not null,
  primary key(movie_id, company_id)
);

-- create country associations
drop table if exists movies_countries;

create table if not exists movies_countries (
  movie_id bigint not null,
  country_id text not null,
  primary key(movie_id, country_id(700))
);

-- create language associations
drop table if exists movies_languages;

create table if not exists movies_languages (
  movie_id bigint not null,
  language_id text not null,
  primary key(movie_id, language_id(700))
);

-- create actors
drop table if exists actors;

create table if not exists actors (
  id serial not null,
  name text not null,
  primary key(id)
);

-- create actor associations
drop table if exists movies_actors;

create table if not exists movies_actors (
  movie_id bigint not null,
  actor_id bigint not null,
  cast text not null,
  cast_order int not null default 0,
  primary key(movie_id, actor_id)
);

-- create crew
drop table if exists crew;

create table if not exists crew (
  id serial not null,
  name text not null,
  primary key(id)
);

-- create crew associations
drop table if exists movies_crew;

create table if not exists movies_crew (
  movie_id bigint not null,
  crew_id bigint not null,
  job_id text not null,
  department_id text not null,
  primary key(movie_id, crew_id)
);
