package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "modernc.org/sqlite"
)

func main() {
	db, err := sql.Open("sqlite", "./incidents.db")
	if err != nil {
		log.Fatalf("Erro ao abrir banco de dados: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}

	if err := InitDB(db); err != nil {
		log.Fatalf("Erro ao inicializar banco de dados: %v", err)
	}

	app := &Application{DB: db}

	// Servir a interface web estatica (HTML/CSS/JS)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	// Endpoints da API HTTP
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/api/dashboard", app.GetDashboardHandler)
	http.HandleFunc("/api/incidents", app.ListIncidentsHandler)
	http.HandleFunc("/api/incidents/create", app.CreateIncidentHandler)
	http.HandleFunc("/api/incidents/detail", app.GetIncidentDetailsHandler)
	http.HandleFunc("/api/incidents/update-status", app.UpdateStatusHandler)
	http.HandleFunc("/api/incidents/comment", app.AddCommentHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro ao iniciar servidor: %v", err)
	}
}
