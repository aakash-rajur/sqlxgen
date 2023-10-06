package models

import (
	"fmt"
	"strings"
)

type HyperParameter struct {
	FriendlyName *string `db:"friendly_name" json:"friendly_name"`
	Value        *string `db:"value" json:"value"`
	Type         *string `db:"type" json:"type"`
}

func (hyperParameter HyperParameter) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("FriendlyName: %v", *hyperParameter.FriendlyName),
			fmt.Sprintf("Value: %v", *hyperParameter.Value),
			fmt.Sprintf("Type: %v", *hyperParameter.Type),
		},
		", ",
	)

	return fmt.Sprintf("HyperParameter{%s}", content)
}

func (_ HyperParameter) TableName() string {
	return "app.hyper_parameters"
}

func (_ HyperParameter) PrimaryKey() []string {
	return []string{
		"value",
		"type",
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

// language=mysql
var hyperParameterInsertSql = `
INSERT INTO app.hyper_parameters(
  friendly_name,
  value,
  type
)
VALUES (
  :friendly_name,
  :value,
  :type
)
RETURNING
  friendly_name,
  value,
  type;
`

// language=mysql
var hyperParameterUpdateSql = `
UPDATE app.hyper_parameters
SET
  friendly_name = :friendly_name,
  value = :value,
  type = :type
WHERE TRUE
  AND value = :value
  AND type = :type
RETURNING
  friendly_name,
  value,
  type;
`

// language=mysql
var hyperParameterFindSql = `
SELECT
  friendly_name,
  value,
  type
FROM app.hyper_parameters
WHERE TRUE
  AND (:value IS NULL or value = :value)
  AND (:type IS NULL or type = :type)
LIMIT 1;
`

// language=mysql
var hyperParameterFindAllSql = `
SELECT
  friendly_name,
  value,
  type
FROM app.hyper_parameters
WHERE TRUE
  AND (:value IS NULL or value = :value)
  AND (:type IS NULL or type = :type);
`

// language=mysql
var hyperParameterDeleteSql = `
DELETE FROM app.hyper_parameters
WHERE TRUE
  AND value = :value
  AND type = :type;
`
