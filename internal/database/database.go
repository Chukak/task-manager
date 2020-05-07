package database

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx"
)

// QueryManage is a common interface for SQL queries
type QueryManage interface {
	Exec(string, ...interface{})
}

// Query pbject works with SQL
type Query struct {
	connection *pgx.Conn
}

// DatabaseManage is a common interface for databases (PSQL)
type DatabaseManage interface {
	Open()
	Close()
}

// Database object, stores host, port, etc for connection to database
type Database struct {
	config *pgx.ConnConfig
	query  *Query
}

// NewDatabase returns a new Database object
func NewDatabase(host string, port uint16, db string, user string, password string) (*Database, error) {
	// user:password@host:port/db
	connectionString := "postgres://%s:%s@%s:%v/%s"
	connectionString = fmt.Sprintf(connectionString, user, password, host, port, db)

	config, err := pgx.ParseConfig(connectionString)
	return &Database{config, nil}, err
}

// Open the database, returns status and error
func (d *Database) Open() (bool, error) {
	var err error
	if d.config == nil {
		err = errors.New("invalid config. Database will not open")
	} else {
		conn, err := pgx.ConnectConfig(context.Background(), d.config)
		if err == nil {
			d.query = &Query{connection: conn}
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
	err = d.query.connection.Close(context.Background())
	return err == nil, err
}

// Exec a SQL query using database object
func (d *Database) Exec(query string, args ...interface{}) (pgx.Rows, error) {
	if d.query == nil {
		var err error
		var rows pgx.Rows
		err = errors.New("database is not open")
		return rows, err
	}
	return d.query.connection.Query(context.Background(), query, args...)
}
