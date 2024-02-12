package store

import (
	"context"
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
		insertSql := instance.InsertQuery()
		rows, err := db.NamedQuery(insertSql, instance)

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

func UpdateByPk[T model[P], P any](db Database, instance T) (T, error) {

	updateSql := instance.UpdateByPkQuery()
	rows, err := db.NamedQuery(updateSql, instance)

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

	return updated, nil
}

func UpdateOne[T model[P], P any](db Database, instance T) (T, error) {

	count, err := Count[T](db, instance)
	if err != nil {
		return nil, err
	}
	if *count != 1 {
		return nil, fmt.Errorf("update-one %s would have matched %d rows", GetTypeName(instance), *count)
	}

	updates, err := UpdateMany[T](db, instance)

	if err != nil {
		return nil, err
	}
	return updates[0], nil
}

func UpdateMany[T model[P], P any](db Database, instances ...T) ([]T, error) {
	updates := make([]T, 0)

	for _, instance := range instances {
		updateSql := instance.UpdateAllQuery()
		rows, err := db.NamedQuery(updateSql, instance)

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
	type CountResult struct {
		Count int64 `db:"count"`
	}
	countResult := new(CountResult)

	countSql := instance.CountQuery()
	rows, err := db.NamedQuery(countSql, instance)
	if err != nil {
		return nil, err
	}

	defer func(rows *sqlx.Rows) {
		_ = rows.Close()
	}(rows)

	hasNext := rows.Next()
	if !hasNext {
		msg := fmt.Sprintf("count '%s' failed", GetTypeName(instance))

		return nil, fmt.Errorf(msg, ErrNotFound)
	}

	err = rows.StructScan(countResult)
	if err != nil {
		return nil, err
	}

	return &countResult.Count, nil
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

type QueryOptions struct {
	SelectList *[]string
	Paginator  *Paginator
	OrderBy    *string
}

type Paginator struct {
	Page     int
	PageSize int
}

func FindMany[T model[P], P any](db Database, instance T) ([]T, error) {
	return FindPage[T](db, instance, nil)
}

func FindPage[T model[P], P any](db Database, instance T, queryOpts *QueryOptions) ([]T, error) {

	findAllSql := instance.FindAllQuery()
	if queryOpts != nil {
		if queryOpts.SelectList != nil && len(*queryOpts.SelectList) > 0 {
			selectList := *queryOpts.SelectList
			fromSql := fmt.Sprintf("FROM %s", instance.TableName())
			whereSql := strings.Split(findAllSql, fromSql)[1]
			findAllSql = fmt.Sprintf("SELECT %s %s %s", strings.Join(selectList, ", "), fromSql, whereSql)
		}
		if queryOpts.Paginator != nil {
			pager := *queryOpts.Paginator
			findAllSql = strings.TrimRight(findAllSql, "\r\n;")
			if queryOpts.OrderBy != nil {
				findAllSql += fmt.Sprintf(" ORDER BY %s", *queryOpts.OrderBy)
			}
			if pager.PageSize != 0 {
				findAllSql += fmt.Sprintf(" LIMIT %d OFFSET %d", pager.PageSize, (pager.Page-1)*pager.PageSize)
			}
			findAllSql += ";"
		} else {
			if queryOpts.OrderBy != nil {
				findAllSql += fmt.Sprintf(" ORDER BY %s", *queryOpts.OrderBy)
			}
		}
	}
	return findMany[T](db, instance, findAllSql)
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

	result, err := findMany[T](db, instance, limitedQuery)
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

func DeleteAll[T model[P], P any](db Database, instance T) (*int64, error) {

	result, err := db.NamedExec(instance.DeleteAllQuery(), instance)

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
	UpdateAllQuery() string
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

func countFields(v any) int {
	return rvCountFields(reflect.ValueOf(v))
}

func rvCountFields(rv reflect.Value) (count int) {

	if rv.Kind() != reflect.Struct {
		return
	}

	fs := rv.NumField()
	count += fs

	for i := 0; i < fs; i++ {
		f := rv.Field(i)
		if rv.Type().Field(i).Anonymous {
			count-- // don't count embedded structs (listed as anonymous fields) as a field
		}

		// recurse each field to see if that field is also an embedded struct
		count += rvCountFields(f)
	}

	return
}

func BulkInsert[T model[P], P any](db *sqlx.DB, ctx context.Context, itemsToSave ...T) ([]*P, error) {
	tx, err := db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return nil, err
	}

	pgMaxParameterCount := 500 // TODO: offer as an option

	// we need to batch the inserts so num `items` * `item` struct field
	// count is less than postgres MaxParameterCount - this is conservative
	firstItem := itemsToSave[0]
	itemPropertyCount := countFields(*firstItem)

	maxBatch := pgMaxParameterCount / itemPropertyCount
	items := make([]*P, 0, len(itemsToSave))
	for i := 0; i < len(itemsToSave); i += maxBatch {
		end := i + maxBatch

		if end > len(itemsToSave) {
			end = len(itemsToSave)
		}

		// Using an anonymous function to ensure that the rows are closed as we go!
		err := func() error {
			firstItemSql := firstItem.InsertQuery()
			rows, err := sqlx.NamedQueryContext(ctx, tx, firstItemSql, itemsToSave[i:end])
			if err != nil {
				tx.Rollback()
				return err
			}
			defer rows.Close()

			insertedRowIdx := 0
			batchItems := make([]*P, len(itemsToSave[i:end]))
			for rows.Next() {
				newItem := new(P)
				if err := rows.StructScan(newItem); err != nil {
					tx.Rollback()
					return err
				}
				batchItems[insertedRowIdx] = newItem
				insertedRowIdx++
			}
			items = append(items, batchItems...)
			return nil
		}()
		if err != nil {
			return nil, err
		}
	}
	tx.Commit()

	return items, nil
}

// *************************
// errors
// *************************

var ErrNotFound = errors.New("entity not found")
var ErrFoundMultiple = errors.New("multiple matching entities")
