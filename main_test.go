package main

import (
    // "fmt"
	"database/sql"
	"testing"
	"log"
	"os"
    // "time"
	_ "github.com/denisenkom/go-mssqldb"
    // r "github.com/moemoe89/go-unit-test-sql/repository"
    // "github.com/google/uuid"
    "github.com/DATA-DOG/go-sqlmock"
    // "github.com/stretchr/testify/assert"
)

const (
	dbDriver = "sqlDriver"
	dbSource = "sqlserver://sa:MyPass@word@127.0.0.1:1439?database=MyDatabase&DBection+timeout=30" 
)

func TestMain(m *testing.M) {
	DB, err := sql.Open(
		"sqlserver",
		"sqlserver://sa:MyPass@word@127.0.0.1:1439?database=MyDatabase&DBection+timeout=30",
	)
    if err != nil {
        log.Fatal("cannot DBect to db:", err)
    }

    pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

    os.Exit(m.Run())
}

func TestInsertData(t *testing.T) {
	DB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
    if err != nil {
        log.Fatal("error init mock", err)
    }
    defer DB.Close()

    svc := Service{
        database: DB,
    }

    userSuccess := User{
        FirstName: "Fadhi",
        LastName: "Muhammad",
        Address: "Klang",
        City: "Selangor",
    }

    mock.ExpectExec(`INSERT INTO Persons(FirstName, LastName, Address, City) VALUES(?, ?, ?, ?)`).WithArgs(userSuccess.FirstName, userSuccess.LastName, userSuccess.Address, userSuccess.City).
    WillReturnResult(sqlmock.NewResult(1,1))
    err = svc.insertData(userSuccess)

    if err != nil {
        t.Errorf("user success, got  err %v", err)
    }
}

// func TestAdd(t *testing.T){

//     got := Add(4, 6)
//     want := 10

//     if got != want {
//         t.Errorf("got %q, wanted %q", got, want)
//     }
// }