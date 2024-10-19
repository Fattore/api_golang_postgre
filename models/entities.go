package models

//vai utilizar para carregar os dados para dentro do banco e quando ler poder carregar para resposta da API
type Todo struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        bool   `json:"done"`
}
