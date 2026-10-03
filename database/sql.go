package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func ConnectDB()*sql.DB{
	dsn:="postgres://postgres:aditya@localhost:5432/goapp"
	db,err:=sql.Open("pgx",dsn)
	if err!=nil{
		log.Fatal(err)
	}

	err=db.Ping()
	if err!=nil{
		log.Fatal(err)
	}

	fmt.Println("Database Connected")
	return db

}