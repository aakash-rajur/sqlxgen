package models

import (
	"fmt"
	"strings"
)

type HyperParameter struct {
	Type               *string `db:"type" json:"type"`
	Value              *string `db:"value" json:"value"`
	FriendlyName       *string `db:"friendly_name" json:"friendly_name"`
	FriendlyNameSearch *string `db:"friendly_name_search" json:"friendly_name_search"`
}

func (hyperParameter HyperParameter) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("Type: %v", *hyperParameter.Type),
			fmt.Sprintf("Value: %v", *hyperParameter.Value),
			fmt.Sprintf("FriendlyName: %v", *hyperParameter.FriendlyName),
			fmt.Sprintf("FriendlyNameSearch: %v", *hyperParameter.FriendlyNameSearch),
		},
		", ",
	)

	return fmt.Sprintf("HyperParameter{%s}", content)
}

func (_ HyperParameter) TableName() string {
	return "public.hyper_parameters"
}

func (_ HyperParameter) PrimaryKey() []string {
	return []string{
		"type",
		"value",
	}
}

func (_ HyperParameter) InsertQuery() string {
	return hyperParameterInsertSql
}

func (_ HyperParameter) UpdateQuery() string {
	return hyperParameterUpdateSql
}

func (_ HyperParameter) FindQuery() string {
	return hyperParameterFindSql
}

func (_ HyperParameter) FindAllQuery() string {
	return hyperParameterFindAllSql
}

func (_ HyperParameter) DeleteQuery() string {
	return hyperParameterDeleteSql
}

// language=postgresql
var hyperParameterInsertSql = `
INSERT INTO public.hyper_parameters(
  type,
  value,
  friendly_name
)
VALUES (
  :type,
  :value,
  :friendly_name
)
RETURNING
  type,
  value,
  friendly_name,
  friendly_name_search;
`

// language=postgresql
var hyperParameterUpdateSql = `
UPDATE public.hyper_parameters
SET
  type = :type,
  value = :value,
  friendly_name = :friendly_name
WHERE TRUE
  AND type = :type
  AND value = :value
RETURNING
  type,
  value,
  friendly_name,
  friendly_name_search;
`

// language=postgresql
var hyperParameterFindSql = `
SELECT
  type,
  value,
  friendly_name,
  friendly_name_search
FROM public.hyper_parameters
WHERE TRUE
  AND (CAST(:type AS TEXT) IS NULL or type = :type)
  AND (CAST(:value AS TEXT) IS NULL or value = :value)
LIMIT 1;
`

// language=postgresql
var hyperParameterFindAllSql = `
SELECT
  type,
  value,
  friendly_name,
  friendly_name_search
FROM public.hyper_parameters
WHERE TRUE
  AND (CAST(:type AS TEXT) IS NULL or type = :type)
  AND (CAST(:value AS TEXT) IS NULL or value = :value);
`

// language=postgresql
var hyperParameterDeleteSql = `
DELETE FROM public.hyper_parameters
WHERE TRUE
  AND type = :type
  AND value = :value;
`
