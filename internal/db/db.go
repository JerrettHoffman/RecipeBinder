package db

import (
	"RecipeBinder/internal"
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// This file serves to wrap the PGX module for DB reads/writes.
// This serves as an abstraction to decouple the implmentation details of PGX from the rest of the code
// Should we need/want to switch DB/DB modules in future, those changes can be made here rather than throughout the app

type postgres struct {
	db *pgxpool.Pool
}

type dbInsertArgs = pgx.NamedArgs

var (
	postgresInstance *postgres
	postgresOnce     sync.Once
)

// TODO: implement better error handling than this maybe retry
// if DB conn fails, we will need to restart the program and try again because "once" function will not be called twice
// panic is probably not the play here
func newPostgres(context context.Context, connString string) (*postgres, error) {
	postgresOnce.Do(func() {
		db, err := pgxpool.New(context, connString)
		if err != nil {
			panic(err)
		}

		postgresInstance = &postgres{db}
	})
	return postgresInstance, nil
}

// where query is valid SQL insert using named arguments and
// args is a struc to map the named arguments from the sql query with data
// ex 'INSERT INTO users (name) VALUES (@userName)'
//
//	{"userName": "John Doe"}
type dbQuery struct {
	query string
	args  dbInsertArgs
}

// Using Query() here rather than Exec() since we need to return the id of the created/found single row
// Ensure input query string has a sql returning command
func (q dbQuery) dbQuerySingleRowReturningId() (internal.ID, error) {
	postgres, err := newPostgres(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return 0, fmt.Errorf("Table connection error: %w", err)
	}

	var id internal.ID

	insertErr := postgres.db.QueryRow(context.Background(), q.query, q.args).Scan(&id)

	if insertErr != nil {
		return 0, fmt.Errorf("Unable to insert row: %w", insertErr)
	}

	return id, nil
}

// Use this function to primarily inser rows in the db when you do not need the ID returned
func (q dbQuery) dbExec() error {
	postgres, err := newPostgres(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return fmt.Errorf("Table connection error: %w", err)
	}

	_, insertErr := postgres.db.Exec(context.Background(), q.query, q.args)

	if insertErr != nil {
		return fmt.Errorf("Unable to insert row: %w", insertErr)
	}

	return nil
}
