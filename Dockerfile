# Imagem base oficial do Go
FROM golang:1.18.2

# Definir diretório de trabalho
WORKDIR /app

# Copiar go.mod e go.sum
COPY go.mod go.sum ./

# Baixar dependências
RUN go mod download

# Copiar o código-fonte do projeto para o contêiner
COPY . .

# Compilar o projeto
RUN go build -o main .

# Expor a porta que seu microserviço vai usar (altere conforme necessário)
EXPOSE 5001

# Comando para executar o binário compilado
CMD ["./main"]
