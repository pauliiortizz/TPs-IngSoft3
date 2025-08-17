# Etapa de construcción
FROM golang:1.21-alpine AS builder

LABEL authors="WebLearn Team"
LABEL description="WebLearn Backend - Optimized Multi-stage Build"

# Instalar dependencias necesarias para el build
RUN apk add --no-cache git ca-certificates tzdata

# Crear usuario no-root para la aplicación
RUN adduser -D -g '' appuser

WORKDIR /build

# Copiar archivos de dependencias primero (mejor cache de Docker)
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copiar código fuente
COPY . .

# Build optimizado para producción
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags='-w -s -extldflags "-static"' \
    -a -installsuffix cgo \
    -o main .

FROM alpine:latest

# Instalar dependencias mínimas necesarias
RUN apk add --no-cache ca-certificates tzdata bash curl

# Crear usuario no-root para la aplicación
RUN adduser -D -g '' appuser

WORKDIR /app

# Copiar binario compilado
COPY --from=builder /build/main ./main

# Copiar script wait-for-it
COPY --from=builder /build/wait-for-it.sh ./wait-for-it.sh
RUN chmod +x ./wait-for-it.sh

RUN chown -R appuser:appuser /app

# Usar usuario no-root
USER appuser

# Variables de entorno por defecto
ENV GIN_MODE=release
ENV PORT=8080
ENV DB_HOST=db
ENV DB_PORT=3306
ENV DB_NAME=weblearn
ENV DB_USER=root

EXPOSE 8080

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:8080/health || exit 1

CMD ["sh", "-c", "./wait-for-it.sh $DB_HOST:$DB_PORT --timeout=60 --strict -- ./main"]
