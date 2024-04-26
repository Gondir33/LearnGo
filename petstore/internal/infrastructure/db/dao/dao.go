package dao

import (
	"context"
	"petstore/internal/infrastructure/db/tabler"
	"reflect"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

//go:generate mockgen -source=./dao.go -destination=../mock/dao_mock.go -package=mock
type IfaceDAO interface {
	BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error)
	Create(ctx context.Context, entity tabler.Tabler, opts ...interface{}) error
	List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error
	Update(ctx context.Context, entity tabler.Tabler, condition Condition, opts ...interface{}) error
}

type Condition struct {
	Equal       map[string]interface{}
	NotEqual    map[string]interface{}
	Order       []*Order
	LimitOffset *LimitOffset
	ForUpdate   bool
	Upsert      bool
}

type Order struct {
	Field string
	Asc   bool
}

type LimitOffset struct {
	Offset int64
	Limit  int64
}

type DAO struct {
	db         *sqlx.DB
	sqlBuilder sq.StatementBuilderType
}

func NewDAO(db *sqlx.DB) IfaceDAO {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	return &DAO{db: db, sqlBuilder: builder}
}

func (d *DAO) BuildSelect(tableName string, condition Condition, fields ...string) (string, []interface{}, error) {
	queryRaw := d.sqlBuilder.Select(fields...).From(tableName)
	queryRaw = queryRaw.Where(condition.Equal)
	for _, order := range condition.Order {
		if order.Asc == true {
			queryRaw = queryRaw.OrderByClause(order.Field + " ASC")
		} else {
			queryRaw = queryRaw.OrderByClause(order.Field)
		}
	}
	queryRaw = queryRaw.Limit(uint64(condition.LimitOffset.Limit)).Offset(uint64(condition.LimitOffset.Offset))
	return queryRaw.ToSql()
}

func filterByTag(tag string, tvalue string) func(fields *[]reflect.StructField) {
	return tabler.FilterByTags(map[string]func(value string) bool{
		tag: func(value string) bool {
			return strings.Contains(value, tvalue)
		},
	})
}

func (d *DAO) Create(ctx context.Context, entity tabler.Tabler, opts ...interface{}) error {
	info := tabler.GetStructInfo(entity)
	InsertRaw := d.sqlBuilder.Insert(entity.TableName())
	InsertRaw = InsertRaw.Columns(info.Fields...).Values(info.Pointers...)
	_, err := InsertRaw.RunWith(d.db).Exec()
	return err
}

func (d *DAO) List(ctx context.Context, dest interface{}, table tabler.Tabler, condition Condition, opts ...interface{}) error {
	info := tabler.GetStructInfo(table)
	sql, args, err := d.BuildSelect(table.TableName(), condition, info.Fields...)
	if err != nil {
		return err
	}
	rows, err := d.db.Queryx(sql, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	destType := reflect.TypeOf(dest).Elem().Elem()
	destValue := reflect.ValueOf(dest)
	destSlice := reflect.MakeSlice(reflect.SliceOf(destType), 0, 0)

	for rows.Next() {
		destElem := reflect.New(destType).Interface()

		err := rows.StructScan(destElem)
		if err != nil {
			return err
		}

		destSlice = reflect.Append(destSlice, reflect.ValueOf(destElem).Elem())
	}

	destValue.Elem().Set(destSlice)

	return nil
}

func (d *DAO) Update(ctx context.Context, entity tabler.Tabler, condition Condition, opts ...interface{}) error {
	info := tabler.GetStructInfo(entity)
	UpdateRaw := d.sqlBuilder.Update(entity.TableName())
	UpdateRaw = UpdateRaw.Where(condition.Equal)
	for i := 0; i < len(info.Fields); i++ {
		if info.Fields[i] != "id" {
			UpdateRaw = UpdateRaw.Set(info.Fields[i], info.Pointers[i])
		}
	}
	_, err := UpdateRaw.RunWith(d.db).Exec()
	return err
}
