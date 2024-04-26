package main

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	type args struct {
		arr1 []User
		arr2 []User
	}
	tests := []struct {
		name string
		args args
		want []User
	}{{
		name: "arr1 0",
		args: args{arr1: []User{}, arr2: []User{{ID: 1, Name: "2", Age: 24}}},
		want: []User{{ID: 1, Name: "2", Age: 24}},
	}, {
		name: "arr2 0",
		args: args{arr2: []User{}, arr1: []User{{ID: 1, Name: "2", Age: 24}}},
		want: []User{{ID: 1, Name: "2", Age: 24}},
	}, {
		name: "arr1 > arr2",
		args: args{arr1: []User{{ID: 3, Name: "212", Age: 214},
			{ID: 8, Name: "21233", Age: 243},
			{ID: 10, Name: "12323", Age: 243}},
			arr2: []User{{ID: 1, Name: "2", Age: 24},
				{ID: 5, Name: "23", Age: 243}}},
		want: []User{{ID: 1, Name: "2", Age: 24},
			{ID: 3, Name: "212", Age: 214},
			{ID: 5, Name: "23", Age: 243},
			{ID: 8, Name: "21233", Age: 243},
			{ID: 10, Name: "12323", Age: 243}},
	}, {
		name: "arr2 > arr1",
		args: args{arr2: []User{{ID: 3, Name: "212", Age: 214},
			{ID: 8, Name: "21233", Age: 243},
			{ID: 10, Name: "12323", Age: 243}},
			arr1: []User{{ID: 1, Name: "2", Age: 24},
				{ID: 5, Name: "23", Age: 243}}},
		want: []User{{ID: 1, Name: "2", Age: 24},
			{ID: 3, Name: "212", Age: 214},
			{ID: 5, Name: "23", Age: 243},
			{ID: 8, Name: "21233", Age: 243},
			{ID: 10, Name: "12323", Age: 243}},
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.arr1, tt.args.arr2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
