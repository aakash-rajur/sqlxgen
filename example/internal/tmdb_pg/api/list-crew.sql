select
count(*) over () as "totalRecordsCount",
c.id as "id",
c.name as "name"
from crew c
where true
and (
  false
  or cast(:search as text) is null
  or c.name_search @@ to_tsquery(:search)
) -- :search type: text
order by (case when :sort = 'desc' then c.id end) desc, c.id -- :sort type: text
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
