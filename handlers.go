package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type Application struct {
	DB *sql.DB
}

type DashboardDTO struct {
	OpenIncidents           int `json:"open_incidents"`
	CriticalUnresolvedCount int `json:"critical_unresolved"`
	ResolvedIncidents       int `json:"resolved_incidents"`
}

func (app *Application) GetDashboardHandler(w http.ResponseWriter, r *http.Request) {
	var dash DashboardDTO

	_ = app.DB.QueryRow("SELECT COUNT(*) FROM incidents WHERE status != 'Resolved'").Scan(&dash.OpenIncidents)
	_ = app.DB.QueryRow("SELECT COUNT(*) FROM incidents WHERE severity = 'Critical' AND status != 'Resolved'").Scan(&dash.CriticalUnresolvedCount)
	_ = app.DB.QueryRow("SELECT COUNT(*) FROM incidents WHERE status = 'Resolved'").Scan(&dash.ResolvedIncidents)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dash)
}

func (app *Application) ListIncidentsHandler(w http.ResponseWriter, r *http.Request) {
	statusFilter := r.URL.Query().Get("status")
	severityFilter := r.URL.Query().Get("severity")

	query := "SELECT id, title, description, severity, owner, status, created_at, updated_at FROM incidents WHERE 1=1"
	var args []interface{}

	if statusFilter != "" {
		query += " AND status = ?"
		args = append(args, statusFilter)
	}
	if severityFilter != "" {
		query += " AND severity = ?"
		args = append(args, severityFilter)
	}
	query += " ORDER BY created_at DESC"

	rows, err := app.DB.Query(query, args...)
	if err != nil {
		http.Error(w, "Erro ao buscar incidentes", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var incidents []Incident
	for rows.Next() {
		var inc Incident
		if err := rows.Scan(&inc.ID, &inc.Title, &inc.Description, &inc.Severity, &inc.Owner, &inc.Status, &inc.CreatedAt, &inc.UpdatedAt); err == nil {
			incidents = append(incidents, inc)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(incidents)
}

func (app *Application) CreateIncidentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Severity    string `json:"severity"`
		Owner       string `json:"owner"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" || req.Description == "" || req.Severity == "" || req.Owner == "" {
		http.Error(w, "Campos obrigatórios ausentes: title, description, severity e owner", http.StatusBadRequest)
		return
	}

	now := time.Now()
	res, err := app.DB.Exec(
		"INSERT INTO incidents (title, description, severity, owner, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)",
		req.Title, req.Description, req.Severity, req.Owner, StatusOpen, now, now,
	)
	if err != nil {
		http.Error(w, "Erro ao salvar incidente", http.StatusInternalServerError)
		return
	}

	id, _ := res.LastInsertId()
	_, _ = app.DB.Exec("INSERT INTO incident_history (incident_id, previous_status, new_status, changed_at) VALUES (?, ?, ?, ?)", id, "N/A", StatusOpen, now)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id, "message": "Incidente criado com sucesso"})
}

func (app *Application) UpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost && r.Method != http.MethodPut {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID do incidente inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		NewStatus string `json:"new_status"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.NewStatus == "" {
		http.Error(w, "Novo status não informado", http.StatusBadRequest)
		return
	}

	var inc Incident
	err = app.DB.QueryRow("SELECT id, severity, status FROM incidents WHERE id = ?", id).Scan(&inc.ID, &inc.Severity, &inc.Status)
	if err != nil {
		http.Error(w, "Incidente não encontrado", http.StatusNotFound)
		return
	}

	if err := ValidateStatusTransition(inc.Severity, inc.Status, req.NewStatus); err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	now := time.Now()
	_, err = app.DB.Exec("UPDATE incidents SET status = ?, updated_at = ? WHERE id = ?", req.NewStatus, now, id)
	if err != nil {
		http.Error(w, "Erro ao atualizar status", http.StatusInternalServerError)
		return
	}

	_, _ = app.DB.Exec("INSERT INTO incident_history (incident_id, previous_status, new_status, changed_at) VALUES (?, ?, ?, ?)", id, inc.Status, req.NewStatus, now)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Status atualizado com sucesso"})
}

func (app *Application) AddCommentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var req struct {
		Author  string `json:"author"`
		Content string `json:"content"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Requisição inválida", http.StatusBadRequest)
		return
	}

	if err := ValidateComment(req.Author, req.Content); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	now := time.Now()
	_, err = app.DB.Exec(
		"INSERT INTO incident_comments (incident_id, author, content, created_at) VALUES (?, ?, ?, ?)",
		id, req.Author, req.Content, now,
	)
	if err != nil {
		http.Error(w, "Erro ao salvar comentário", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Comentário adicionado com sucesso"})
}

func (app *Application) GetIncidentDetailsHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var inc Incident
	err = app.DB.QueryRow("SELECT id, title, description, severity, owner, status, created_at, updated_at FROM incidents WHERE id = ?", id).
		Scan(&inc.ID, &inc.Title, &inc.Description, &inc.Severity, &inc.Owner, &inc.Status, &inc.CreatedAt, &inc.UpdatedAt)
	if err != nil {
		http.Error(w, "Incidente não encontrado", http.StatusNotFound)
		return
	}

	var timeline []TimelineItem

	// 1. Buscar Historico de Status
	rowsH, _ := app.DB.Query("SELECT previous_status, new_status, changed_at FROM incident_history WHERE incident_id = ?", id)
	defer rowsH.Close()
	for rowsH.Next() {
		var item TimelineItem
		item.Type = "status_change"
		if err := rowsH.Scan(&item.OldStatus, &item.NewStatus, &item.Timestamp); err == nil {
			timeline = append(timeline, item)
		}
	}

	// 2. Buscar Comentarios
	rowsC, _ := app.DB.Query("SELECT author, content, created_at FROM incident_comments WHERE incident_id = ?", id)
	defer rowsC.Close()
	for rowsC.Next() {
		var item TimelineItem
		item.Type = "comment"
		if err := rowsC.Scan(&item.Author, &item.Content, &item.Timestamp); err == nil {
			timeline = append(timeline, item)
		}
	}

	// 3. Ordenar Cronologicamente por Data/Hora
	sort.Slice(timeline, func(i, j int) bool {
		return timeline[i].Timestamp.Before(timeline[j].Timestamp)
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"incident": inc,
		"timeline": timeline,
	})
}
