CREATE EXTENSION IF NOT EXISTS btree_gin;

COMMENT ON EXTENSION btree_gin IS 'support for indexing common datatypes in GIN';

DROP FUNCTION IF EXISTS public.concat_ws_i(text, VARIADIC text[]);
DROP FUNCTION IF EXISTS public.array_to_string_i(text[], text);
--
-- TOC entry 229 (class 1255 OID 112390)
-- Name: array_to_string_i(text[], text); Type: FUNCTION; Schema: public; Owner: -
--

CREATE OR REPLACE FUNCTION public.array_to_string_i(text[], text) RETURNS text
    LANGUAGE internal IMMUTABLE STRICT PARALLEL SAFE
    AS $$array_to_text$$;


--
-- TOC entry 230 (class 1255 OID 112368)
-- Name: concat_ws_i(text, text[]); Type: FUNCTION; Schema: public; Owner: -
--

CREATE OR REPLACE FUNCTION public.concat_ws_i(text, VARIADIC text[]) RETURNS text
    LANGUAGE internal IMMUTABLE PARALLEL SAFE
    AS $$text_concat_ws$$;
