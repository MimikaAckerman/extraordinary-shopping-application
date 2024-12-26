package Handler

import (
	"bk-compras-extraordinarias/config"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Función para actualizar el estado de compra y notificar al solicitante
func ActualizarEstadoCompra(c *gin.Context) {
	// Obtener el ID del registro y el nuevo estado
	id := c.Param("id")
	var nuevoEstado struct {
		EstadoCompra string `json:"estado_compra" binding:"required"`
	}
	if err := c.ShouldBindJSON(&nuevoEstado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Verificar si el registro existe y obtener datos del solicitante
	querySelect := `
		SELECT usuario, estado_compra
		FROM compras_extraordinarias
		WHERE id = @id
	`
	var usuario string
	var estadoActual string

	err := config.DB.QueryRow(querySelect, sql.Named("id", id)).Scan(&usuario, &estadoActual)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Registro no encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar el registro: " + err.Error()})
		}
		return
	}

	// Si el estado no cambia, no hacer nada
	if estadoActual == nuevoEstado.EstadoCompra {
		c.JSON(http.StatusOK, gin.H{"message": "El estado ya está actualizado"})
		return
	}

	// Actualizar el estado en la base de datos
	queryUpdate := `
		UPDATE compras_extraordinarias
		SET estado_compra = @nuevoEstado
		WHERE id = @id
	`
	_, err = config.DB.Exec(queryUpdate, sql.Named("nuevoEstado", nuevoEstado.EstadoCompra), sql.Named("id", id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el estado: " + err.Error()})
		return
	}

	// Enviar correo al solicitante
	err = EnviarCorreoCambioEstado(usuario, nuevoEstado.EstadoCompra)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el correo: " + err.Error()})
		return
	}

	// Respuesta de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Estado actualizado y notificación enviada al solicitante"})
}

// Función para enviar un correo al solicitante
// Función para enviar un correo al solicitante
func EnviarCorreoCambioEstado(correoSolicitante, nuevoEstado string) error {
	body := fmt.Sprintf(`
		<div style="font-family: Arial, sans-serif; padding: 20px; border: 1px solid #dddddd;">
			<img src="https://redaccion.camarazaragoza.com/wp-content/uploads/2021/04/grupo-ub.png" alt="Grupo UB" style="width: 150px; margin-bottom: 20px;">
			<h2>Actualización del Estado de su Compra</h2>
			<p><strong>Nuevo Estado:</strong> %s</p>
			<p>Le informamos que el estado de su compra ha cambiado recientemente. Por favor, contacte con su aprobador en caso de requerir más información.</p>
		</div>
	`, nuevoEstado)

	// Asegúrate de llamar a EnviarCorreo correctamente
	return EnviarCorreo(correoSolicitante, "Actualización de Estado de Compra", body)
}
