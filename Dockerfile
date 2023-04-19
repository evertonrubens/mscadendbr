# Use uma imagem oficial do Golang como base
FROM golang:1.18.2

# Defina a pasta de trabalho dentro do container
WORKDIR /app

# Copie o arquivo go.mod e go.sum (se existir) para a pasta de trabalho
COPY go.mod go.sum ./

# Faça o download das dependências
RUN go mod download

# Copie o restante do código para a pasta de trabalho
COPY . .

# Construa o executável do seu aplicativo
RUN go build -o main .

# Exponha a porta em que seu aplicativo será executado
EXPOSE 5001

# Execute o aplicativo
CMD ["./main"]
