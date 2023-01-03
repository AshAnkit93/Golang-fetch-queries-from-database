package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	// connection properties
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "Lotus",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "eagle_yard",
		AllowNativePasswords: true,
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected")

	var i string

	var usd string
	fmt.Println("Enter user ID : ")
	fmt.Scanln(&usd)

	rows, err := db.Query("SELECT * from t_user_details where UserID=?", usd)

	checkErr(err)

	for rows.Next() {
		var roleID, status int
		var UserID, fname, uname, paswrd, dept, desig, email, createBy, createTS string

		err = rows.Scan(&UserID, &fname, &uname, &paswrd, &roleID, &dept, &desig, &email, &status, &createBy, &createTS)
		checkErr(err)

		fmt.Println(UserID, fname, uname, paswrd, roleID, dept, desig, email, status, createBy, createTS)
	}
	defer db.Close()
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
