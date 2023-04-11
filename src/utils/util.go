package tools

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

func GenerateID() string {
	// Cria uma string aleatória com 32 bytes de comprimento
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)

	// Gera o hash com base na string aleatória e converte para string
	return fmt.Sprintf("%x", sha256.Sum256(randomBytes))
}
