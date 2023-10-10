package models

import (
	"fmt"
	"strings"
)

type HyperParameter struct {
	FriendlyName *string `db:"friendly_name" json:"friendly_name"`
	Type         *string `db:"type" json:"type"`
	Value        *string `db:"value" json:"value"`
}

func (hyperParameter HyperParameter) String() string {
	content := strings.Join(
		[]string{
			fmt.Sprintf("FriendlyName: %v", *hyperParameter.FriendlyName),
			fmt.Sprintf("Type: %v", *hyperParameter.Type),
			fmt.Sprintf("Value: %v", *hyperParameter.Value),
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

// language=mysql
var hyperParameterInsertSql = `
INSERT INTO app.hyper_parameters(
  friendly_name,
  type,
  value
)
VALUES (
  :friendly_name,
  :type,
  :value
)
RETURNING
  friendly_name,
  type,
  value;
`

// language=mysql
var hyperParameterUpdateSql = `
UPDATE app.hyper_parameters
SET
  friendly_name = :friendly_name,
  type = :type,
  value = :value
WHERE TRUE
  AND type = :type
  AND value = :value
RETURNING
  friendly_name,
  type,
  value;
`

// language=mysql
var hyperParameterFindSql = `
SELECT
  friendly_name,
  type,
  value
FROM app.hyper_parameters
WHERE TRUE
  AND (:type IS NULL or type = :type)
  AND (:value IS NULL or value = :value)
LIMIT 1;
`

// language=mysql
var hyperParameterFindAllSql = `
SELECT
  friendly_name,
  type,
  value
FROM app.hyper_parameters
WHERE TRUE
  AND (:type IS NULL or type = :type)
  AND (:value IS NULL or value = :value);
`

// language=mysql
var hyperParameterDeleteSql = `
DELETE FROM app.hyper_parameters
WHERE TRUE
  AND type = :type
  AND value = :value;
`
