#!/bin/bash

echo "🚀 Iniciando entornos QA y PROD..."

# Levantar entorno QA
echo "📋 Levantando entorno QA..."
docker-compose -f docker-compose.qa.yml up -d

# Levantar entorno PROD
echo "🏭 Levantando entorno PROD..."
docker-compose -f docker-compose.prod.yml up -d

echo "✅ Entornos iniciados exitosamente!"
echo ""
echo "🔗 URLs de acceso:"
echo "QA Frontend:  http://localhost:8001"
echo "QA Backend:   http://localhost:8081"
echo "PROD Frontend: http://localhost:8002"
echo "PROD Backend:  http://localhost:8082"
echo ""
echo "📊 Para ver el estado:"
echo "docker ps"
