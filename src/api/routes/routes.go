package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"cadEndBr/src/api/controllers"
	"cadEndBr/src/api/middleware"
)

func SetupRoutes(router *mux.Router) {

	//Rotas Publicas
	publicRoutes := router.PathPrefix("/public/v1").Subrouter()
	publicRoutes.HandleFunc("/token", controllers.TokenHandler).Methods(http.MethodPost)

	//Rotas Privadas
	privateRoutes := router.PathPrefix("/api/v1").Subrouter()

	// Adicionando o middleware JWT para a rota privada
	privateRoutes.Use(middleware.JWTMiddleware)

	// Endpoints privados
	privateRoutes.HandleFunc("/enderecos", controllers.CreateEnderecoHandler).Methods(http.MethodPost)
	privateRoutes.HandleFunc("/enderecos", controllers.GetAllEnderecosHandler).Methods(http.MethodGet)
	privateRoutes.HandleFunc("/enderecos/id/{id}", controllers.GetEnderecoByIdHandler).Methods(http.MethodGet)
	privateRoutes.HandleFunc("/enderecos/cep/{cep}", controllers.GetEnderecosByCepHandler).Methods(http.MethodGet)
	privateRoutes.HandleFunc("/enderecos/nomePF/{nome}", controllers.GetEnderecosByNomePFHandler).Methods(http.MethodGet)

	// Habilitando CORS para todas as rotas
	router.Use(cors.Default().Handler)
}

func ConvertToMiddlewareFunc(handler http.HandlerFunc) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			handler(w, r)
			next.ServeHTTP(w, r)
		})
	}
}
