--
-- PostgreSQL database dump
--

-- Dumped from database version 15.4
-- Dumped by pg_dump version 16.0

-- Started on 2023-10-05 16:01:34 IST

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

DROP INDEX IF EXISTS public.movies_title_search_idx;
DROP INDEX IF EXISTS public.movies_runtime_idx;
DROP INDEX IF EXISTS public.movies_revenue_idx;
DROP INDEX IF EXISTS public.movies_popularity_idx;
DROP INDEX IF EXISTS public.movies_languages_language_idx;
DROP INDEX IF EXISTS public.movies_language_idx;
DROP INDEX IF EXISTS public.movies_keywords_search_idx;
DROP INDEX IF EXISTS public.movies_genres_genre_idx;
DROP INDEX IF EXISTS public.movies_crew_job_idx;
DROP INDEX IF EXISTS public.movies_crew_department_idx;
DROP INDEX IF EXISTS public.movies_crew_crew_idx;
DROP INDEX IF EXISTS public.movies_countries_country_idx;
DROP INDEX IF EXISTS public.movies_companies_company_idx;
DROP INDEX IF EXISTS public.movies_budget_idx;
DROP INDEX IF EXISTS public.movies_actors_character_search_idx;
DROP INDEX IF EXISTS public.movies_actors_actor_idx;
DROP INDEX IF EXISTS public.hyper_parameters_search_idx;
DROP INDEX IF EXISTS public.crew_name_search_idx;
DROP INDEX IF EXISTS public.companies_name_search_idx;
DROP INDEX IF EXISTS public.actors_name_search_idx;
ALTER TABLE IF EXISTS ONLY public.t_movies DROP CONSTRAINT IF EXISTS t_movies_pkey;
ALTER TABLE IF EXISTS ONLY public.movies DROP CONSTRAINT IF EXISTS movies_pkey;
ALTER TABLE IF EXISTS ONLY public.movies_languages DROP CONSTRAINT IF EXISTS movies_languages_pkey;
ALTER TABLE IF EXISTS ONLY public.movies_genres DROP CONSTRAINT IF EXISTS movies_genres_pkey;
ALTER TABLE IF EXISTS ONLY public.movies_crew DROP CONSTRAINT IF EXISTS movies_crew_pkey;
ALTER TABLE IF EXISTS ONLY public.movies_countries DROP CONSTRAINT IF EXISTS movies_countries_pkey;
ALTER TABLE IF EXISTS ONLY public.movies_companies DROP CONSTRAINT IF EXISTS movies_companies_pkey;
ALTER TABLE IF EXISTS ONLY public.movies_actors DROP CONSTRAINT IF EXISTS movies_actors_pkey;
ALTER TABLE IF EXISTS ONLY public.hyper_parameters DROP CONSTRAINT IF EXISTS hyper_parameters_pkey;
ALTER TABLE IF EXISTS ONLY public.crew DROP CONSTRAINT IF EXISTS crew_pkey;
ALTER TABLE IF EXISTS ONLY public.companies DROP CONSTRAINT IF EXISTS companies_pkey;
ALTER TABLE IF EXISTS ONLY public.actors DROP CONSTRAINT IF EXISTS actors_pkey;
ALTER TABLE IF EXISTS public.movies ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.crew ALTER COLUMN id DROP DEFAULT;
ALTER TABLE IF EXISTS public.actors ALTER COLUMN id DROP DEFAULT;
DROP TABLE IF EXISTS public.t_movies_credits;
DROP TABLE IF EXISTS public.t_movies;
DROP TABLE IF EXISTS public.movies_languages;
DROP SEQUENCE IF EXISTS public.movies_id_seq;
DROP TABLE IF EXISTS public.movies_genres;
DROP TABLE IF EXISTS public.movies_crew;
DROP TABLE IF EXISTS public.movies_countries;
DROP TABLE IF EXISTS public.movies_companies;
DROP TABLE IF EXISTS public.movies_actors;
DROP TABLE IF EXISTS public.movies;
DROP TABLE IF EXISTS public.hyper_parameters;
DROP SEQUENCE IF EXISTS public.crew_id_seq;
DROP TABLE IF EXISTS public.crew;
DROP TABLE IF EXISTS public.companies;
DROP SEQUENCE IF EXISTS public.actors_id_seq;
DROP TABLE IF EXISTS public.actors;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 228 (class 1259 OID 125685)
-- Name: actors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.actors (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    name_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, name)) STORED
);


