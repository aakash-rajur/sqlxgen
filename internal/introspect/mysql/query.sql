--
drop table if exists sample_query_introspection;
--
create table if not exists sample_query_introspection
{{.Query}};
--
select
c.column_name as column_name,
c.data_type as type,
'0' as type_id,
false as is_array,
false as is_sequence,
c.is_nullable = 'YES' as nullable,
c.generation_expression != '' as "generated"
from information_schema.columns c
left join information_schema.key_column_usage kc on (
  true
  and kc.table_schema = c.table_schema
  and kc.table_name = c.table_name
  and kc.column_name = c.column_name
  and kc.constraint_name = 'PRIMARY'
)
where true
and c.table_schema = 'public'
and c.table_name = 'sample_query_introspection';
