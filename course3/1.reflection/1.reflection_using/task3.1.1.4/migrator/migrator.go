package migrator

import (
	"database/sql"
	"test/tabler"

	"golang.org/x/sync/errgroup"
)

type Migratorer interface {
	Migrate(tables ...func(tabler tabler.Tabler)) error
}

type Migrator struct {
	db           *sql.DB
	sqlGenerator SQLGenerator
}

func NewMigrator(db *sql.DB, sqlGenerator SQLGenerator) *Migrator {
	return &Migrator{
		db:           db,
		sqlGenerator: sqlGenerator,
	}
}

func (m *Migrator) Migrate(tables ...tabler.Tabler) error {
	var errGroup errgroup.Group
	for _, table := range tables {
		createSQL := m.sqlGenerator.CreateTableSQL(table)
		errGroup.Go(func() error {
			_, err := m.db.Exec(createSQL)
			return err
		})
	}

	return errGroup.Wait()
}
