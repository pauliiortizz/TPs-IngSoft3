#!/bin/bash

echo "🛑 Deteniendo entornos QA y PROD..."

# Detener entorno QA
echo "📋 Deteniendo entorno QA..."
docker-compose -f docker-compose.qa.yml down

# Detener entorno PROD
echo "🏭 Deteniendo entorno PROD..."
docker-compose -f docker-compose.prod.yml down

echo "✅ Entornos detenidos exitosamente!"
