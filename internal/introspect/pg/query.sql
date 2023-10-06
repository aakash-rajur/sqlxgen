--
drop table if exists sample_query_introspection;
--
create temp table if not exists sample_query_introspection as
{{.Query}};
--
select
attr.attname as column_name,
regexp_replace(tp.typname, '^_(\w+)$', '\1') as type,
tp.oid as type_id,
tp.typcategory = 'A' as is_array,
false as is_sequence,
not attr.attnotnull as nullable,
attr.attgenerated = 's' as generated
from pg_attribute attr
inner join pg_type tp on tp.oid = attr.atttypid
where true
and attr.attrelid = cast('sample_query_introspection' as regclass)
and attr.attnum > 0
and not attr.attisdropped
order by attnum;
