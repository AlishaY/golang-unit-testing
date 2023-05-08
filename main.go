package main

import (
	// "myRestAPI/database"
	// "github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"log"
	// "net/http"

	// "net/http"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
)

type Service struct {
   database *sql.DB
  }
var DB *sql.DB
var err error

func main() {
   // fmt.Println("Listening on port 3000")
   // http.ListenAndServe(":3000", nil)

	DB, err = sql.Open(
		"sqlserver",
		"sqlserver://sa:MyPass@word@127.0.0.1:1439?database=MyDatabase&connection+timeout=30",
	)

	if err != nil {
		log.Fatal(err)
	}

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Docker SQL Connected!")

   getData();
   
}

func getData() {
   var fname string
   var lname string

   rows, err := DB.Query("SELECT FirstName, LastName FROM Persons", 1)
   if err != nil {
      log.Fatal(err)
   }
   defer rows.Close()

   for rows.Next() {
      err := rows.Scan(&fname, &lname)
      if err != nil {
         log.Fatal(err)
      }
      fmt.Printf("the name: %s %s \n", fname, lname)
   }
   err = rows.Err()
   if err != nil {
      log.Fatal(err)
   }

   return 
}

func updateData() {
   //UPDATE DATA
   res, err := DB.Exec("UPDATE Persons SET PersonID=5 WHERE PersonID IS NULL")
   if err != nil {
      log.Fatal(err)
   }
   rowCount, err := res.RowsAffected()
   if err != nil {
      log.Fatal()
   }
   log.Printf("affected = %d \n", rowCount)
}

func (s *Service) insertData(u User) {
   //INSERT DATA
   res, err := DB.Exec("INSERT INTO Persons(FirstName, LastName, Address, City) VALUES('Harrith', 'Hasmadi', 'Kempadang', 'Penang')")
   if err != nil {
      log.Fatal(err)
   }
   rowCnt, err := res.RowsAffected()
   if err != nil {
      log.Fatal(err)
   }
   log.Printf("affected = %d\n", rowCnt)
}

func delData() {
   //DELETE DATA
   res, err := DB.Exec("DELETE FROM Persons WHERE PersonID=5")
   if err != nil {
      log.Fatal(err)
   }
   rowCnt, err := res.RowsAffected()
   if err != nil {
      log.Fatal(err)
   }
   log.Printf("affected = %d\n", rowCnt)
}

func Add(x, y int) (res int) {
	return x + y
}
