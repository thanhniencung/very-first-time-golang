package model

import (
        "database/sql"
        "fmt"
        _ "github.com/lib/pq"
    )

const (  
  user     = "postgres"
  password = "123456"
  dbname   = "University"
)

var db *sql.DB

func InitDB() {
    var err error
    
    psqlInfo := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)
    db, err = sql.Open("postgres", psqlInfo)
    
    if err != nil {  
      panic(err)
    }

    if err = db.Ping(); err != nil {
        panic(err)
    }

    fmt.Println("========= You connected to your database. =========")
}


