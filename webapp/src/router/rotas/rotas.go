package rotas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	fmt.Println(`Configurando as rotasssss...`,rotas)

	for _, rota := range rotas {
		fmt.Println(rota)
		router.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}
	fileServer := http.FileServer(http.Dir("./assets"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}