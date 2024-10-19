package handlers

import (
	"API/models"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	todo, err := models.Get(int64(id))
	if err != nil {
		log.Printf("Erro ao obter registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "aplication/json")
	//vai fazer o encode do map de resposta da aplicação
	json.NewEncoder(w).Encode(todo)
}
