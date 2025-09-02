# decisiones.md - TP 3

# 1. Configuración inicial del proyecto
## Decisión
Usamos **Agile** ya que con las otras opciones no aparecian las User Stories. 

# 2. Gestión del trabajo con Azure Boards
Hicimos 3 equipos, ademas del principal. 
Creamos un sprint para el equipo 1 (de dos semanas) y le asignamos las US con sus tareas designadas.

- Agile (Epic → User Story → Task) y sprints.

# 3. Control de versiones con Azure Repos

## 🔹 Rama principal
- **main**  
  - Rama protegida.  
  - Configurada con políticas obligatorias:  
    - Requiere Pull Request para mergear cambios.  
    - Requiere al menos **1 reviewer**.  
    - No se permite push directo a `main`.  

---

## 🔹 Ramas de feature
Creamos ramas de funcionalidad siguiendo la convención `feature/<nombre>`.    
- `feature/registro-email` → Implementación del registro de usuario.  
- `feature/login-jwt` → Implementación del inicio de sesión con JWT.  

---

## 🔹 Flujo de trabajo (Workflow)
1. Se parte siempre desde la rama `main`.  
2. Se crea una rama de feature:  
   ```bash
   git checkout -b feature/registro-email
   git push origin feature/registro-email

Hicimos dos cambios para probar los PR
- git checkout feature/registro-email

# Creamos un archivo de prueba
- echo "print('Funcionalidad de registro')" > registro.py

# Guardamos y subimos
- git add registro.py
- git commit -m "Agrego script de prueba para registro"
- git push origin feature/registro-email


