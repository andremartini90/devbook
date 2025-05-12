package router

import (
	"fmt"
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	fmt.Println("Configurando as rotas...")
	r := mux.NewRouter()

	return rotas.Configurar(r)
}