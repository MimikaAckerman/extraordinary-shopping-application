package Handler

import (
	"bk-compras-extraordinarias/config"
	"bk-compras-extraordinarias/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetServicioCentros(c *gin.Context) {
	// Consulta SQL
	query := `select distinct uc.Center , uc.CodNav ,uc.Zone from UbUserCenters uuc
				join UbCenters uc on uuc.PkCenterId=uc.PkCenter
				where uuc.IsOwner=1 and CodNav like 'S0%' `

	// Ejecutar la consulta
	rows, err := config.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al ejecutar la consulta: " + err.Error()})
		return
	}
	defer rows.Close()

	// Recoger los resultados
	var resultados []models.ServicioCentro
	for rows.Next() {
		var centro models.ServicioCentro
		if err := rows.Scan(&centro.Center, &centro.CodNav, &centro.Zone); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar los resultados: " + err.Error()})
			return
		}

		// Si Zone es NULL, reemplázalo con un valor vacío
		if !centro.Zone.Valid {
			centro.Zone.String = ""
		}

		resultados = append(resultados, centro)
	}

	// Verificar errores en la iteración
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en la iteración de resultados: " + err.Error()})
		return
	}

	// Devolver los resultados como JSON
	c.JSON(http.StatusOK, resultados)
}
