# Entorno Reproducible WebLearn

## 🎯 Objetivo
Este documento explica cómo asegurar que el entorno WebLearn se ejecute de manera idéntica en cualquier máquina.

## 📋 Prerrequisitos
- Docker Engine 20.10+
- Docker Compose 2.0+
- 4GB RAM disponible
- Puertos libres: 8001, 8002, 8081, 8082, 3308, 3309

## 🚀 Configuración Inicial
\`\`\`bash
# 1. Clonar el repositorio
git clone <tu-repositorio>
cd weblearn

# 2. Ejecutar script de configuración
chmod +x setup-environment.sh
./setup-environment.sh

# 3. Levantar todos los servicios
docker-compose up -d
\`\`\`

## 🔧 Estructura del Entorno
- **QA**: Entorno de pruebas con logs detallados
- **PROD**: Entorno de producción optimizado
- **Bases de datos**: MySQL 8.0 con volúmenes persistentes
- **Redes**: Aislamiento completo entre entornos

## 📊 Verificación de Estado
\`\`\`bash
# Ver estado de todos los servicios
docker-compose ps

# Ver logs de un servicio específico
docker-compose logs backend-qa

# Verificar health checks
docker-compose exec backend-qa curl http://localhost:8080/health
\`\`\`

## 🔒 Garantías de Reproducibilidad
1. **Versiones fijas**: Todas las imágenes usan tags específicos
2. **Variables de entorno**: Configuración centralizada en archivos .env
3. **Volúmenes nombrados**: Persistencia de datos garantizada
4. **Health checks**: Verificación automática de servicios
5. **Límites de recursos**: Consumo controlado de memoria
6. **Redes aisladas**: Separación completa entre entornos
