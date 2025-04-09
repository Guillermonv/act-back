package main

import (
	"database/sql"
	"log"
	"net/http"

	"act-back/internal/service"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "example:example@tcp(localhost:3306)/example")
	if err != nil {
		log.Fatal("Error abriendo la base de datos:", err)
	}

	srv := &service.ActivityService{DB: db}

	http.HandleFunc("/activities", srv.HandleGetActivitiesGrouped)
	log.Println("Servidor corriendo en http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
