package main

import (
	"reflect"
	"testing"
)

func TestUser_TableName(t *testing.T) {
	type fields struct {
		ID        int
		FirstName string
		LastName  string
		Email     string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{{
		name:   "users",
		fields: fields{},
		want:   "users",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:        tt.fields.ID,
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
				Email:     tt.fields.Email,
			}
			if got := u.TableName(); got != tt.want {
				t.Errorf("User.TableName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	type args struct {
		tabler Tabler
	}
	fakeDataGenerator := &GoFakeitGenerator{}
	tests := []struct {
		name   string
		sqlite *SQLiteGenerator
		args   args
		want   string
	}{{
		name:   "create",
		sqlite: &SQLiteGenerator{},
		args:   args{fakeDataGenerator.GenerateFakeUser()},
		want:   "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, first_name VARCHAR(100), last_name VARCHAR(100), email VARCHAR(100) UNIQUE)",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlite := &SQLiteGenerator{}
			if got := sqlite.CreateTableSQL(tt.args.tabler); got != tt.want {
				t.Errorf("SQLiteGenerator.CreateTableSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	type args struct {
		model Tabler
	}
	tests := []struct {
		name   string
		sqlite *SQLiteGenerator
		args   args
		want   string
	}{{
		name:   "1st",
		sqlite: &SQLiteGenerator{},
		args:   args{User{ID: 46383, FirstName: "Benton", LastName: "Jewess", Email: "bufordgleichner@herman.io"}},
		want:   "INSERT INTO users (id, first_name, last_name, email) VALUES (46383, 'Benton', 'Jewess', 'bufordgleichner@herman.io')",
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sqlite := &SQLiteGenerator{}
			if got := sqlite.CreateInsertSQL(tt.args.model); got != tt.want {
				t.Errorf("SQLiteGenerator.CreateInsertSQL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fake := &GoFakeitGenerator{}
	got := fake.GenerateFakeUser()
	tests := []struct {
		name string
		fake *GoFakeitGenerator
		want User
	}{{
		name: "fake",
		fake: &GoFakeitGenerator{},
		want: got,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GoFakeitGenerator.GenerateFakeUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
