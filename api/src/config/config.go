package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// String de conexão com o banco de dados
	StringConexaoBanco = ""

	// Porta da API
	Porta = 0

	// Chave usada para assinar o token
	SecretKey []byte
)

// Carregar carrega as variáveis de ambiente
func Carregar() {
	var erro error

	// Carregar variáveis do arquivo .env
	if erro = godotenv.Load(); erro != nil {
		log.Fatal("Erro ao carregar .env:", erro)
	}

	// Definir a porta da API (padrão: 9000)
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	// Criar string de conexão com o banco MySQL
	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NOME"),
	)

	// Chave usada para assinar o token
	SecretKey = []byte(os.Getenv("SECRET_KEY"))
}
