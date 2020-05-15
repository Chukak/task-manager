package database

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/pgxpool"
)

// QueryResult type result of SQL query
type QueryResult pgx.Rows

// QueryExecType implements a SQL query type
type QueryExecType int

const (
	// SELECT query
	SELECT QueryExecType = 0
	// INSERT query
	INSERT QueryExecType = 1
	// UPDATE query
	UPDATE QueryExecType = 2
	// DELETE query
	DELETE QueryExecType = 3
)

// QueryManage is a common interface for SQL queries
type QueryManage interface {
	GetConnection()
	CloseConnection(*pgxpool.Conn)
	Exec(QueryExecType, *pgxpool.Conn, string, ...interface{})
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
	if config != nil && err == nil {
		config.MaxConnLifetime = 7 * time.Second
		config.MaxConns = 20
	}
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

// GetConnection get connection from pool
func (d *Database) GetConnection() (*pgxpool.Conn, error) {
	if d.query == nil {
		var err error = errors.New("database is not open")
		return nil, err
	}
	return d.query.connectionPool.Acquire(context.Background())
}

// CloseConnection close a connection
func (d *Database) CloseConnection(conn *pgxpool.Conn) {
	if conn != nil {
		conn.Release()
	}
}

// Exec a SQL query using database object
func (d *Database) Exec(tp QueryExecType, conn *pgxpool.Conn, query string, args ...interface{}) (QueryResult, error) {
	var rows QueryResult
	var err error
	if conn == nil {
		err = errors.New("connection is null")
		return rows, err
	}
	switch tp {
	case SELECT:
		rows, err = conn.Query(context.Background(), query, args...)
	default:
		_, err = conn.Exec(context.Background(), query, args...)
	}
	return rows, err
}
