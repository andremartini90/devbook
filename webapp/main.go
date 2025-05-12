package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	fmt.Println("Rodando, WebApp!")
	utils.CarregarTemplates()
	r := router.Gerar()

	fmt.Println("Escutando na porta 3000...",r)
	log.Fatal(http.ListenAndServe(":3000", r))
}