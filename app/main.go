package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	horaUTC := time.Now().UTC().Format(time.RFC3339)

	var response Resposta
	response.Nome = "Projeto Korp"
	response.Horario = horaUTC

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)

	fmt.Println("Servidor rodando na porta 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
