package dao

import (
	"context"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type SampleEntity struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
}

func (s SampleEntity) TableName() string {
	return "sample_table"
}

func setupTestDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	// Create a sample table for testing
	_, err = db.Exec(`
		CREATE TABLE sample_table (
			id INTEGER PRIMARY KEY,
			name TEXT,
			age INTEGER
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func TestDAO_Create(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dao := NewDAO(db)

	entity := SampleEntity{
		Name: "John",
		Age:  30,
	}

	err = dao.Create(context.Background(), &entity)
	assert.NoError(t, err)

	// Verify that the entity has been inserted
	var count int
	err = db.Get(&count, "SELECT COUNT(*) FROM sample_table")
	assert.NoError(t, err)
	assert.Equal(t, 1, count)
}

func TestDAO_List(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dao := NewDAO(db)

	// Insert some sample data
	_, err = db.Exec("INSERT INTO sample_table (name, age) VALUES (?, ?)", "Alice", 25)
	assert.NoError(t, err)
	_, err = db.Exec("INSERT INTO sample_table (name, age) VALUES (?, ?)", "Bob", 35)
	assert.NoError(t, err)

	var entities []SampleEntity
	err = dao.List(context.Background(), &entities, SampleEntity{}, Condition{}, nil)
	assert.NoError(t, err)

	// Verify that the correct number of entities has been retrieved
	assert.Equal(t, 2, len(entities))
	assert.Equal(t, "Alice", entities[0].Name)
	assert.Equal(t, 25, entities[0].Age)
	assert.Equal(t, "Bob", entities[1].Name)
	assert.Equal(t, 35, entities[1].Age)
}

func TestDAO_Update(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	dao := NewDAO(db)

	// Insert a sample entity
	_, err = db.Exec("INSERT INTO sample_table (name, age) VALUES (?, ?)", "Alice", 25)
	assert.NoError(t, err)

	entity := SampleEntity{
		Name: "Bob",
		Age:  30,
	}

	// Update the entity
	err = dao.Update(context.Background(), &entity, Condition{Equal: map[string]interface{}{"id": 1}}, nil)
	assert.NoError(t, err)

	// Verify that the entity has been updated
	var updatedEntity SampleEntity
	err = db.Get(&updatedEntity, "SELECT * FROM sample_table WHERE id = 1")
	assert.NoError(t, err)
	assert.Equal(t, "Bob", updatedEntity.Name)
	assert.Equal(t, 30, updatedEntity.Age)
}
