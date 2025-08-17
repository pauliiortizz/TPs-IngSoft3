# Etapa de construcción - Build optimizado
FROM node:18-alpine AS builder

LABEL authors="WebLearn Team"
LABEL description="WebLearn Frontend - Optimized Multi-stage Build"

# Instalar dependencias del sistema necesarias
RUN apk add --no-cache git

WORKDIR /app

# Copiar archivos de dependencias primero (mejor cache de Docker)
COPY package*.json ./
RUN npm ci --only=production --silent

# Copiar código fuente
COPY . .

# Build optimizado para producción
RUN npm run build

# Etapa de producción - Nginx optimizado
FROM nginx:1.25-alpine

# Instalar curl para health checks
RUN apk add --no-cache curl

# Crear usuario no-root para nginx
RUN addgroup -g 1001 -S nginx-app && \
    adduser -S -D -H -u 1001 -h /var/cache/nginx -s /sbin/nologin -G nginx-app -g nginx-app nginx-app

# Copiar archivos construidos
COPY --from=builder /app/build /usr/share/nginx/html

# Copiar configuración personalizada de nginx
COPY nginx.conf /etc/nginx/nginx.conf
COPY default.conf /etc/nginx/conf.d/default.conf

# Configurar permisos
RUN chown -R nginx-app:nginx-app /usr/share/nginx/html && \
    chown -R nginx-app:nginx-app /var/cache/nginx && \
    chown -R nginx-app:nginx-app /var/log/nginx && \
    chown -R nginx-app:nginx-app /etc/nginx/conf.d

# Crear directorio para PID de nginx
RUN mkdir -p /var/run/nginx && \
    chown -R nginx-app:nginx-app /var/run/nginx

# Variables de entorno
ENV NGINX_HOST=localhost
ENV NGINX_PORT=80
ENV API_URL=http://backend:8080

# Usar usuario no-root
USER nginx-app

EXPOSE 80

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:80/ || exit 1

# Comando de inicio
CMD ["nginx", "-g", "daemon off;"]
