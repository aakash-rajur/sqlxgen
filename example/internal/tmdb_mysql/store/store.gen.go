package store

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/jmoiron/sqlx"
)

func Insert[T model](db Database, instances ...T) error {
	for _, instance := range instances {
		rows, err := db.NamedQuery(instance.InsertQuery(), instance)

		if err != nil {
			return err
		}

		rows.Next()

		err = rows.StructScan(&instance)

		_ = rows.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func Update[T model](db Database, instances ...T) error {
	for _, instance := range instances {
		rows, err := db.NamedQuery(instance.UpdateQuery(), instance)

		if err != nil {
			return err
		}

		rows.Next()

		err = rows.StructScan(&instance)

		_ = rows.Close()

		if err != nil {
			return err
		}
	}

	return nil
}

func Find[T model](db Database, instance T) (T, error) {
	var result T

	rows, err := db.NamedQuery(instance.FindQuery(), instance)

	if err != nil {
		return result, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	rows.Next()

	err = rows.StructScan(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func FindMany[T model](db Database, instance T) ([]T, error) {
	rows, err := db.NamedQuery(instance.FindAllQuery(), instance)

	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	result := make([]T, 0)

	for rows.Next() {
		var instance T

		err = rows.StructScan(&instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}

	return result, nil
}

func Delete[T model](db Database, instance T) error {
	_, err := db.NamedExec(instance.DeleteQuery(), instance)

	return err
}

type model interface {
	TableName() string

	PrimaryKey() []string

	InsertQuery() string

	UpdateQuery() string

	FindQuery() string

	FindAllQuery() string

	DeleteQuery() string
}

func Query[R any](db Database, args queryable) ([]R, error) {
	re, err := regexp.Compile(`-{2,}\s*([\w\W\s\S]*?)(\n|\z)`)

	if err != nil {
		return nil, err
	}

	query := re.ReplaceAllString(args.Sql(), "$2")

	rows, err := db.NamedQuery(query, args)

	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	result := make([]R, 0)

	for rows.Next() {
		var instance R

		err = rows.StructScan(&instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}

	return result, nil
}

type queryable interface {
	Sql() string
}

type Database interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
}

type JsonObject map[string]interface{}

func (j *JsonObject) Scan(src any) error {
	jsonBytes, ok := src.([]byte)

	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	err := json.Unmarshal(jsonBytes, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonObject) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type JsonArray []map[string]interface{}

func (j *JsonArray) Scan(src any) error {
	jsonBytes, ok := src.([]byte)

	if !ok {
		return fmt.Errorf("expected []byte, got %T", src)
	}

	err := json.Unmarshal(jsonBytes, &j)

	if err != nil {
		return err
	}

	return nil
}

func (j *JsonArray) Value() (driver.Value, error) {
	return json.Marshal(j)
}
