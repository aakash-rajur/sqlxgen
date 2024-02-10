package store

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/jmoiron/sqlx"
)

// orm

func GetTypeName[T any](instance T) string {
	t := reflect.TypeOf(instance)
	typeName := t.Name()
	if t.Kind() == reflect.Pointer {
		typeName = t.Elem().Name()
	}
	return typeName
}

func InsertOne[T model[P], P any](db Database, instance T) (T, error) {
	inserted, err := Insert[T](db, instance)

	if err != nil {
		return nil, err
	}

	return inserted[0], nil
}

func Insert[T model[P], P any](db Database, instances ...T) ([]T, error) {
	inserts := make([]T, 0)

	for _, instance := range instances {
		rows, err := db.NamedQuery(instance.InsertQuery(), instance)

		if err != nil {
			return nil, err
		}

		hasNext := rows.Next()

		if !hasNext {
			return nil, fmt.Errorf("unable to insert %s", GetTypeName(instance))
		}

		inserted := new(P)

		err = rows.StructScan(inserted)

		if err != nil {
			return nil, err
		}

		err = rows.Close()

		if err != nil {
			return nil, err
		}

		inserts = append(inserts, inserted)
	}

	return inserts, nil
}

// TODO: implement bulk insert
// func InsertBulk[T model[P], P any](db Database, instances ...T) (int64, error) {}

func UpdateOne[T model[P], P any](db Database, instance T) (T, error) {

	count, err := Count[T](db, instance)
	if err != nil {
		return nil, err
	}
	if *count != 1 {
		return nil, fmt.Errorf("update-one %s would have matched %d rows", GetTypeName(instance), *count)
	}

	updates, err := Update[T](db, instance)

	if err != nil {
		return nil, err
	}
	return updates[0], nil
}

func Update[T model[P], P any](db Database, instances ...T) ([]T, error) {
	updates := make([]T, 0)

	for _, instance := range instances {
		rows, err := db.NamedQuery(instance.UpdateQuery(), instance)

		if err != nil {
			return nil, err
		}

		hasNext := rows.Next()

		if !hasNext {
			return nil, fmt.Errorf("unable to update %s", GetTypeName(instance))
		}

		updated := new(P)

		err = rows.StructScan(updated)

		if err != nil {
			return nil, err
		}

		err = rows.Close()

		if err != nil {
			return nil, err
		}

		updates = append(updates, updated)
	}

	return updates, nil
}

func Count[T model[P], P any](db Database, instance T) (*int64, error) {
	var count int64

	err := db.QueryRow(instance.CountQuery()).Scan(&count)

	if err != nil {
		return nil, err
	}

	return &count, nil
}

func findSingle[T model[P], P any](db Database, instance T, sqlQuery string) (T, error) {
	result := new(P)

	rows, err := db.NamedQuery(sqlQuery, instance)

	if err != nil {
		return result, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	hasNext := rows.Next()

	if !hasNext {
		msg := fmt.Sprintf("'%s' not found", GetTypeName(instance))

		return result, fmt.Errorf(msg, ErrNotFound)
	}

	err = rows.StructScan(result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func findMany[T model[P], P any](db Database, instance T, sqlQuery string) ([]T, error) {
	rows, err := db.NamedQuery(sqlQuery, instance)

	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	result := make([]T, 0)

	for rows.Next() {
		instance := new(P)

		err = rows.StructScan(instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}

	return result, nil
}

func Find[T model[P], P any](db Database, instance T) (T, error) {
	// ** Deprecated:
	//    Use FindFirst or FindOne depending on your expectation if > 1 rows match criteria.
	return FindFirst(db, instance)
}

// Find limit 1
func FindFirst[T model[P], P any](db Database, instance T) (T, error) {
	return findSingle(db, instance, instance.FindFirstQuery())
}

// Find and return 1, err if > 1
func FindOne[T model[P], P any](db Database, instance T) (T, error) {
	limitedQuery := strings.TrimRight(instance.FindAllQuery(), ";") + " LIMIT 2;"

	result, err := findMany(db, instance, limitedQuery)
	if err != nil {
		return nil, err
	}
	if len(result) > 1 {
		return nil, ErrFoundMultiple
	}
	return result[0], nil
}

func FindByPk[T model[P], P any](db Database, instance T) (T, error) {
	return findSingle(db, instance, instance.FindByPkQuery())
}

func FindMany[T model[P], P any](db Database, instance T) ([]T, error) {
	return findMany(db, instance, instance.FindAllQuery())
}

// Delete by Pk, err if not found
func DeleteByPk[T model[P], P any](db Database, instance T) error {

	result, err := db.NamedExec(instance.DeleteByPkQuery(), instance)
	if err != nil {
		return err
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAff != 1 {
		return fmt.Errorf("unable to delete %s", instance)
	}
	return nil
}

func DeleteOne[T model[P], P any](db Database, instance T) error {
	count, err := Count[T](db, instance)
	if err != nil {
		return err
	}
	if *count != 1 {
		return fmt.Errorf("delete-one %s would have matched %d rows", GetTypeName(instance), *count)
	}

	_, err = Delete[T](db, instance)

	return err
}

func Delete[T model[P], P any](db Database, instance T) (*int64, error) {

	result, err := db.NamedExec(instance.DeleteQuery(), instance)

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	return &rowsAff, nil
}

type model[P any] interface {
	*P

	TableName() string
	PrimaryKey() []string

	InsertQuery() string
	UpdateQuery() string
	UpdateByPkQuery() string

	CountQuery() string
	FindFirstQuery() string
	FindByPkQuery() string
	FindAllQuery() string

	DeleteByPkQuery() string
	DeleteQuery() string
}

// Custom sql query used like:
//
//	func (args *CatalogQryArgs) Query(db store.Database) ([]*CatalogQryResult, error) {
//		return store.Query[*CatalogQryResult](db, args)
//	}
func Query[R result[pR], Q queryable[pQ], pR, pQ any](db Database, args Q) ([]R, error) {
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
		instance := new(pR)

		err = rows.StructScan(instance)

		if err != nil {
			return nil, err
		}

		result = append(result, instance)
	}

	err = rows.Close()

	if err != nil {
		return nil, err
	}

	return result, nil
}

type queryable[P any] interface {
	*P

	Sql() string
}

type result[P any] interface {
	*P
}

// supplementary types

type Database interface {
	NamedExec(query string, arg interface{}) (sql.Result, error)

	NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)

	QueryRow(query string, args ...interface{}) *sql.Row
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

// errors

var ErrNotFound = errors.New("entity not found")
var ErrFoundMultiple = errors.New("multiple matching entities")
