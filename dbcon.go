package DBUtils

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
)

func GetDbHandle() (*sql.DB, error) {
	fmt.Sprintf("Opening connection to %s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"))
	cfg := mysql.Config{
		User:   "root",
		Passwd: "XXX",
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", "localhost", "3306"),
		DBName: "test",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	return db, err
}