--
-- TOC entry 227 (class 1259 OID 125684)
-- Name: actors_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.actors_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3626 (class 0 OID 0)
-- Dependencies: 227
-- Name: actors_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.actors_id_seq OWNED BY public.actors.id;


--
-- TOC entry 223 (class 1259 OID 125652)
-- Name: companies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.companies (
    id bigint NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    name_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, name)) STORED
);


--
-- TOC entry 231 (class 1259 OID 125709)
-- Name: crew; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.crew (
    id integer NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    name_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, name)) STORED
);


--
-- TOC entry 230 (class 1259 OID 125708)
-- Name: crew_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.crew_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3627 (class 0 OID 0)
-- Dependencies: 230
-- Name: crew_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.crew_id_seq OWNED BY public.crew.id;


--
-- TOC entry 219 (class 1259 OID 125592)
-- Name: hyper_parameters; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.hyper_parameters (
    type text NOT NULL,
    value text DEFAULT ''::text NOT NULL,
    friendly_name text DEFAULT ''::text NOT NULL,
    friendly_name_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, friendly_name)) STORED
);


--
-- TOC entry 221 (class 1259 OID 125604)
-- Name: movies; Type: TABLE; Schema: public; Owner: -
--


CREATE TABLE public.movies (
    id integer NOT NULL,
    title text DEFAULT ''::text NOT NULL,
    original_title text DEFAULT ''::text NOT NULL,
    original_language text DEFAULT ''::text NOT NULL,
    overview text DEFAULT ''::text NOT NULL,
    synopsis varchar DEFAULT ''::text NOT NULL,
    summary varchar(50) DEFAULT ''::text NOT NULL,
	search_vector tsvector NULL,
	client_id varchar NULL,
	-- metadata public.hstore NULL,
	location_accuracy int4 NULL,
	data_synced_at timestamp NOT NULL DEFAULT now(),
	completed_coordinates point NULL,
	is_completed bool NULL,
	distance_to_place numeric NULL,
    runtime integer DEFAULT 0 NOT NULL,
    release_date date DEFAULT '1970-01-01'::date NOT NULL,
    tagline text DEFAULT ''::text NOT NULL,
    status text DEFAULT ''::text NOT NULL,
    homepage text DEFAULT ''::text NOT NULL,
    popularity double precision DEFAULT 0 NOT NULL,
    vote_average double precision DEFAULT 0 NOT NULL,
    vote_count integer DEFAULT 0 NOT NULL,
    budget bigint DEFAULT 0 NOT NULL,
    revenue bigint DEFAULT 0 NOT NULL,
    keywords text[] DEFAULT '{}'::text[] NOT NULL,
    title_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, public.concat_ws_i(' '::text, VARIADIC ARRAY[title, original_title]))) STORED,
    keywords_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, public.array_to_string_i(keywords, ' '::text))) STORED
);


--
-- TOC entry 229 (class 1259 OID 125696)
-- Name: movies_actors; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_actors (
    movie_id bigint NOT NULL,
    actor_id bigint NOT NULL,
    "character" text DEFAULT ''::text NOT NULL,
    cast_order integer DEFAULT 0 NOT NULL,
    character_search tsvector GENERATED ALWAYS AS (to_tsvector('english'::regconfig, "character")) STORED
);


--
-- TOC entry 224 (class 1259 OID 125662)
-- Name: movies_companies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_companies (
    movie_id bigint NOT NULL,
    company_id bigint NOT NULL
);


--
-- TOC entry 225 (class 1259 OID 125668)
-- Name: movies_countries; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_countries (
    movie_id bigint NOT NULL,
    country_id text NOT NULL
);


--
-- TOC entry 232 (class 1259 OID 125720)
-- Name: movies_crew; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_crew (
    movie_id bigint NOT NULL,
    crew_id bigint NOT NULL,
    job_id text DEFAULT ''::text NOT NULL,
    department_id text DEFAULT ''::text NOT NULL
);


--
-- TOC entry 222 (class 1259 OID 125644)
-- Name: movies_genres; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_genres (
    movie_id bigint NOT NULL,
    genre_id text NOT NULL
);


