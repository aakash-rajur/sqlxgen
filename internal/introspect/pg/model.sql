select
ns.nspname  as schema_name,
cls.relname as table_name,
json_agg(
  json_build_object(
    'column_name', attr.attname,
    'type', regexp_replace(tp.typname, '^_(\w+)$', '\1'),
    'type_id', tp.oid,
    'is_array', tp.typcategory = 'A',
    'is_sequence', coalesce(col.column_default like 'nextval(%', false),
    'nullable', not attr.attnotnull,
    'generated', attr.attgenerated = 's',
    'pk_name', coalesce(kcu.constraint_name, ''),
    'pk_ordinal_position', coalesce(kcu.ordinal_position, 0)
  ) order by kcu.ordinal_position, attr.attnum
) as columns
from pg_catalog.pg_attribute attr
inner join pg_catalog.pg_class cls on cls.oid = attr.attrelid
inner join pg_catalog.pg_namespace ns on ns.oid = cls.relnamespace
inner join pg_catalog.pg_type tp on tp.oid = attr.atttypid
left join information_schema.table_constraints tc on (
  true
  and tc.table_schema = ns.nspname
  and tc.table_name = cls.relname
  and tc.constraint_type = 'PRIMARY KEY'
)
left join information_schema.key_column_usage kcu on (
  true
  and kcu.table_schema = ns.nspname
  and kcu.table_name = cls.relname
  and kcu.column_name = attr.attname
  and kcu.constraint_name = tc.constraint_name
  and kcu.constraint_schema = tc.constraint_schema
  and tc.constraint_type = 'PRIMARY KEY'
)
left join information_schema.columns col on (
  true
  and col.table_schema = ns.nspname
  and col.table_name = cls.relname
  and col.column_name = attr.attname
)
where true
and ns.nspname = :schema
and cls.relkind in ('r', 't', 'v', 'm', 'p')
and attr.attnum >= 1
group by ns.nspname, cls.relname
order by cls.relname;