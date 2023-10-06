select
c.table_schema as schema_name,
c.table_name as table_name,
json_arrayagg(
  json_object(
    'column_name', c.column_name,
    'type', c.data_type,
    'type_id', '0',
    'nullable', c.is_nullable = 'YES',
    'is_array', false,
    'is_sequence', c.extra like '%auto_increment%',
    'generated', c.generation_expression != '',
    'pk_name', coalesce(kc.constraint_name, ''),
    'pk_ordinal_position', coalesce(kc.ordinal_position, 0)
  )
) as columns
from information_schema.columns c
left join information_schema.KEY_COLUMN_USAGE kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = :schema
group by c.table_schema, c.table_name;