--
-- TOC entry 220 (class 1259 OID 125603)
-- Name: movies_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.movies_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3628 (class 0 OID 0)
-- Dependencies: 220
-- Name: movies_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.movies_id_seq OWNED BY public.movies.id;


--
-- TOC entry 226 (class 1259 OID 125676)
-- Name: movies_languages; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.movies_languages (
    movie_id bigint NOT NULL,
    language_id text NOT NULL
);


--
-- TOC entry 217 (class 1259 OID 108427)
-- Name: t_movies; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.t_movies (
    id bigint NOT NULL,
    budget double precision DEFAULT 0 NOT NULL,
    homepage text DEFAULT ''::text,
    genre jsonb DEFAULT '[]'::jsonb NOT NULL,
    keywords jsonb DEFAULT '[]'::jsonb NOT NULL,
    original_language text DEFAULT ''::text NOT NULL,
    original_title text DEFAULT ''::text NOT NULL,
    overview text DEFAULT ''::text,
    popularity double precision DEFAULT 0 NOT NULL,
    production_companies jsonb DEFAULT '[]'::jsonb NOT NULL,
    production_countries jsonb DEFAULT '[]'::jsonb NOT NULL,
    release_date date DEFAULT '1970-01-01'::date NOT NULL,
    revenue double precision DEFAULT 0 NOT NULL,
    runtime double precision DEFAULT 0,
    spoken_languages jsonb DEFAULT '[]'::jsonb NOT NULL,
    status text DEFAULT ''::text NOT NULL,
    tagline text DEFAULT ''::text,
    title text DEFAULT ''::text NOT NULL,
    vote_average double precision DEFAULT 0 NOT NULL,
    vote_count bigint DEFAULT 0 NOT NULL
);


--
-- TOC entry 218 (class 1259 OID 108453)
-- Name: t_movies_credits; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.t_movies_credits (
    movie_id bigint NOT NULL,
    title text DEFAULT ''::text NOT NULL,
    casting jsonb DEFAULT '[]'::jsonb NOT NULL,
    crew jsonb DEFAULT '[]'::jsonb NOT NULL
);


--
-- TOC entry 3423 (class 2604 OID 125688)
-- Name: actors id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.actors ALTER COLUMN id SET DEFAULT nextval('public.actors_id_seq'::regclass);


--
-- TOC entry 3429 (class 2604 OID 125712)
-- Name: crew id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.crew ALTER COLUMN id SET DEFAULT nextval('public.crew_id_seq'::regclass);


--
-- TOC entry 3403 (class 2604 OID 125607)
-- Name: movies id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies ALTER COLUMN id SET DEFAULT nextval('public.movies_id_seq'::regclass);


--
-- TOC entry 3465 (class 2606 OID 125694)
-- Name: actors actors_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.actors
    ADD CONSTRAINT actors_pkey PRIMARY KEY (id);


--
-- TOC entry 3453 (class 2606 OID 125660)
-- Name: companies companies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.companies
    ADD CONSTRAINT companies_pkey PRIMARY KEY (id);


--
-- TOC entry 3472 (class 2606 OID 125718)
-- Name: crew crew_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.crew
    ADD CONSTRAINT crew_pkey PRIMARY KEY (id);


--
-- TOC entry 3437 (class 2606 OID 125601)
-- Name: hyper_parameters hyper_parameters_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.hyper_parameters
    ADD CONSTRAINT hyper_parameters_pkey PRIMARY KEY (type, value);


--
-- TOC entry 3469 (class 2606 OID 125705)
-- Name: movies_actors movies_actors_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_actors
    ADD CONSTRAINT movies_actors_pkey PRIMARY KEY (movie_id, actor_id);


--
-- TOC entry 3456 (class 2606 OID 125666)
-- Name: movies_companies movies_companies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_companies
    ADD CONSTRAINT movies_companies_pkey PRIMARY KEY (movie_id, company_id);


--
-- TOC entry 3459 (class 2606 OID 125674)
-- Name: movies_countries movies_countries_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_countries
    ADD CONSTRAINT movies_countries_pkey PRIMARY KEY (movie_id, country_id);


--
-- TOC entry 3477 (class 2606 OID 125728)
-- Name: movies_crew movies_crew_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_crew
    ADD CONSTRAINT movies_crew_pkey PRIMARY KEY (movie_id, crew_id);


