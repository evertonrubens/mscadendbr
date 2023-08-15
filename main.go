// @title msCadEndBr
// @version 1.0
// @description Documentação da API do seu microservice
// @BasePath /api/v1

package main

import (
	"log"
	"net/http"
	"os"

	"cadEndBr/src/api/routes"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	serverName := os.Getenv("SERVER_NAME")
	if serverName == "DEV" {
		log.Println("Iniciando o servidor de DEV no IP: ", os.Getenv("SERVER_IP")+":"+port)
		log.Println("Servidor iniciado com sucesso!")
		log.Fatal(http.ListenAndServe(":"+port, r))
	}

	/*if serverName == "PRODUCAO" {
		log.Println("Iniciando o servidor de PRODUÇÃO")
	  log.Println("Servidor iniciado com sucesso!")
	  log.Fatal(http.ListenAndServeTLS(":"+port, r))
	}*/

}
