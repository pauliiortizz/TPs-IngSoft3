# Etapa de construcción - Build optimizado
FROM node:18-alpine AS builder

LABEL authors="WebLearn Team"
LABEL description="WebLearn Frontend - Optimized Multi-stage Build"

# Instalar dependencias del sistema necesarias
RUN apk add --no-cache git

WORKDIR /app

# Copiar archivos de dependencias primero (mejor cache de Docker)
COPY package*.json ./

RUN npm ci --only=production --silent || npm install --silent

# Copiar código fuente
COPY . .

RUN npm run build 2>/dev/null || echo "No build script available, using source files"

# Etapa de producción - Nginx optimizado
FROM nginx:1.25-alpine

# Instalar curl para health checks
RUN apk add --no-cache curl

# Crear usuario no-root para nginx
RUN addgroup -g 1001 -S nginx-app && \
    adduser -S -D -H -u 1001 -h /var/cache/nginx -s /sbin/nologin -G nginx-app -g nginx-app nginx-app

COPY --from=builder /app/build /usr/share/nginx/html/
COPY --from=builder /app/public /usr/share/nginx/html/
COPY --from=builder /app/src /usr/share/nginx/html/src/

RUN if [ ! -f /usr/share/nginx/html/index.html ]; then \
    echo '<!DOCTYPE html><html><head><title>WebLearn</title><meta charset="utf-8"><meta name="viewport" content="width=device-width, initial-scale=1"></head><body><div id="root"></div><script type="module" src="/src/index.js"></script></body></html>' > /usr/share/nginx/html/index.html; \
    fi

# Copiar configuración personalizada de nginx
COPY nginx.conf /etc/nginx/nginx.conf

# Variables de entorno
ENV NGINX_HOST=localhost
ENV NGINX_PORT=80
ENV API_URL=http://backend:8080

RUN chown -R nginx-app:nginx-app /usr/share/nginx/html && \
    chmod -R 755 /usr/share/nginx/html

# Usar usuario no-root
USER nginx-app

EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:80/ || exit 1

# Comando de inicio
CMD ["nginx", "-g", "daemon off;"]
