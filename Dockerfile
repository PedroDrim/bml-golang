FROM golang:1.24.2-alpine3.21

# Create app directory
WORKDIR /app

# Copy files to directory
COPY ./ /app/

# Executando testes
RUN ./go test ./...

# Iniciando CLI
ENTRYPOINT ["sh","Bench.sh"]
