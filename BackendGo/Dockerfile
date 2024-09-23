# Construção do backend em Go
FROM golang:1.18

# Definir diretório de trabalho
WORKDIR /app

# Copiar código fonte
COPY . .

# Baixar dependências
RUN go mod download

# Compilar o executável
RUN go build -o backendGo .

# Expor porta 8080
EXPOSE 8080

# Executar a aplicação
CMD ["./backendGo"]
