select
count(*) over () as "totalRecordsCount",
a.id as "id",
a.name as "name"
from actors a
where true
and (
  false
  or :search is null
  or match(a.name) against (:search in natural language mode)
) -- :search type: text
order by (case when :sort = 'desc' then a.id end) desc, a.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
