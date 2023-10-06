select
count(*) over (partition by hp.type) as "totalRecordsCount",
hp.value as "value",
hp.friendly_name as "friendly_name"
from hyper_parameters hp
where true
and hp.type = :type -- :type type: text
and (:value is null or hp.value = :value) -- :value type: text
and (:search is null or match(hp.friendly_name) against (:search in natural language mode)) -- :search type: text
order by hp.friendly_name
limit :limit -- :limit type: int
offset :offset; -- :offset type: int
