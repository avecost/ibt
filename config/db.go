package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB(dataSourceName string) {
	var err error

	DB, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	if err = DB.Ping(); err != nil {
		log.Panic(err)
	}
}

func DBPrepareStatement(stmtStr string) (*sql.Stmt, error) {
	stmt, err := DB.Prepare(stmtStr)
	if err != nil {
		return nil, err
	}
	return stmt, nil
}

func DBCloseStatement(stmt *sql.Stmt) {
	stmt.Close()
}
