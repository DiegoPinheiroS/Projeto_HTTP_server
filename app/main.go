package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Resposta struct {
	Nome    string `json:"nome"`
	Horario string `json:"horario"`
}

var (
	httpRequestsTotal = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Volume total de requisições no endpoint /projeto-korp",
		},
	)
)

func init() {
	prometheus.MustRegister(httpRequestsTotal)
}

func projetoKorpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	httpRequestsTotal.Inc()

	horaUTC := time.Now().UTC().Format(time.RFC3339)

	response := Resposta{
		Nome:    "Projeto Korp",
		Horario: horaUTC,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/projeto-korp", projetoKorpHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Iniciando o servidor...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Erro crítico ao iniciar o servidor: %v", err)
	}
}
