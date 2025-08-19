# WebLearn - Aplicación de Cursos Containerizada

Aplicación web de gestión de cursos desarrollada con Go (backend) y JavaScript/React (frontend), completamente containerizada con Docker.

## Arquitectura

- **Backend**: API REST en Go con Gin framework
- **Frontend**: Aplicación web con React servida por Nginx
- **Base de datos**: MySQL 8.0
- **Containerización**: Docker con multi-stage builds
- **Orquestación**: Docker Compose para entornos QA y PROD

## Prerrequisitos

- Docker Engine 20.10+
- Docker Compose 2.0+
- Git
- Cuenta en Docker Hub (opcional, para pull de imágenes)

## Instrucciones de Uso

### 1. Clonar el Repositorio

\`\`\`bash
git clone <tu-repositorio>
cd TP2
\`\`\`

### 2. Construir las Imágenes (Opcional)

Las imágenes ya están disponibles en Docker Hub, pero puedes construirlas localmente:

\`\`\`bash
# Backend
docker build -t weblearn-backend ./backend

# Frontend  
docker build -t weblearn-frontend ./frontend
\`\`\`

### 3. Ejecutar los Contenedores

#### Opción A: Entorno Unificado (QA + PROD)
\`\`\`bash
# Levantar ambos entornos simultáneamente
docker-compose up -d

# Verificar estado
docker-compose ps
\`\`\`

#### Opción B: Entornos Separados
\`\`\`bash
# Solo QA
docker-compose -f docker-compose.qa.yml up -d

# Solo PROD
docker-compose -f docker-compose.prod.yml up -d
\`\`\`

### 4. Acceder a la Aplicación

#### URLs de Acceso

**Entorno QA:**
- Frontend: http://localhost:8001
- Backend API: http://localhost:8081
- Base de datos: localhost:3308

**Entorno PROD:**
- Frontend: http://localhost:8002
- Backend API: http://localhost:8082
- Base de datos: localhost:3309

#### Puertos Utilizados

| Servicio | QA | PROD |
|----------|----|----- |
| Frontend | 8001 | 8002 |
| Backend | 8081 | 8082 |
| MySQL | 3308 | 3309 |

## Conectarse a la Base de Datos

### Credenciales por Defecto

**QA:**
\`\`\`
Host: localhost
Puerto: 3308
Usuario: root
Contraseña: qa_password
Base de datos: weblearn_qa
\`\`\`

**PROD:**
\`\`\`
Host: localhost
Puerto: 3309
Usuario: root
Contraseña: prod_password
Base de datos: weblearn_prod
\`\`\`

### Conexión desde Terminal

\`\`\`bash
# QA
mysql -h localhost -P 3308 -u root -pqa_password weblearn_qa

# PROD
mysql -h localhost -P 3309 -u root -pprod_password weblearn_prod
\`\`\`

### Conexión desde Aplicación

Las aplicaciones se conectan automáticamente usando las variables de entorno configuradas en docker-compose.yml.

## Verificar Funcionamiento

### 1. Estado de Contenedores
\`\`\`bash
docker-compose ps
\`\`\`
Todos los servicios deben mostrar estado "Up" y "healthy".

### 2. Logs de Aplicación
\`\`\`bash
# Ver logs de todos los servicios
docker-compose logs

# Logs específicos
docker-compose logs backend-qa
docker-compose logs frontend-prod
\`\`\`

### 3. Health Checks
\`\`\`bash
# Backend QA
curl http://localhost:8081/health

# Backend PROD
curl http://localhost:8082/health
\`\`\`

### 4. Verificar Base de Datos
\`\`\`bash
# Conectar y verificar tablas
docker exec -it weblearn-db-qa mysql -u root -pqa_password -e "USE weblearn_qa; SHOW TABLES;"
\`\`\`

### 5. Prueba de Persistencia
\`\`\`bash
# Reiniciar contenedores
docker-compose restart

# Verificar que los datos persisten
curl http://localhost:8081/cursos
\`\`\`

## 🔧 Comandos Útiles

### Gestión de Contenedores
\`\`\`bash
# Detener todos los servicios
docker-compose down

# Detener y eliminar volúmenes
docker-compose down -v

# Reconstruir imágenes
docker-compose build --no-cache

# Ver uso de recursos
docker stats
\`\`\`

### Gestión de Volúmenes
\`\`\`bash
# Listar volúmenes
docker volume ls | grep weblearn

# Inspeccionar volumen
docker volume inspect weblearn_mysql_data_qa

# Backup de base de datos
docker exec weblearn-db-qa mysqldump -u root -pqa_password weblearn_qa > backup_qa.sql
\`\`\`

### Limpieza del Sistema
\`\`\`bash
# Eliminar contenedores detenidos
docker container prune

# Eliminar imágenes no utilizadas
docker image prune -a

# Eliminar volúmenes no utilizados
docker volume prune
\`\`\`

## 🐳 Imágenes en Docker Hub

Las imágenes están disponibles públicamente:

- **Backend**: `delfisalinasmich/weblearn-backend:v1.0`
- **Frontend**: `delfisalinasmich/weblearn-frontend:v1.0`

### Tags Disponibles
- `v1.0`: Versión estable de producción
- `latest`: Última versión de desarrollo

## Variables de Entorno

### QA Environment
\`\`\`env
APP_ENV=qa
GIN_MODE=debug
DB_HOST=weblearn-db-qa
DB_PORT=3306
DB_USER=root
DB_PASSWORD=qa_password
DB_NAME=weblearn_qa
\`\`\`

### PROD Environment
\`\`\`env
APP_ENV=production
GIN_MODE=release
DB_HOST=weblearn-db-prod
DB_PORT=3306
DB_USER=root
DB_PASSWORD=prod_password
DB_NAME=weblearn_prod
\`\`\`

## Monitoreo y Logs

### Logs en Tiempo Real
\`\`\`bash
# Seguir logs de todos los servicios
docker-compose logs -f

# Logs específicos con timestamps
docker-compose logs -f --timestamps backend-qa
\`\`\`

### Métricas de Contenedores
\`\`\`bash
# Uso de recursos en tiempo real
docker stats

# Información detallada de un contenedor
docker inspect weblearn-backend-qa
\`\`\`

## Solución de Problemas

### Problemas Comunes

1. **Puerto ocupado**
   \`\`\`bash
   # Verificar qué proceso usa el puerto
   lsof -i :8001
   
   # Cambiar puerto en docker-compose.yml si es necesario
   \`\`\`

2. **Contenedor no inicia**
   \`\`\`bash
   # Ver logs detallados
   docker-compose logs nombre-servicio
   
   # Verificar configuración
   docker-compose config
   \`\`\`

3. **Base de datos no conecta**
   \`\`\`bash
   # Verificar que MySQL esté listo
   docker-compose logs db-qa
   
   # Probar conexión manual
   docker exec -it weblearn-db-qa mysql -u root -pqa_password
   \`\`\`

4. **Volúmenes no persisten**
   \`\`\`bash
   # Verificar volúmenes montados
   docker inspect weblearn-db-qa | grep Mounts -A 10
   \`\`\`

### Reinicio Completo
\`\`\`bash
# Detener todo y limpiar
docker-compose down -v
docker system prune -f

# Levantar desde cero
docker-compose up -d
\`\`\`

## 📁 Estructura del Proyecto

\`\`\`
TP2/
├── backend/
│   ├── Dockerfile
│   ├── main.go
│   └── ...
├── frontend/
│   ├── Dockerfile
│   ├── package.json
│   └── ...
├── docker-compose.yml
├── docker-compose.qa.yml
├── docker-compose.prod.yml
├── .env.qa
├── .env.prod
├── decisiones.md
└── README.md
\`\`\`

## Documentación Adicional

- `decisiones.md`: Justificaciones técnicas y arquitecturales
- Logs de aplicación: Disponibles via `docker-compose logs`
- Health checks: Configurados en todos los servicios

## Contribución

1. Fork del repositorio
2. Crear rama feature (`git checkout -b feature/nueva-funcionalidad`)
3. Commit cambios (`git commit -am 'Agregar nueva funcionalidad'`)
4. Push a la rama (`git push origin feature/nueva-funcionalidad`)
5. Crear Pull Request

---

**Curso**: Ingeniería de Software III  
**Universidad**: Universidad Católica de Córdoba
