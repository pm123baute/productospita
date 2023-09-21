package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, err := sql.Open("mysql", "root:123@tcp(192.168.1.11:3306)/pita")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	item := "Harina"
	query := "SELECT COUNT(*) AS cantidad_items FROM productos WHERE nombre Like ?;"
	var count int
	err = db.QueryRow(query, "%"+item+"%").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total:", count)
}
