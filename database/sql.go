package sqlplus

import (
	"database/sql"
	"time"
)
import _ "github.com/go-sql-driver/mysql"

type DB struct {
	*sql.DB
	driverName  string
	active      int
	idle        int
	idleTimeout int64

	idx int64
}

type Row struct {
	err error
	*sql.Rows
	db    *DB
	query string
	args  []interface{}
}

func Open(driverName, dataSourceName string, active, idle int, idleTimeout int64) *DB {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Duration(idleTimeout))
	db.SetMaxIdleConns(idle)
	db.SetMaxOpenConns(active)
	return &DB{
		DB:          db,
		driverName:  driverName,
		active:      active,
		idle:        idle,
		idleTimeout: idleTimeout,
	}
}

func (db *DB) exec(query string, args ...interface{})(res sql.Result, err error){
	res, err = db.DB.Exec(query, args)
	return
}

func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	row, err := db.query(query, args)
	if err != nil {
		return err
	}
	defer row.Close()
	return nil
}

func (db *DB) Insert() {}

func (db *DB) Delete() {}

func (db *DB) Update() {}

func (db *DB) Begin() {}

func (db *DB) query(query string, args ...interface{}) (*Row, error) {
	rows, err := db.DB.Query(query, args)
	if err != nil {
		return nil, err
	}
	return &Row{
		err:   err,
		Rows:  rows,
		db:    db,
		query: query,
		args:  args,
	}, nil
}

func slowlog() {}
