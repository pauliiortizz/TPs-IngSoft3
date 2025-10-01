# Frameworks de testing elegidos

1. Backend (Node.js + Express):
- Jest: framework de testing rápido y ampliamente usado en proyectos JS.
- Supertest: permite testear endpoints HTTP simulando requests sin levantar el servidor real.
- SQLite in-memory: usado para los tests en lugar de MySQL, lo que permite ejecutar pruebas aisladas sin depender de una base de datos externa.

2. Frontend (React):
- Jest: framework integrado por defecto con Create React App.
- React Testing Library (@testing-library/react): para testear componentes React a nivel de usuario (interacciones y contenido renderizado).
- @testing-library/jest-dom: agrega matchers personalizados como toBeInTheDocument.
- Axios (mockeado con Jest): se utiliza en los servicios para llamadas HTTP, pero en los tests se reemplaza por un mock para aislar dependencias externas.

# Estrategia de mocking

1. Backend:
- La conexión real a MySQL fue reemplazada por SQLite en memoria, usando sequelize.sync({ force: true }) para crear las tablas al inicio de cada suite.
- Esto asegura independencia de datos, tests reproducibles y sin riesgos para la DB real.

2. Frontend:
- Se creó un mock de Axios con Jest (jest.mock('axios')) para simular respuestas de la API.
- Esto evita depender del backend durante las pruebas de frontend y permite verificar que los métodos (axios.get) son llamados correctamente.

# Casos de prueba implementados
1. Backend
- GET /ping → responde con { message: "pong" }.
- GET /users → devuelve un arreglo (vacío al inicio, ya que se usa SQLite in-memory).
- POST /users → crea un usuario válido y devuelve status 201.
- POST /users sin email → devuelve error de validación.
- PUT /users/:id → actualiza un usuario existente.
- PUT /users/:id con ID inexistente → responde 404.
- DELETE /users/:id → elimina un usuario existente.
- DELETE /users/:id con ID inexistente → responde 404.

![Casos_de_Prueba_Backend](image-2.png)

2. Frontend
- App.test.js → verifica que se renderice el título de la aplicación.
- userService.test.js → prueba que el servicio getAllUsers devuelva usuarios mockeados y llame a axios.get('/users').
- UserList.test.js → renderiza la lista de usuarios usando el servicio mockeado.
- Validación de casos edge en frontend: renderizado correcto cuando el arreglo está vacío.

![Casos_de_Prueba_Frontend](image-3.png)


# Integración con CI/CD

- Se configuró el pipeline en Azure DevOps para ejecutar npm test en frontend y backend dentro de la stage Build and Test.

- Solo si los tests pasan, el pipeline continúa con Deploy QA y luego con Deploy PROD.

- Esto asegura que únicamente versiones validadas lleguen a entornos finales.

![alt text](image-7.png)


![alt text](image-8.png)