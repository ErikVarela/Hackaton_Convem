package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

// InitDB cria as tabelas no SQLite e insere os dados de seed se o banco estiver vazio.
func InitDB(db *sql.DB) error {
	// 1. Criar tabela de incidentes
	queryIncidents := `
	CREATE TABLE IF NOT EXISTS incidents (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		severity TEXT NOT NULL,
		owner TEXT NOT NULL,
		status TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		updated_at DATETIME NOT NULL
	);`

	if _, err := db.Exec(queryIncidents); err != nil {
		return fmt.Errorf("erro ao criar tabela de incidentes: %w", err)
	}

	// 2. Criar tabela de historico de status
	queryHistory := `
	CREATE TABLE IF NOT EXISTS incident_history (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		incident_id INTEGER NOT NULL,
		previous_status TEXT NOT NULL,
		new_status TEXT NOT NULL,
		changed_at DATETIME NOT NULL,
		FOREIGN KEY (incident_id) REFERENCES incidents(id) ON DELETE CASCADE
	);`

	if _, err := db.Exec(queryHistory); err != nil {
		return fmt.Errorf("erro ao criar tabela de historico: %w", err)
	}

	// 3. Criar tabela de comentarios (Change Request #1)
	queryComments := `
	CREATE TABLE IF NOT EXISTS incident_comments (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		incident_id INTEGER NOT NULL,
		author TEXT NOT NULL,
		content TEXT NOT NULL,
		created_at DATETIME NOT NULL,
		FOREIGN KEY (incident_id) REFERENCES incidents(id) ON DELETE CASCADE
	);`

	if _, err := db.Exec(queryComments); err != nil {
		return fmt.Errorf("erro ao criar tabela de comentarios: %w", err)
	}

	// 4. Inserir dados de exemplo (Seed) se a tabela estiver vazia
	return seedDatabase(db)
}

func seedDatabase(db *sql.DB) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM incidents").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	now := time.Now()

	seeds := []struct {
		Title       string
		Description string
		Severity    string
		Owner       string
		Status      string
	}{
		{
			Title:       "Payment API instability",
			Description: "A API de pagamentos apresentando alta latencia e falhas intermitentes.",
			Severity:    "Critical",
			Owner:       "Ana",
			Status:      "Open",
		},
		{
			Title:       "Reconciliation delay",
			Description: "Atraso no processamento dos arquivos de reconciliacao financeira.",
			Severity:    "High",
			Owner:       "Bruno",
			Status:      "In Progress",
		},
		{
			Title:       "Incorrect customer notification",
			Description: "Envio duplicado de notificacao de confirmacao de e-mail.",
			Severity:    "Medium",
			Owner:       "Carla",
			Status:      "Resolved",
		},
	}

	for _, s := range seeds {
		res, err := db.Exec(
			`INSERT INTO incidents (title, description, severity, owner, status, created_at, updated_at) 
			 VALUES (?, ?, ?, ?, ?, ?, ?)`,
			s.Title, s.Description, s.Severity, s.Owner, s.Status, now, now,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir seed %s: %w", s.Title, err)
		}

		id, _ := res.LastInsertId()
		_, err = db.Exec(
			`INSERT INTO incident_history (incident_id, previous_status, new_status, changed_at) 
			 VALUES (?, ?, ?, ?)`,
			id, "N/A", s.Status, now,
		)
		if err != nil {
			return fmt.Errorf("erro ao inserir historico inicial para %s: %w", s.Title, err)
		}
	}

	log.Println("Dados iniciais de seed inseridos com sucesso!")
	return nil
}
