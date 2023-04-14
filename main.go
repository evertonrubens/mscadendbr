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
	err = godotenv.Load(dir + "/prd.env")
	if err != nil {
		log.Print("Arquivo .env não encontrado.")
	}
}

func main() {

	// cria um novo roteador
	r := mux.NewRouter()

	// define as rotas
	routes.SetupRoutes(r)

	// inicia o servidor na porta definida nas variáveis de ambiente
	log.Println("Iniciando o servidor no IP: ", os.Getenv("IP_SERVER")+os.Getenv("PORT_SERVER"))
	log.Fatal(http.ListenAndServe(os.Getenv("PORT_SERVER"), r))
}
