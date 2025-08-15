import React, { useEffect, useState } from 'react';
import { useCookies } from 'react-cookie';

export const MisCursos = () => {
    // Estado para almacenar los cursos suscritos y el estado de carga
    const [cursosSuscritos, setCursosSuscritos] = useState([]);
    const [loading, setLoading] = useState(true);
    const [cookies, setCookie, removeCookie] = useCookies(['token']);

    // useEffect se ejecuta al montar el componente, llamando a fetchCursosSuscritos
    useEffect(() => {
        fetchCursosSuscritos();
    }, []);


    const fetchCursosSuscritos = () => {
        // Obtener el token y el userId del almacenamiento local
        const token = cookies.token;
        const userId = localStorage.getItem('userId');

        // Verificar si el token y el userId están presentes
        if (!token || !userId) {
            console.error('Usuario no autenticado');
            return;
        }

        // Realizar la petición fetch para obtener los cursos suscritos del usuario
        fetch(`http://localhost:8080/usuarios/${userId}/cursos`, {
            headers: {
                'Authorization': `Bearer ${token}` // Enviar el token en los encabezados de la petición
            }
        })
            .then(response => response.json()) // Convertir la respuesta a JSON
            .then(data => {
                setCursosSuscritos(data); // Guardar los cursos en el estado
                setLoading(false);  // Actualizar el estado de carga a falso
            })
            .catch(error => {
                console.error('Error al cargar los cursos suscritos:', error);
                setLoading(false); // En caso de error, actualizar el estado de carga a falso
            });
    };

    // Renderizar un mensaje de carga mientras se obtienen los datos
    if (loading) {
        return <p>Todavía no ha iniciado sesión</p>;
    }

    // Renderizar la lista de cursos o un mensaje si no hay cursos
    return (
        <div className="mis-cursos">
            <h2>Mis Cursos</h2>
            {cursosSuscritos.length > 0 ? (
                <div className="cursos-list">
                    {cursosSuscritos.map(curso => (
                        <div key={curso.id_curso} className="curso-item">
                            <p>{curso.Titulo}</p>
                        </div>
                    ))}
                </div>
            ) : (
                <p>No estás suscrito a ningún curso aún.</p>
            )}
        </div>
    );
}

// Exportar el componente por defecto
export default MisCursos;
