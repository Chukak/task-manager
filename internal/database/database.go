package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
)

type QueryResult pgx.Rows

// QueryManage is a common interface for SQL queries
type QueryManage interface {
	Exec(string, ...interface{})
}

// Query pbject works with SQL
type Query struct {
	connectionPool *pgxpool.Pool
}

// DatabaseManage is a common interface for databases (PSQL)
type DatabaseManage interface {
	Open()
	Close()
}

// Database object, stores host, port, etc for connection to database
type Database struct {
	config *pgxpool.Config
	query  *Query
}

// NewDatabase returns a new Database object
func NewDatabase(host string, port uint16, db string, user string, password string) (*Database, error) {
	// user:password@host:port/db
	connectionString := "postgres://%s:%s@%s:%v/%s"
	connectionString = fmt.Sprintf(connectionString, user, password, host, port, db)

	config, err := pgxpool.ParseConfig(connectionString)
	return &Database{config, nil}, err
}

// Open the database, returns status and error
func (d *Database) Open() (bool, error) {
	var err error
	if d.config == nil {
		err = errors.New("invalid config. Database will not open")
	} else {
		pool, err := pgxpool.ConnectConfig(context.Background(), d.config)
		if err == nil {
			d.query = &Query{connectionPool: pool}
		}
	}
	return err == nil, err
}

// Close the database, returns status and error
func (d *Database) Close() (bool, error) {
	var err error
	if d.query == nil {
		err = errors.New("database is not open")
	}
	d.query.connectionPool.Close()
	return err == nil, err
}

// Exec a SQL query using database object
func (d *Database) Exec(query string, args ...interface{}) (QueryResult, error) {
	if d.query == nil {
		var err error
		var rows QueryResult
		err = errors.New("database is not open")
		return rows, err
	}
	conn, _ := d.query.connectionPool.Acquire(context.Background())
	defer conn.Release()
	return conn.Query(context.Background(), query, args...)
}
