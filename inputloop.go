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
	i = "yes"
	for {
		if i == "yes" || i == "YES" {

			var usd string
			fmt.Println("Enter user ID : ")
			fmt.Scanln(&usd)

			rows, err := db.Query("SELECT * from t_user_details")

			checkErr(err)

			for rows.Next() {
				var roleID, status int
				var uid, fname, uname, paswrd, dept, desig, email, createBy, createTS string

				err = rows.Scan(&userid, &fname, &uname, &paswrd, &roleID, &dept, &desig, &email, &status, &createBy, &createTS)
				checkErr(err)

				fmt.Println(userid, fname, uname, paswrd, roleID, dept, desig, email, status, createBy, createTS)

				fmt.Println("Continue Quering : ")
				fmt.Scanln(&i)
				if i == "no" || i == "NO" {
					fmt.Println("Exiting")
					break
				}
			}
			defer db.Close()
		}
	}
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
