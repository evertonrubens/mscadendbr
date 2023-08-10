package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/rs/cors"
)

// Middleware, algoritmo responsável por interceptar as requisições entre as
// rotas as implementações validando o token para aquelas que exigirem esta validação.
func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cors.Default().HandlerFunc(w, r)

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		// Extrair o token do Header Autorization
		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		// Validação do Token
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

		// Continuar para o proximo handler
		next.ServeHTTP(w, r)
	})
}

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
