package Handler

import (
	"bk-compras-extraordinarias/config"
	"fmt"

	"bk-compras-extraordinarias/models"
	"database/sql"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InsertComprasExtraordinarias(c *gin.Context) {
	// Vincular los datos del formulario
	var prodData models.ComprasExtraordinarias
	if err := c.ShouldBind(&prodData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Valores predeterminados
	prodData.EstadoAprobacion = "Pendiente"
	prodData.EstadoAprobacion2 = "Pendiente"
	prodData.EstadoCompra = "Pendiente"
	prodData.Aprobador = "Pendiente"
	prodData.Aprobador2 = "Pendiente"
	prodData.Fecha = time.Now().Format("2006-01-02T15:04:05")

	// Manejar el archivo adjunto (opcional)
	var datosAdjuntoBytes []byte
	if file, err := c.FormFile("DatosAdjunto"); err == nil {
		src, err := file.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al abrir el archivo adjunto: " + err.Error()})
			return
		}
		defer src.Close()

		// Leer el archivo completo
		datosAdjuntoBytes, err = io.ReadAll(src)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer el archivo adjunto: " + err.Error()})
			return
		}

		// Obtener el nombre del archivo y su extensión
		prodData.NombreAdjunto = file.Filename
	} else {
		// Valores predeterminados si no se proporciona un archivo
		datosAdjuntoBytes = nil
		prodData.NombreAdjunto = ""
	}

	// Preparar la consulta SQL
	query := `
		INSERT INTO compras_extraordinarias (
			titulo, aprobador, estado_aprobacion, estado_aprobacion2, estado_compra, descripcion, fecha, 
			link1, link2, link3, direccion, servicio, tipo_peticion, urgencia, 
			proyecto, datos_adjunto, nombre_adjunto, coste, aprobador2, usuario, centro
		)
		VALUES (
			@p1, @p2, @p3, @p4, @p5, @p6, @p7, 
			@p8, @p9, @p10, @p11, @p12, @p13, @p14, 
			@p15, @p16, @p17, @p18, @p19, @p20, @p21
		)
	`

	// Ejecutar la consulta
	_, err := config.DB.Exec(query,
		sql.Named("p1", prodData.Titulo),
		sql.Named("p2", prodData.Aprobador),
		sql.Named("p3", prodData.EstadoAprobacion),
		sql.Named("p4", prodData.EstadoAprobacion2),
		sql.Named("p5", prodData.EstadoCompra),
		sql.Named("p6", prodData.Descripcion),
		sql.Named("p7", prodData.Fecha),
		sql.Named("p8", prodData.Link1),
		sql.Named("p9", prodData.Link2),
		sql.Named("p10", prodData.Link3),
		sql.Named("p11", prodData.Direccion),
		sql.Named("p12", prodData.Servicio),
		sql.Named("p13", prodData.TipoPeticion),
		sql.Named("p14", prodData.Urgencia),
		sql.Named("p15", prodData.Proyecto),
		sql.Named("p16", datosAdjuntoBytes), // Archivo adjunto
		sql.Named("p17", prodData.NombreAdjunto),
		sql.Named("p18", prodData.Coste),
		sql.Named("p19", prodData.Aprobador2),
		sql.Named("p20", prodData.Usuario),
		sql.Named("p21", prodData.Centro),
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al insertar los datos: " + err.Error()})
		return
	}

	// Enviar correo de resguardo al solicitante
	err = EnviarCorreoSolicitante(prodData.Titulo, prodData.Descripcion, prodData.Proyecto, fmt.Sprintf("%.2f", prodData.Coste), prodData.Usuario)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el correo al solicitante: " + err.Error()})
		return
	}

	// Respuesta de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Datos insertados correctamente"})
}
