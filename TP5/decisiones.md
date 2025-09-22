# Decisiones Técnicas - TP05 - Salinas Delfina, Ortiz Paulina

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

<img width="788" height="238" alt="image" src="https://github.com/user-attachments/assets/d0e38132-241b-4687-a595-06b8ca6fdf18" />

---

## 3. Variables y Secrets

- Se configuraron variables de entorno para distinguir entornos:
  - `.env.qa` → `REACT_APP_API_URL=https://tp05-backend-qa-chdtg5exgzarc7hd.brazilsouth-01.azurewebsites.net/`
  - `.env.prod` → `REACT_APP_API_URL=https://tp05-backend-prod-cudvdwd9c0exdbhc.brazilsouth-01.azurewebsites.net/`
- Se almacenaron **deployment tokens** de Static Web Apps como secretos en Azure DevOps:
  - `$(SWA_TOKEN_QA)`
  - `$(SWA_TOKEN_PROD)`

<img width="1068" height="243" alt="image" src="https://github.com/user-attachments/assets/1d38c256-8961-4b22-96d8-3e2e72a170fd" />

<img width="947" height="471" alt="image" src="https://github.com/user-attachments/assets/125504d0-0332-4df8-a277-9467907dfe3e" />


---

## 4. Gestión de aprobaciones

- En el environment **PROD** de Azure DevOps se configuró **aprobación manual** antes del despliegue.
- Responsables:
  - QA → despliegue automático tras la build.
  - PROD → requiere confirmación manual del responsable del equipo.
    
<img width="1070" height="293" alt="image" src="https://github.com/user-attachments/assets/6bce2f62-267e-4a0d-9aa4-26100e2f1d95" />

---

## 5. Health Checks

- Se validaron los endpoints expuestos por los backends:
  - QA → `https://tp05-backend-qa-chdtg5exgzarc7hd.brazilsouth-01.azurewebsites.net/users`

<img width="1311" height="173" alt="image" src="https://github.com/user-attachments/assets/24108b5d-11c3-435a-8444-e5e7e1123060" />

  - PROD → `https://tp05-backend-prod-cudvdwd9c0exdbhc.brazilsouth-01.azurewebsites.net/users`

<img width="1292" height="185" alt="image" src="https://github.com/user-attachments/assets/d6448572-4a75-4067-9e20-5fcd7fb184f6" />

- Ambos devuelven datos JSON correctamente.
- Se planificó validar el frontend accediendo a la URL de Static Web Apps, confirmando integración con la API.
  - QA → `https://polite-sky-043673010.1.azurestaticapps.net/`
    <img width="482" height="168" alt="image" src="https://github.com/user-attachments/assets/5cd3817d-b28c-48e8-8d53-1ede35370083" />

  - PROD → `https://ambitious-cliff-063ec7210.1.azurestaticapps.net/`
    <img width="504" height="174" alt="image" src="https://github.com/user-attachments/assets/c3f27fc5-9cd1-4441-9fc1-52dba7dcea65" />


---


## 6. Decisiones importantes

- Se eligió **Azure** como proveedor cloud por simplicidad y conexión nativa con Azure DevOps.
- Separación clara de ambientes QA y PROD para:
  - Minimizar riesgos en despliegues productivos.
  - Permitir validación previa en QA.
- Uso de **Aprobación Manual** como gate de seguridad en el pase a Producción.
- Uso de **artefactos** para garantizar consistencia entre build y release.

---
