package response

import (
	"encoding/json"
	"net/http"
)

// JSON serializa o payload como JSON e escreve o status code informado.
func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		http.Error(w, "erro ao serializar resposta", http.StatusInternalServerError)
	}
}

// Error escreve uma resposta de erro padronizada em JSON.
func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}
