package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"gotest.tools/assert"

	_ "github.com/lib/pq"
)

const (
	port   = 5432
	user   = "postgres"
	dbname = "mytest"
)

func TestQuery(t *testing.T) {
	db := getDBConn()
	defer db.Close()

	id := -1
	insertSQL := `INSERT INTO testing.my_table (my_text_value) VALUES ($1) RETURNING id`
	err := db.QueryRow(insertSQL, "Lorem Ipsum").Scan(&id)
	if err != nil {
		panic(err)
	}

	deletedID := -10
	deleteSQL := `DELETE FROM testing.my_table WHERE id = $1 RETURNING id;`
	err = db.QueryRow(deleteSQL, id).Scan(&deletedID)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, id, deletedID)
}

func getDBConn() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"), port, user, os.Getenv("DB_PASS"), dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
