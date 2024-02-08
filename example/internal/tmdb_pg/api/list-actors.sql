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
