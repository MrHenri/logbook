package main

import (
	"log"
	"net/http"

	"github.com/logbook/web/petCRUD/functions"
	"github.com/logbook/web/petCRUD/server"
)

func main() {
	log.Println("Server started on: http://localhost:8080")
	http.HandleFunc("/", functions.Index)
	http.HandleFunc("/show", functions.Show)
	http.HandleFunc("/new", functions.New)
	http.HandleFunc("/edit", functions.Edit)
	http.HandleFunc("/insert", server.Create)
	http.HandleFunc("/update", server.Update)
	http.HandleFunc("/delete", server.Delete)
	http.ListenAndServe(":8080", nil)
}
