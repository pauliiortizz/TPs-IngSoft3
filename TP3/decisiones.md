# decisiones.md - TP 3

# 1. Configuración inicial del proyecto
## Decisión
Usar **Scrum** con sprints de 2 semanas y seguimiento en Azure Boards.

## Contexto
- Requisitos cambiantes y necesidad de feedback frecuente del usuario.
- Equipo pequeño/mediano con entregas incrementales.
- Azure Boards soporta artefactos de Scrum (Epic → Product Backlog Item → Task) y sprints. 

## Alternativas consideradas
- **Agile (Agile process)**: flexible, pero Scrum aporta más cadencia y ceremonias.
- **Basic**: muy simple para equipos sin experiencia ágil; se queda corto para épicas/planificación.
- **CMMI**: más pesado, orientado a proyectos con fuerte gobierno.

## Consecuencias
- Timeboxing de 2 semanas.
- Métricas: Velocidad, Burn-down, % de historias Done.
- Board con columnas To Do / Doing / Code Review / Testing / Done.
- Estimación por **Story Points**; descomposición a **Tasks** (horas).

## Estado
Aceptada (fecha: AAAA-MM-DD).

# 2. Gestión del trabajo con Azure Boards

## Jerarquía de trabajo
Epic → User Story (PBI) → Task
Bug se gestiona en paralelo y se liga a la US afectada cuando corresponda.

## Definiciones
- **Epic**: funcionalidad de alto nivel (2–6 sprints).
- **User Story**: valor de negocio negociable, estimada en **Story Points**.
- **Task**: trabajo técnico, estimado en **horas**.
- **Bug**: defecto reproducible con pasos, resultado esperado y actual.

## Estados del flujo
To Do → Doing → Code Review → Testing → Done  
(Ajuste de columnas en Boards por equipo.)

## Sprints
Duración: 2 semanas. Fechas explícitas por Iteration Path.  
Capacidad: sumatoria de horas en Tasks; objetivo: completar US comprometidas.

## Criterios (Definition of Done)
- Código mergeado a `main` vía PR aprobado.
- Pipeline verde (build + test).
- Aceptación funcional validada.
- Documentación mínima (README/Changelog).

# 3. Control de versiones con Azure Repos

## Reglas
- `main` siempre *deployable*; protegido con Branch Policies (PR + 1 reviewer + build).
- Cada cambio nace en `feature/<tema>` desde `main`.
- Trabajo se integra vía **Pull Request** con:
  - Reviewer(s) mínimo 1
  - Pipeline de PR en verde (Build validation)
  - Work items vinculados

## Lanzamientos y hotfix
- Se etiqueta release en `main` (tag `vX.Y.Z`).
- Hotfix: `hotfix/<incidente>` desde `main`, PR a `main` con prioridad.

## Por qué no GitFlow clásico
- GitFlow añade ramas `develop`/`release` que incrementan complejidad.
- Nuestro equipo busca ciclos cortos (2 semanas) y *continuous delivery*.


