import React from 'react';
import { Link } from 'react-router-dom';

export const ComponenteColumna = ({ isAdmin }) => {
    return (
        <div className="componente-columna">
            <p><Link className="home-texto" to="/">Home</Link></p>
            <p><Link className="miscursos-texto" to="/miscursos">Mis Cursos</Link></p>
        </div>
    );
};
