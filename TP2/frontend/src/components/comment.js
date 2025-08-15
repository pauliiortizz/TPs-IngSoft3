import React, { useState } from 'react';
import { Dialog, DialogTitle, DialogContent, DialogActions, Button, TextField } from '@mui/material';

const CommentModal = ({ isOpen, onClose, onSubmit }) => {
    const [comment, setComment] = useState('');

    // Manejar el cambio en el campo de texto del comentario
    const handleCommentChange = (e) => {
        setComment(e.target.value);  // Actualizar el estado con el nuevo valor del comentario
    };

    // Manejar la acción de enviar el comentario
    const handleSubmit = () => {
        onSubmit(comment);  // Llamar a la función onSubmit con el comentario actual
        setComment(''); // Reset comment field after submission
    };

    return (
        // Componente de diálogo de Material-UI que se muestra u oculta según el valor de isOpen
        <Dialog open={isOpen} onClose={onClose}>
            <DialogTitle>Ingresar Comentario</DialogTitle>
            <DialogContent>
                <TextField
                    label="Comentario" // Etiqueta del campo de texto
                    value={comment} // Valor actual del comentario
                    onChange={handleCommentChange} // Función que maneja el cambio en el campo de texto
                    fullWidth // Ocupa todo el ancho disponible
                    multiline // Permite múltiples líneas de texto
                    rows={4} // Número de filas visibles en el campo de texto
                />
            </DialogContent>
            <DialogActions>
                <Button onClick={onClose} color="secondary">Cancelar</Button>{/* Botón para cerrar el diálogo sin enviar */}
                <Button onClick={handleSubmit} color="primary">Enviar</Button>{/* Botón para enviar el comentario */}
            </DialogActions>
        </Dialog>
    );
};

export default CommentModal;
