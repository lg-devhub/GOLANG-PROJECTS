package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID    int    `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var clientes = []Cliente{}

func listarClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(clientes)
}

func criarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var cliente Cliente

	err := json.NewDecoder(r.Body).Decode(&cliente)

	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	cliente.ID = len(clientes) + 1

	clientes = append(clientes, cliente)

	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(cliente)
}

func buscarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	for _, cliente := range clientes {
		if cliente.ID == id {
			json.NewEncoder(w).Encode(cliente)
			return
		}
	}

	http.Error(w, "Cliente não encontrado", http.StatusNotFound)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/clientes", listarClientes).Methods("GET")
	router.HandleFunc("/clientes", criarCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", buscarCliente).Methods("GET")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Erro ao iniciar servidor:", err)
	}
}