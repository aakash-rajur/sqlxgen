select
count(*) over () as "totalRecordsCount",
c.id as "id",
c.name as "name"
from companies c
where true
and (
  false
  or :search is null
  or match(c.name) against (:search in natural language mode)
) -- :search type: text
order by (case when :sort = 'desc' then c.id end) desc, c.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
