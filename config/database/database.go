package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//InitDB is function for initialize database connection
func OpenDB(user, password, dbName string) (*sql.DB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbName))

	if err != nil {
		return nil, fmt.Errorf("Erro ao tentar conectar no banco de dados %v \n", err)
	}

	return db, nil

}

func CloseDB(db *sql.DB) {
	db.Close()
}
