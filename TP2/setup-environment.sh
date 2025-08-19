#!/bin/bash

echo "🚀 Configurando entorno reproducible WebLearn..."

# Verificar que Docker esté instalado
if ! command -v docker &> /dev/null; then
    echo "❌ Docker no está instalado. Por favor instala Docker primero."
    exit 1
fi

# Verificar que Docker Compose esté instalado
if ! command -v docker-compose &> /dev/null; then
    echo "❌ Docker Compose no está instalado. Por favor instala Docker Compose primero."
    exit 1
fi

# Crear archivos de entorno si no existen
if [ ! -f .env.qa ]; then
    echo "⚠️  Archivo .env.qa no encontrado. Creando archivo por defecto..."
    cat > .env.qa << EOF
# QA Environment Variables
ENVIRONMENT=qa
DB_HOST=db-qa
DB_PORT=3306
DB_NAME=weblearn_qa
DB_USER=weblearn_user
DB_PASSWORD=weblearn_pass_qa
MYSQL_ROOT_PASSWORD=root_pass_qa
MYSQL_DATABASE=weblearn_qa
MYSQL_USER=weblearn_user
MYSQL_PASSWORD=weblearn_pass_qa
LOG_LEVEL=debug
API_URL=http://localhost:8081
EOF
fi

if [ ! -f .env.prod ]; then
    echo "⚠️  Archivo .env.prod no encontrado. Creando archivo por defecto..."
    cat > .env.prod << EOF
# PROD Environment Variables
ENVIRONMENT=production
DB_HOST=db-prod
DB_PORT=3306
DB_NAME=weblearn_prod
DB_USER=weblearn_user
DB_PASSWORD=weblearn_pass_prod
MYSQL_ROOT_PASSWORD=root_pass_prod
MYSQL_DATABASE=weblearn_prod
MYSQL_USER=weblearn_user
MYSQL_PASSWORD=weblearn_pass_prod
LOG_LEVEL=info
API_URL=http://localhost:8082
EOF
fi

# Crear archivo init-db.sql si no existe
if [ ! -f init-db.sql ]; then
    echo "⚠️  Archivo init-db.sql no encontrado. Creando estructura básica..."
    cat > init-db.sql << EOF
-- Estructura básica de la base de datos WebLearn
CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role ENUM('student', 'admin') DEFAULT 'student',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS cursos (
    id INT AUTO_INCREMENT PRIMARY KEY,
    titulo VARCHAR(200) NOT NULL,
    descripcion TEXT,
    categoria VARCHAR(100),
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS inscripciones (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    curso_id INT,
    fecha_inscripcion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (curso_id) REFERENCES cursos(id)
);
EOF
fi

echo "✅ Entorno configurado correctamente."
echo "📋 URLs de acceso:"
echo "   QA Frontend:  http://localhost:8001"
echo "   QA Backend:   http://localhost:8081"
echo "   PROD Frontend: http://localhost:8002"
echo "   PROD Backend:  http://localhost:8082"
echo ""
echo "🚀 Para iniciar: docker-compose up -d"
echo "🛑 Para detener: docker-compose down"
