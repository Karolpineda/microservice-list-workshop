# Usa una imagen base oficial de Go
FROM golang:1.23.6-alpine as builder
 
WORKDIR /app
 
COPY go.mod go.sum ./
RUN go mod tidy
 
# Copia todo el código fuente al contenedor, incluyendo el archivo .env
COPY . .
 
# Compila el binario
RUN go build -o main .
 
# Imagen base ligera para la ejecución
FROM alpine:latest
 
# Instala dependencias necesarias como el cliente de PostgreSQL
RUN apk --no-cache add postgresql-client
 
WORKDIR /root/
 
# Copiar el binario compilado
COPY --from=builder /app/main .
 
COPY --from=builder /app/.env .
 
# Exponer el puerto
EXPOSE 8097
 
# Comando para ejecutar la aplicación
CMD ["./main"]
