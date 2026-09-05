package main

import (
	"errors"
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

var ErrInvalidStatusTransition = errors.New("transição de status inválida: incidentes Critical não podem passar direto de Open para Resolved sem passar por In Progress")

// ValidateStatusTransition verifica as regras de negocio de transição de status
func ValidateStatusTransition(severity, currentStatus, newStatus string) error {
	// Regra de Negocio do Desafio:
	// Um incidente Critical nao pode passar diretamente de Open para Resolved.
	if severity == SeverityCritical && currentStatus == StatusOpen && newStatus == StatusResolved {
		return ErrInvalidStatusTransition
	}
	return nil
}
