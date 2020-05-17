package database

import (
	"os"
	"strconv"
	"testing"

	test "github.com/chukak/task-manager/pkg/checks"
)

func TestDatabaseInitialization(t *testing.T) {
	test.SetT(t)

	host := os.Getenv("DB_HOST")

	val, err := strconv.Atoi(os.Getenv("DB_PORT"))
	test.CheckEqual(err, nil)

	port := uint16(val)
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	db, err := NewDatabase(host, port, database, user, password)
	test.CheckEqual(err, nil)
	test.CheckEqual(db.config.ConnConfig.Host, host)
	test.CheckEqual(db.config.ConnConfig.Port, port)
	test.CheckEqual(db.config.ConnConfig.Database, database)
	test.CheckEqual(db.config.ConnConfig.User, user)
	test.CheckEqual(db.config.ConnConfig.Password, password)
}

func TestDatabaseFunctionality(t *testing.T) {
	test.SetT(t)

	host := os.Getenv("DB_HOST")
	val, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	port := uint16(val)
	database := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	db, _ := NewDatabase(host, port, database, user, password)
	ok, err := db.Open()
	test.CheckTrue(ok)
	test.CheckEqual(err, nil)

	conn, err := db.GetConnection()
	test.CheckEqual(err, nil)
	rows, err := db.Exec(SELECT, conn,
		"SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';")
	test.CheckEqual(err, nil)
	// two tables
	test.CheckTrue(rows.Next())
	test.CheckEqual(rows.Err(), nil)
	var names [2]string
	test.CheckEqual(rows.Scan(&names[0]), nil)
	test.CheckEqual(rows.Scan(&names[1]), nil)
	test.CheckNotEqual(len(names[0]), 0)
	test.CheckNotEqual(len(names[1]), 0)
	db.CloseConnection(conn)

	ok, err = db.Close()
	test.CheckTrue(ok)
	test.CheckEqual(err, nil)
}

func TestDatabaseInvalidConfig(t *testing.T) {
	test.SetT(t)

	host := os.Getenv("DB_HOST")
	port := 0
	database := os.Getenv("DB_NAME")
	user := "user"
	password := os.Getenv("DB_PASSWORD")

	db, _ := NewDatabase(host, uint16(port), database, user, password)
	ok, err := db.Open()
	test.CheckFalse(ok)
	test.CheckNotEqual(err, nil)

	conn, err := db.GetConnection()
	test.CheckNotEqual(err, nil)
	rows, err := db.Exec(SELECT, conn,
		"SELECT table_name FROM information_schema.tables WHERE table_schema = 'public';")
	test.CheckNotEqual(err, nil)
	test.CheckEqual(err.Error(), "connection is null")
	db.CloseConnection(conn)
	// nothing to do, otherwise SEGFAULT!
	_ = rows
}
