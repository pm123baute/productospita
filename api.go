package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"

	_ "github.com/go-sql-driver/mysql"
)

type Parametro struct {
	producto string
}

var params Parametro

type consulta struct {
	PRD string `json:"producto"`
}

var producto string

func handleRequests(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var con consulta
	err = json.Unmarshal(body, &con)
	if err != nil {
		panic(err)
	}
	producto = con.PRD
	fmt.Printf("Producto: %s\n", producto)
	w.WriteHeader(http.StatusOK)
	consultabd()
}

func main() {

	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)
	http.Handle("/api/consulta/", cors(http.HandlerFunc(handleRequests)))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func consultabd() {
	db, err := sql.Open("mysql", "root:123@tcp(192.168.1.11:3306)/pita")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sql := "SELECT COUNT(*) AS cnt_items FROM productos WHERE nombre Like ?;"
	var count int
	err = db.QueryRow(sql, "%"+producto+"%").Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(count)
}
