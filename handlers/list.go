package handlers

import (
	"API/models"
	"encoding/json"
	"log"
	"net/http"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("Erro ao obter registros: %v", err)
	}

	w.Header().Add("Content-Type", "aplication/json")
	//vai fazer o encode do map de resposta da aplicação
	json.NewEncoder(w).Encode(todos)
}
