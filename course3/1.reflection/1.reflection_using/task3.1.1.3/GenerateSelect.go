package main

import (
	sq "github.com/Masterminds/squirrel"
)

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
	sqlBuilder sq.StatementBuilderType
}

func NewDAO() *DAO {
	builder := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)
	return &DAO{sqlBuilder: builder}
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

/*
func main() {
	d := NewDAO()
	s, _, err := d.BuildSelect("users", Condition{
		Equal: map[string]interface{}{
			"username": "test",
		},
		LimitOffset: &LimitOffset{
			Offset: 5,
			Limit:  3,
		},
		Order: []*Order{
			{
				Field: "id",
				Asc:   true,
			},
		},
	}, "id", "username")

	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}
*/
