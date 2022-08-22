package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/paul-ss/wb-L0/service/config"
	log "github.com/sirupsen/logrus"
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
			log.Error("db close: ", err.Error())
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

func (pg *PgConn) UpdateLastMsgId(sub string, id uint64) error {
	_, err := pg.db.Exec(
		"insert into messages (subId, lastMsgId) "+
			"values ($1, $2) "+
			"on conflict (subId) "+
			"do update "+
			"set lastMsgId = $2 ", sub, id)
	return err
}

func (pg *PgConn) GetLastMsgId(sub string) (id uint64, ok bool) {
	if err := pg.db.QueryRow(
		"select lastMsgId from messages "+
			"where subId = $1", sub).Scan(&id); err != nil {
		return
	}

	ok = true
	return
}
