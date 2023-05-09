package main

import (
	// "myRestAPI/database"
	// "github.com/gin-gonic/gin"
	"database/sql"
	"fmt"
	"log"
	// "net/http"
	// _ "github.com/go-sql-driver/mysql"
	_ "github.com/denisenkom/go-mssqldb"
)

type User struct {
   FirstName      string
   LastName      string
   Address   string
   City       string
  }

type Service struct {
   database *sql.DB
  }
// var DB *sql.DB
var err error

func initDB() (*sql.DB, error) {
   db, err := sql.Open(
		"sqlserver",
		"sqlserver://sa:MyPass@word@127.0.0.1:1439?database=MyDatabase&connection+timeout=30",
	)
   if err != nil {
    return db, err
   }
  
   return db, nil
  }

func main() {

	db, err := initDB()
   if err != nil {
   log.Fatal(err)
   }

   svc := &Service{
   database: db,
   }

   fmt.Println("connected",svc)

   // getData();
   // insertData();
   
}

// func getData() {
//    var fname string
//    var lname string

//    rows, err := DB.Query("SELECT FirstName, LastName FROM Persons", 1)
//    if err != nil {
//       log.Fatal(err)
//    }
//    defer rows.Close()

//    for rows.Next() {
//       err := rows.Scan(&fname, &lname)
//       if err != nil {
//          log.Fatal(err)
//       }
//       fmt.Printf("the name: %s %s \n", fname, lname)
//    }
//    err = rows.Err()
//    if err != nil {
//       log.Fatal(err)
//    }

//    return 
// }

// func (s *Service) getData(id int) ([]User, error) {
//    var users []User
   
//    res, rr:= DB.Query(`SELECT FirstName, LastName FROM Persons`, id)

//    if err != nil {
//       return users, err
//    }

//    for res.Next() {
//       var u User
//       err := res.Scan(&u.FirstName, &u.LastName, u.)
//    }
// }

// func updateData() {
//    //UPDATE DATA
//    res, err := DB.Exec("UPDATE Persons SET PersonID=5 WHERE PersonID IS NULL")
//    if err != nil {
//       log.Fatal(err)
//    }
//    rowCount, err := res.RowsAffected()
//    if err != nil {
//       log.Fatal()
//    }
//    log.Printf("affected = %d \n", rowCount)
// }

func (s *Service) insertData(u User) error{
   //INSERT DATA
   _, err := s.database.Exec(`INSERT INTO Persons(FirstName, LastName, Address, City) VALUES(?, ?, ?, ?)`, 
   u.FirstName,
   u.LastName,
   u.Address,
   u.City)
   if err != nil {
      log.Fatal(err)
   }
   // rowCnt, err := res.RowsAffected()
   // if err != nil {
   //    log.Fatal(err)
   // }
   // log.Printf("affected = %d\n", rowCnt)

   return nil
}

// func delData() {
//    //DELETE DATA
//    res, err := DB.Exec("DELETE FROM Persons WHERE PersonID=5")
//    if err != nil {
//       log.Fatal(err)
//    }
//    rowCnt, err := res.RowsAffected()
//    if err != nil {
//       log.Fatal(err)
//    }
//    log.Printf("affected = %d\n", rowCnt)
// }

// func Add(x, y int) (res int) {
// 	return x + y
// }
