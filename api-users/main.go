package main

import(
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Usuario struct {
	ID		int		`json:"id"`
	Nome	string	`json:"nome"`
}

var usuarios = []Usuario{
	{
		ID:		1,
		Nome:	"Luiz",
	},
	{
		ID:		2,
		Nome:	"Gustavo",
	},
}


func listarUsuarios(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	json.NewEncoder(w).Encode(usuarios)
}

func listarUsuariosQuantidade(w http.ResponseWriter, r *http.Request) {

	w.Header().Set(
		"Content-Type",
		"application/json",
	)

	resposta := map[string]int{
		"total": len(usuarios),
	}

	json.NewEncoder(w).Encode(resposta)
}

func buscarUsuario(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(
		r.URL.Path,
		"/usuarios/",
	)
	for _, usuario := range usuarios {
		if strconv.Itoa(usuario.ID) == id {
			w.Header().Set(
				"Content-Type",
				"application/json",
			)
			json.NewEncoder(w).Encode(usuario)

			return
		}
	}

	http.Error(
		w,
		"Usuário Não Encontrado",
		http.StatusNotFound,
	)
}


func main() {
	http.HandleFunc(
		"/usuarios",
		listarUsuarios,
	)

	http.HandleFunc(
		"/usuarios/",
		buscarUsuario,
	)

	http.HandleFunc(
		"/usuarios/quantidade",
		listarUsuariosQuantidade,
	)

	fmt.Println("Servidor iniciado na porta: 8080")
	err := http.ListenAndServe(
		":8080",
		nil,
	)

	if err != nil {
		fmt.Println(err)
	}
}