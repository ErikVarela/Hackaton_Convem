package main

import (
	"errors"
	"strings"
	"time"
)

// Constantes de Severidade
const (
	SeverityLow      = "Low"
	SeverityMedium   = "Medium"
	SeverityHigh     = "High"
	SeverityCritical = "Critical"
)

// Constantes de Status
const (
	StatusOpen       = "Open"
	StatusInProgress = "In Progress"
	StatusResolved   = "Resolved"
)

type Incident struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Severity    string    `json:"severity"`
	Owner       string    `json:"owner"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type IncidentHistory struct {
	ID             int       `json:"id"`
	IncidentID     int       `json:"incident_id"`
	PreviousStatus string    `json:"previous_status"`
	NewStatus      string    `json:"new_status"`
	ChangedAt      time.Time `json:"changed_at"`
}

type IncidentComment struct {
	ID         int       `json:"id"`
	IncidentID int       `json:"incident_id"`
	Author     string    `json:"author"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"created_at"`
}

// TimelineItem unifica historico e comentarios em uma unica linha do tempo (Change Request #1)
type TimelineItem struct {
	Type      string    `json:"type"` // "status_change" ou "comment"
	Author    string    `json:"author,omitempty"`
	Content   string    `json:"content,omitempty"`
	OldStatus string    `json:"old_status,omitempty"`
	NewStatus string    `json:"new_status,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

var ErrInvalidStatusTransition = errors.New("transição de status inválida: incidentes Critical não podem passar direto de Open para Resolved sem passar por In Progress")
var ErrInvalidComment = errors.New("autor e conteúdo do comentário são obrigatórios e não podem ser vazios")

func ValidateStatusTransition(severity, currentStatus, newStatus string) error {
	if severity == SeverityCritical && currentStatus == StatusOpen && newStatus == StatusResolved {
		return ErrInvalidStatusTransition
	}
	return nil
}

func ValidateComment(author, content string) error {
	if strings.TrimSpace(author) == "" || strings.TrimSpace(content) == "" {
		return ErrInvalidComment
	}
	return nil
}
