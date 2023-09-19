package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/handlers"
)

type Parameters struct {
	Producto  string
}

var params Parameters

type productos struct {
	Producto  string `json:"user"`
}

func handleRequests(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var cnt productos
	err = json.Unmarshal(body, &cnt)
	if err != nil {
		panic(err)
	}
	cantidad:=cnt.Producto
	fmt.Printf("Cantidad: %s\n", cantidad)
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Configura los manejadores de CORS para permitir solicitudes desde todos los orígenes.
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedOrigins([]string{"*"}),
	)

	// Crea el enrutador de HTTP y registra el manejador de rutas "/api".
	http.Handle("/api/pita/consultacntnombre", cors(http.HandlerFunc(handleRequests)))

	// Inicia el servidor en el puerto 8080.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
