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
		log.Fatal("Error conectando a la base de datos:", err)
	}

	s := &service.ActivityService{DB: db}

	http.HandleFunc("/activities", s.HandleGetActivities)
	http.HandleFunc("/activities/grouped", s.HandleGetActivitiesGrouped)
	http.HandleFunc("/activities/update", s.HandleUpdateActivity) // o método PUT en /activities si preferís
	http.HandleFunc("/activities/populate", s.HandlePopulateActivities)

	// fallback para rutas no existentes

	log.Println("Servidor escuchando en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
