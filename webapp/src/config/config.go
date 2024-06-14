package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL para comunicação com a API
	APIURL = ""
	// Porta onde a aplicação web está rodando
	Porta = 0
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizada para criptografar os dados do cookie
	BlockKey []byte
)

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	erro := godotenv.Load()
    if erro != nil {
        log.Fatalf("Error loading .env file")
    }
	
	Porta, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Porta = 9000
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

}