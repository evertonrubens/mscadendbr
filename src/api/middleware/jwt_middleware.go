package middleware

import (
	"fmt"
	"log"
	"os"
	"strings"	
	"net/http"
	"github.com/rs/cors"
	"github.com/dgrijalva/jwt-go"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		
		cors.Default().HandlerFunc(w, r)


		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Extract token from header
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Validate token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

/*
func JWTMiddleware() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		cors.Default().HandlerFunc(w, r)

		tokenString := r.Header.Get("Authorization")

		//tokenString := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", 1)

		if tokenString == "" {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token não encontrado"))
			log.Println("O Bearer Token não foi informado -> ", tokenString)
			return
		}
		log.Println("Analisando o token informado no Authorization:", tokenString)

		token, err := parseJWTToken(tokenString)

		if err != nil {
			log.Println("Erro dentro do JWTMiddleware, apos a execução do parseJWTToken -> ", err.Error())
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(err.Error()))
			return
		}

		if !token.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Token Inválido"))
			return
		}
	}
}
*/

func parseJWTToken(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, verifyToken)
	log.Println("valor do tokenstr: ", tokenStr)
	if err != nil {
		log.Println("Escrevendo um possível erro no momento de fazer o parseJWTToken: ", err)
		return nil, err
	}
	return token, nil
}

func verifyToken(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return []byte("secret"), nil
}
