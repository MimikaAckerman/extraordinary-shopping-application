package Handler

import (
	"bk-compras-extraordinarias/config"
	"bk-compras-extraordinarias/models" // Importa el modelo correcto
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetComprasExtraordinarias devuelve todos los registros de la tabla compras_extraordinarias
func GetComprasExtraordinarias(c *gin.Context) {
	// Consulta SQL para seleccionar todos los registros, incluyendo datos_adjuntos
	query := `
		SELECT id, titulo, aprobador,aprobador2,estado_aprobacion, estado_compra, descripcion, fecha,
		       link1, link2, link3, direccion, servicio, tipo_peticion, urgencia, proyecto, coste, usuario, datos_adjunto
		FROM compras_extraordinarias
	`

	// Ejecutar la consulta
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al consultar los datos: " + err.Error()})
		return
	}
	defer rows.Close()

	// Crear un slice para almacenar los resultados
	var compras []models.ComprasExtraordinariasView

	for rows.Next() {
		var compra models.ComprasExtraordinariasView
		var datosAdjunto []byte // Variable para el campo binario

		// Escanear las columnas en la estructura
		err := rows.Scan(
			&compra.ID, &compra.Titulo, &compra.Aprobador, &compra.Aprobador2, &compra.EstadoAprobacion,
			&compra.EstadoCompra, &compra.Descripcion, &compra.Fecha,
			&compra.Link1, &compra.Link2, &compra.Link3, &compra.Direccion,
			&compra.Servicio, &compra.TipoPeticion, &compra.Urgencia, &compra.Proyecto, &compra.Coste, &compra.Usuario,
			&datosAdjunto, // Escanear el binario

		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al escanear los datos: " + err.Error()})
			return
		}

		// Convertir el binario a base64 y asignarlo al campo
		compra.ToBase64(datosAdjunto)

		// Añadir el registro al slice
		compras = append(compras, compra)
	}

	// Comprobar si hay registros
	if len(compras) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "No se encontraron registros en la tabla"})
		return
	}

	// Devolver los datos en formato JSON
	c.JSON(http.StatusOK, compras)
}
