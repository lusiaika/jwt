package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"tugasregisterjwt/database"
	"tugasregisterjwt/handler"

	"github.com/gorilla/mux"
)

// Replace with your own connection parameters
var sqlserver = "localhost"
var sqlport = 1433
var sqldbName = "register"
var sqluser = "lusiaika"
var sqlpassword = "123456"

func main() {
	// Create connection string
	// connString := fmt.Sprintf("server=%s;database=%s;port=%d;trusted_connection=yes",
	// 	sqlserver, sqldbName, sqlport)

	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		sqluser, sqlpassword, sqlserver, sqlport, sqldbName)

	sql := database.NewSqlConnection(connString)
	handler.SqlConnect = sql
	r := mux.NewRouter()

	handler.InstallUserAPI(r)

	r.Use()
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//http.ListenAndServe(PORT, nil)
	log.Fatal(srv.ListenAndServe())
}
