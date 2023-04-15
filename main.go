// @title msCadEndBr
// @version 1.0
// @description Documentação da API do seu microservice
// @BasePath /api/v1

package main

import (
	"log"
	"net/http"
	"os"

	"msCadEndBr/src/api/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	// carrega variáveis de ambiente
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load(dir + "/dev.env")
	if err != nil {
		log.Print("Arquivo .env não encontrado.")
	}
}

func main() {

	// cria um novo roteador
	r := mux.NewRouter()

	// define as rotas
	routes.SetupRoutes(r)

	port := os.Getenv("PORT_SERVER")
	if port == "" {
		port = "5001"
	}
	// inicia o servidor na porta definida nas variáveis de ambiente
	log.Println("Iniciando o servidor no IP: ", os.Getenv("IP_SERVER")+":"+port)
	log.Println("Servidor iniciado com sucesso!")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
