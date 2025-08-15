import React from 'react';

const ArchivosSubidos = ({ files }) => {
    return (
        <div className="ArchivosSubidos">
            <h2>Archivos Subidos</h2>
            {files.map(file => (
                // Mapea cada archivo en la lista `files` a un elemento div
                <div key={file} className="File">
                    {/* Enlace al archivo subido, abre en una nueva pestaña */}
                    <a href={`http://localhost:8080/uploads/${file}`} target="_blank" rel="noopener noreferrer">{file}</a>
                </div>
            ))}
        </div>
    );
};

export default ArchivosSubidos;