# Decisiones Técnicas - TP05

Este documento resume las decisiones tomadas, configuraciones realizadas y evidencias recopiladas durante el trabajo práctico de despliegue en la nube con Azure DevOps y Azure Portal.

---

## 1. Cloud Resources

- **Backend**
  - Se crearon **Azure Web Apps** separados para QA y Producción:
    - `[tp05-backend-qa.azurewebsites.net](https://tp05-backend-qa-chdtg5exgzarc7hd.brazilsouth-01.azurewebsites.net/)`
    - `[tp05-backend-prod.azurewebsites.net](https://tp05-backend-prod-cudvdwd9c0exdbhc.brazilsouth-01.azurewebsites.net/)`
  - Se configuró el pipeline para empaquetar el backend en `.zip` y desplegarlo en cada App Service.

<img width="1153" height="72" alt="image" src="https://github.com/user-attachments/assets/5f04f9fb-76fc-4925-a23a-4f0f301375bd" />

- **Frontend**
  - Se crearon **Azure Static Web Apps** para QA y Producción:
    - `ambitious-cliff-063ec7210.1.azurestaticapps.net` (QA)
    - `...` (Producción)
  - El despliegue se realiza desde la carpeta `build/` generada por React.

<img width="1156" height="73" alt="image" src="https://github.com/user-attachments/assets/f9e20bdc-3484-45e7-8fbc-a577f8ca8062" />

---

## 2. Release Pipeline

- Se utilizó **Azure DevOps Pipelines** para orquestar Build y Release.
- Se creó un **Build stage** que compila y genera artefactos para:
  - **Frontend (React)** → `drop-front`
  - **Backend (Node.js + Express)** → `drop-back`
- Se configuraron stages separados:
  - **Deploy_QA** → despliegue automático al entorno QA.
  - **Deploy_PROD** → despliegue manual al entorno PROD.

---

## 3. Variables y Secrets

- Se configuraron variables de entorno para distinguir entornos:
  - `.env.qa` → `REACT_APP_API_URL=https://tp05-backend-qa.azurewebsites.net`
  - `.env.prod` → `REACT_APP_API_URL=https://tp05-backend-prod.azurewebsites.net`
- Se almacenaron **deployment tokens** de Static Web Apps como secretos en Azure DevOps:
  - `$(SWA_TOKEN_QA)`
  - `$(SWA_TOKEN_PROD)`

---

## 4. Gestión de aprobaciones

- En el environment **PROD** de Azure DevOps se configuró **aprobación manual** antes del despliegue.
- Responsables:
  - QA → despliegue automático tras la build.
  - PROD → requiere confirmación manual del responsable del equipo.

---

## 5. Health Checks

- Se validaron los endpoints expuestos por los backends:
  - QA → `https://tp05-backend-qa.azurewebsites.net/users`
  - PROD → `https://tp05-backend-prod.azurewebsites.net/users`
- Ambos devuelven datos JSON correctamente.
- Se planificó validar el frontend accediendo a la URL de Static Web Apps, confirmando integración con la API.

---

## 6. Evidencias

- **Capturas de pantalla** (adjuntar en entrega final):
  - Azure Portal con recursos creados (Web Apps y Static Web Apps).
  - Configuración de variables de entorno.
  - Pipeline en ejecución mostrando build + deploy exitosos.
  - Health checks (endpoints QA/PROD funcionando).
- **Archivos relevantes**:
  - `azure-pipelines.yml`
  - `staticwebapp.config.json` (para fallback de rutas SPA)
  - `.env.qa` y `.env.prod`

---

## 7. Decisiones importantes

- Se eligió **Azure** como proveedor cloud por simplicidad y conexión nativa con Azure DevOps.
- Separación clara de ambientes QA y PROD para:
  - Minimizar riesgos en despliegues productivos.
  - Permitir validación previa en QA.
- Uso de **Aprobación Manual** como gate de seguridad en el pase a Producción.
- Uso de **artefactos** para garantizar consistencia entre build y release.

---
