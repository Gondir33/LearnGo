package main

import (
	"database/sql"
	"log"
	"reflect"
	"testing"
)

func TestMigrator(t *testing.T) {
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	migrator := NewMigrator(db, &SQLiteGenerator{})
	if !reflect.DeepEqual(migrator, &Migrator{db, &SQLiteGenerator{}}) {
		t.Errorf("NewMigrator don't work: want %v, get: %v", &Migrator{db, &SQLiteGenerator{}}, migrator)
	}

	fake := &GoFakeitGenerator{}
	fakeUser := fake.GenerateFakeUser()
	if err := migrator.Migrate(fakeUser); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
	if err := migrator.Migrate(fakeUser); err == nil {
		log.Fatalf("should failed but not %v", err)
	}
	db.Close()
	if err := migrator.Migrate(fakeUser); err == nil {
		log.Fatalf("should failed but not %v", err)
	}

}
