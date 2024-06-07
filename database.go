package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB

func initDB() {
    var err error
    dsn := "root:@tcp(127.0.0.1:3306)/mini-projet" 
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Connexion a la base de donnée réussie")
}
