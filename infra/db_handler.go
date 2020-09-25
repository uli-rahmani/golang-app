package infra

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"saturday/constants"
	"saturday/domain"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// DBHandler - Database struct.
type DBHandler struct {
	DB *sqlx.DB
}

// ConnectDB - function for connect DB.
func (d *DBHandler) ConnectDB(dbAcc *domain.DBAccount) {
	dbs, err := sqlx.Open("postgres", "user="+dbAcc.Username+" password="+dbAcc.Password+" dbname="+dbAcc.DBName+" host="+dbAcc.URL+" port="+dbAcc.Port+" connect_timeout="+dbAcc.Timeout)
	if err != nil {
		log.Println(constants.ConnectDBFail + " | " + err.Error())
	}

	d.DB = dbs

	err = d.DB.Ping()
	if err != nil {
		fmt.Printf("postgres username: %s, password: %s, url: %s, port: %s, dbname: %s", dbAcc.Username, dbAcc.Password, dbAcc.URL, dbAcc.Port, dbAcc.DBName)
		log.Println(constants.ConnectDBFail, err.Error())
	}

	log.Println(constants.ConnectDBReadSuccess)
	d.DB.SetConnMaxLifetime(time.Duration(dbAcc.MaxLifeTime))
}

// Close - function for connection lost.
func (d *DBHandler) Close() {
	if err := d.DB.Close(); err != nil {
		log.Println(constants.ClosingDBReadFailed + " | " + err.Error())
	} else {
		log.Println(constants.ClosingDBReadSuccess)
	}
}

func (d *DBHandler) Exec(query string, args ...interface{}) (sql.Result, error) {
	result, err := d.DB.Exec(query, args...)
	return result, err
}

func (d *DBHandler) Select(dest interface{}, query string, args ...interface{}) error {
	err := d.DB.Select(dest, query, args...)
	return err
}

func (d *DBHandler) Rebind(query string) string {
	return d.DB.Rebind(query)
}

func (d *DBHandler) In(query string, params ...interface{}) (string, []interface{}, error) {
	query, args, err := sqlx.In(query, params...)
	return query, args, err
}
