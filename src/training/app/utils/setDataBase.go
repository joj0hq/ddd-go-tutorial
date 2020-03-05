package utils

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func OpenDataBase() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	var (
		DBMS     = os.Getenv("DBMS")
		USER     = os.Getenv("DB_USER")
		PASS     = os.Getenv("DB_PASS")
		PORT     = os.Getenv("DB_PORT")
		HOST     = os.Getenv("HOST")
		PROTOCOL = "tcp(" + HOST + ":" + PORT + ")"
		DBNAME   = os.Getenv("DB_NAME")
		CONNECT  = USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
	)
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil, err
	}
	return db, nil
}
