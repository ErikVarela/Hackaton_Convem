package main

import (
	"testing"
)

func TestValidateStatusTransition(t *testing.T) {
	tests := []struct {
		name          string
		severity      string
		currentStatus string
		newStatus     string
		expectError   bool
	}{
		{
			name:          "Critical: Open -> Resolved (DEVE FALHAR)",
			severity:      SeverityCritical,
			currentStatus: StatusOpen,
			newStatus:     StatusResolved,
			expectError:   true,
		},
		{
			name:          "Critical: Open -> In Progress (PERMITIDO)",
			severity:      SeverityCritical,
			currentStatus: StatusOpen,
			newStatus:     StatusInProgress,
			expectError:   false,
		},
		{
			name:          "Critical: In Progress -> Resolved (PERMITIDO)",
			severity:      SeverityCritical,
			currentStatus: StatusInProgress,
			newStatus:     StatusResolved,
			expectError:   false,
		},
		{
			name:          "High: Open -> Resolved (PERMITIDO PARA NAO CRITICAL)",
			severity:      SeverityHigh,
			currentStatus: StatusOpen,
			newStatus:     StatusResolved,
			expectError:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateStatusTransition(tt.severity, tt.currentStatus, tt.newStatus)
			if tt.expectError && err == nil {
				t.Errorf("esperava erro mas obteve nil para o caso %s", tt.name)
			}
			if !tt.expectError && err != nil {
				t.Errorf("nao esperava erro mas obteve %v no caso %s", err, tt.name)
			}
		})
	}
}

func TestValidateComment(t *testing.T) {
	tests := []struct {
		name        string
		author      string
		content     string
		expectError bool
	}{
		{"Comentario Valido", "Ana", "Provedor contatado.", false},
		{"Autor Vazio", "", "Provedor contatado.", true},
		{"Conteudo Vazio", "Ana", "   ", true},
		{"Ambos Vazios", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateComment(tt.author, tt.content)
			if tt.expectError && err == nil {
				t.Errorf("esperava erro mas obteve nil no caso %s", tt.name)
			}
			if !tt.expectError && err != nil {
				t.Errorf("nao esperava erro mas obteve %v no caso %s", err, tt.name)
			}
		})
	}
}
