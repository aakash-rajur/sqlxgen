select
count(*) over (partition by hp.type) as "totalRecordsCount",
hp.value as "value",
hp.friendly_name as "friendlyName"
from hyper_parameters hp
where true
and hp.type = :type -- :type type: text
and (cast(:value as text) is null or hp.value = :value) -- :value type: text
and (cast(:search as text) is null or hp.friendly_name_search @@ to_tsquery(:search)) -- :search type: text
order by hp.friendly_name
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
