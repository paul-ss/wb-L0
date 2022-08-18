package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/paul-ss/wb-L0/service/config"
	"os"
)

var (
	conn *PgConn
)

func NewPgConn() *PgConn {
	if conn != nil && conn.db != nil && conn.db.Ping() == nil {
		return conn
	}

	db, err := sql.Open("postgres", config.DBConnString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	conn = &PgConn{db}
	return conn
}

func Close() {
	if conn != nil && conn.db != nil {
		err := conn.db.Close()
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}

		conn = nil
	}
}

type PgConn struct {
	db *sql.DB
}

func (pg *PgConn) StoreOrder(id string, data []byte) error {
	_, err := pg.db.Exec(
		"insert into orders (id, data) "+
			"values ($1, $2) ",
		id, data)

	return err
}

func (pg *PgConn) GetOrderById(id string) ([]byte, error) {
	var order []byte

	err := pg.db.QueryRow(
		"select data from orders "+
			"where id = $1 ",
		id).Scan(&order)

	return order, err
}
