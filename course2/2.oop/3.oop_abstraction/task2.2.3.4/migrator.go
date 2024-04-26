package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Migrator struct {
	db           *sql.DB
	sqLGenerator SQLGenerator
}

func (m Migrator) Migrate(models ...Tabler) error {
	for _, model := range models {
		_, err := m.db.Exec(m.sqLGenerator.CreateTableSQL(model))
		if err != nil {
			return err
		}
		_, err = m.db.Exec(m.sqLGenerator.CreateInsertSQL(model))
		if err != nil {
			return err
		}
	}
	return nil
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	m := &Migrator{db: db, sqLGenerator: sqlGenerator}
	return m
}

/*
// Основная функция
func main() {
	// Подключение к SQLite БД
	db, err := sql.Open("sqlite3", "file:my_database.db?cache=shared&mode=rwc")
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Создание мигратора с использованием вашего SQLGenerator
	migrator := NewMigrator(db, &SQLiteGenerator{})

	// Миграция таблицы User
	if err := migrator.Migrate(User{}); err != nil {
		log.Fatalf("failed to migrate: %v", err)
	}
}
*/
