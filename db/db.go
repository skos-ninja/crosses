package db

import (
	"context"
	"database/sql"

	// Imported to use the postgres sql driver
	_ "github.com/lib/pq"
)

// DB is the export database service
type DB struct {
	conn *sql.DB
}

// Config is the database service configuration object
type Config struct {
	ConnectionString string `json:"conn_string"`
}

// NewDB creates a new DB service with a connection to postgres
func NewDB(config Config) (*DB, error) {
	dbConn, err := sql.Open("postgres", config.ConnectionString)
	if err != nil {
		return nil, err
	}

	return &DB{
		conn: dbConn,
	}, nil
}

// Ping will ping the database
func (db *DB) Ping(ctx context.Context) error {
	return db.conn.PingContext(ctx)
}

func (db *DB) Exec(ctx context.Context, sql string, args ...interface{}) (sql.Result, error) {
	return db.conn.ExecContext(ctx, sql, args...)
}

func (db *DB) QueryWithCtx(ctx context.Context, sql string, args ...interface{}) (*sql.Rows, error) {
	return db.conn.QueryContext(ctx, sql, args...)
}

func (db *DB) QueryRowWithCtx(ctx context.Context, sql string, args ...interface{}) *sql.Row {
	return db.conn.QueryRowContext(ctx, sql, args...)
}