--
-- TOC entry 3450 (class 2606 OID 125650)
-- Name: movies_genres movies_genres_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_genres
    ADD CONSTRAINT movies_genres_pkey PRIMARY KEY (movie_id, genre_id);


--
-- TOC entry 3462 (class 2606 OID 125682)
-- Name: movies_languages movies_languages_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies_languages
    ADD CONSTRAINT movies_languages_pkey PRIMARY KEY (movie_id, language_id);


--
-- TOC entry 3443 (class 2606 OID 125628)
-- Name: movies movies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.movies
    ADD CONSTRAINT movies_pkey PRIMARY KEY (id);


--
-- TOC entry 3435 (class 2606 OID 108452)
-- Name: t_movies t_movies_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.t_movies
    ADD CONSTRAINT t_movies_pkey PRIMARY KEY (id);


--
-- TOC entry 3463 (class 1259 OID 125695)
-- Name: actors_name_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX actors_name_search_idx ON public.actors USING gin (name_search);


--
-- TOC entry 3451 (class 1259 OID 125661)
-- Name: companies_name_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX companies_name_search_idx ON public.companies USING gin (name_search);


--
-- TOC entry 3470 (class 1259 OID 125719)
-- Name: crew_name_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX crew_name_search_idx ON public.crew USING gin (name_search);


--
-- TOC entry 3438 (class 1259 OID 125602)
-- Name: hyper_parameters_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX hyper_parameters_search_idx ON public.hyper_parameters USING gin (type, friendly_name_search);


--
-- TOC entry 3466 (class 1259 OID 125706)
-- Name: movies_actors_actor_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_actors_actor_idx ON public.movies_actors USING btree (actor_id);


--
-- TOC entry 3467 (class 1259 OID 125707)
-- Name: movies_actors_character_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_actors_character_search_idx ON public.movies_actors USING gin (character_search);


--
-- TOC entry 3439 (class 1259 OID 125640)
-- Name: movies_budget_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_budget_idx ON public.movies USING btree (budget);


--
-- TOC entry 3454 (class 1259 OID 125667)
-- Name: movies_companies_company_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_companies_company_idx ON public.movies_companies USING btree (company_id);


--
-- TOC entry 3457 (class 1259 OID 125675)
-- Name: movies_countries_country_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_countries_country_idx ON public.movies_countries USING btree (country_id);


--
-- TOC entry 3473 (class 1259 OID 125729)
-- Name: movies_crew_crew_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_crew_crew_idx ON public.movies_crew USING btree (crew_id);


--
-- TOC entry 3474 (class 1259 OID 125731)
-- Name: movies_crew_department_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_crew_department_idx ON public.movies_crew USING btree (department_id, movie_id);


--
-- TOC entry 3475 (class 1259 OID 125730)
-- Name: movies_crew_job_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_crew_job_idx ON public.movies_crew USING btree (job_id, movie_id);


--
-- TOC entry 3448 (class 1259 OID 125651)
-- Name: movies_genres_genre_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_genres_genre_idx ON public.movies_genres USING btree (genre_id);


--
-- TOC entry 3440 (class 1259 OID 125643)
-- Name: movies_keywords_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_keywords_search_idx ON public.movies USING gin (keywords_search);


--
-- TOC entry 3441 (class 1259 OID 125638)
-- Name: movies_language_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_language_idx ON public.movies USING btree (original_language);


--
-- TOC entry 3460 (class 1259 OID 125683)
-- Name: movies_languages_language_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_languages_language_idx ON public.movies_languages USING btree (language_id);


--
-- TOC entry 3444 (class 1259 OID 125639)
-- Name: movies_popularity_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_popularity_idx ON public.movies USING btree (popularity);


--
-- TOC entry 3445 (class 1259 OID 125641)
-- Name: movies_revenue_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_revenue_idx ON public.movies USING btree (revenue);


--
-- TOC entry 3446 (class 1259 OID 125637)
-- Name: movies_runtime_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_runtime_idx ON public.movies USING btree (runtime);


--
-- TOC entry 3447 (class 1259 OID 125642)
-- Name: movies_title_search_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX movies_title_search_idx ON public.movies USING gin (title_search);


-- Completed on 2023-10-05 16:01:34 IST

--
-- PostgreSQL database dump complete
--

