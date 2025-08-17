# Decisiones Técnicas - WebLearn Containerization

## 1. Elección de la Aplicación

**Aplicación elegida:** WebLearn - Plataforma de cursos online

**Justificación:**
- Aplicación full-stack completa (Go backend + React frontend + MySQL)
- Arquitectura real de producción con múltiples servicios
- Permite demostrar containerización de diferentes tecnologías
- Ya tenía funcionalidad completa implementada del TP anterior

## 2. Construcción de Imágenes Personalizadas

### Backend (Go)
**Imagen base:** `golang:1.21-alpine` (build) + `alpine:latest` (runtime)

**Justificación:**
- **Multi-stage build** para optimizar tamaño final
- **Alpine Linux** por ser liviana y segura (5MB vs 100MB+ de Ubuntu)
- **Compilación estática** para mejor rendimiento
- **Usuario no-root** para seguridad
- **Health checks** para monitoreo automático

**Instrucciones clave:**
\`\`\`dockerfile
# Build stage optimizado
FROM golang:1.21-alpine AS builder
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Runtime stage mínimo
FROM alpine:latest
RUN adduser -D -s /bin/sh appuser
USER appuser
\`\`\`

### Frontend (React)
**Imagen base:** `node:18-alpine` (build) + `nginx:1.25-alpine` (runtime)

**Justificación:**
- **Multi-stage build** para separar build de runtime
- **Node.js 18** para mejor rendimiento y soporte ES modules
- **Nginx Alpine** como servidor web optimizado
- **Configuración personalizada** con headers de seguridad
- **Compresión gzip** para mejor rendimiento

## 3. Estrategia de Versionado

**Tags utilizados:**
- `v1.0` - Versión específica para releases
- `latest` - Última versión estable

**Justificación:**
- **Versionado semántico** para control de releases
- **Tags específicos** para reproducibilidad
- **Latest tag** para desarrollo y testing

**Publicación en Docker Hub:**
- **Repositorios:** `delfisalinasmich/weblearn-backend` y `delfisalinasmich/weblearn-frontend`
- **Proceso de release:**
  \`\`\`bash
  # Tag con versión específica
  docker tag weblearn-backend delfisalinasmich/weblearn-backend:v1.0
  docker tag weblearn-frontend delfisalinasmich/weblearn-frontend:v1.0
  
  # Tag latest para desarrollo
  docker tag weblearn-backend delfisalinasmich/weblearn-backend:latest
  docker tag weblearn-frontend delfisalinasmich/weblearn-frontend:latest
  
  # Push a Docker Hub
  docker push delfisalinasmich/weblearn-backend:v1.0
  docker push delfisalinasmich/weblearn-frontend:v1.0
  \`\`\`

**Estrategia futura:**
- **v1.x** para patches y bug fixes
- **v2.x** para cambios mayores con breaking changes
- **Tags de entorno** como `qa`, `staging`, `prod` para diferentes ambientes

## 4. Base de Datos

**Base de datos elegida:** MySQL 8.0

**Justificación:**
- **Compatibilidad** con la aplicación existente
- **Soporte ARM64** para desarrollo en Mac Apple Silicon
- **Volúmenes persistentes** para durabilidad de datos
- **Configuración optimizada** para contenedores

**Configuración:**
\`\`\`yaml
db:
  image: mysql:8.0
  volumes:
    - mysql_data:/var/lib/mysql
  environment:
    MYSQL_ROOT_PASSWORD: ${DB_PASSWORD}
    MYSQL_DATABASE: weblearn
\`\`\`

## 5. Optimizaciones Implementadas

### Seguridad
- Usuarios no-root en todos los contenedores
- Headers de seguridad en nginx
- Variables de entorno para credenciales
- Imágenes base actualizadas

### Rendimiento
- Multi-stage builds para imágenes más pequeñas
- Compilación estática en Go
- Compresión gzip en nginx
- Cache optimizado para assets estáticos

### Monitoreo
- Health checks en todos los servicios
- Logs estructurados
- Configuración para diferentes entornos

## 6. Arquitectura Final

\`\`\`
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   Frontend      │    │    Backend      │    │    Database     │
│   (nginx)       │◄──►│     (Go)        │◄──►│    (MySQL)      │
│   Port: 80      │    │   Port: 8080    │    │   Port: 3306    │
└─────────────────┘    └─────────────────┘    └─────────────────┘
\`\`\`

**Red:** `app-network` (bridge) para comunicación entre servicios
**Volúmenes:** `mysql_data` para persistencia de base de datos
