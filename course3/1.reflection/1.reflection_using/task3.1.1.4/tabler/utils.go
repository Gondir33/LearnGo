package tabler

import (
	"reflect"
)

type Tabler interface {
	TableName() string
}

type StructInfo struct {
	Fields   []string
	Pointers []interface{}
}

func GetStructInfo(u interface{}, args ...func(*[]reflect.StructField)) StructInfo {
	val := reflect.ValueOf(u).Elem()

	structFields := reflect.VisibleFields(reflect.TypeOf(u).Elem())
	for i := range args {
		if args[i] == nil {
			continue
		}
		args[i](&structFields)
	}

	var res StructInfo

	for _, field := range structFields {
		valueField := val.FieldByName(field.Name)
		res.Pointers = append(res.Pointers, valueField.Addr().Interface())
		res.Fields = append(res.Fields, field.Tag.Get("db"))
	}

	return res
}

func FilterByFields(fieldsint ...int) func(fields *[]reflect.StructField) {
	return func(fields *[]reflect.StructField) {
		var res []reflect.StructField

		for _, field := range fieldsint {
			res = append(res, (*fields)[field])
		}

		*fields = res
	}
}

func FilterByTags(tags map[string]func(value string) bool) func(fields *[]reflect.StructField) {
	return func(fields *[]reflect.StructField) {
		var res []reflect.StructField

		for key, value := range tags {
			for _, field := range *fields {
				if value(field.Tag.Get(key)) == true {
					res = append(res, field)
				}
			}
		}

		*fields = res
	}
}
