package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/Luftalian/writers_app/interfaces/database"
)

type DbHandler struct {
	Conn *sqlx.DB
}

func NewDbHandler() *DbHandler {
	conn, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}
	log.Printf("Connected to Database")
	DbHandler := new(DbHandler)
	DbHandler.Conn = conn
	return DbHandler
}

func (handler *DbHandler) Select(dest interface{}, query string, args ...interface{}) error {
	return handler.Conn.Select(dest, query, args...)
}

func (handler *DbHandler) Exec(query string, args ...interface{}) (database.Result, error) {
	result, err := handler.Conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return DbResult{result}, nil
}

func (handler *DbHandler) Get(dest interface{}, query string, args ...interface{}) error {
	return handler.Conn.Get(dest, query, args...)
}

type DbResult struct {
	Result sql.Result
}

func (r DbResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r DbResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}
