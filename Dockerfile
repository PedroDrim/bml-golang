# Etapa de Build
FROM alpine:3.14 AS build

# Instala dependências necessárias para extração dos arquivos
RUN apk add --no-cache unzip

# Define diretório de trabalho
WORKDIR /build

# Copia os arquivos de dados para a etapa de build
COPY ./data /build/

# Descompacta os arquivos de simulação
RUN unzip simulationInput_D.zip -d .

FROM golang:1.24.2-alpine3.21

# Create app directory
WORKDIR /app

# Copy files to directory
COPY ./ /app/

# Copia os arquivos extraídos da etapa de build
COPY --from=build /build /app/data

# Executando testes
#RUN go test ./...

# Iniciando CLI
ENTRYPOINT ["sh","Bench.sh"]
