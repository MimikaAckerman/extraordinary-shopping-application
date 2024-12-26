package routes

import (
	"bk-compras-extraordinarias/Handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes configura las rutas y middlewares del servidor
func RegisterRoutes() *gin.Engine {
	r := gin.Default()

	// Configurar CORS
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Permitir todas las orígenes
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * 3600, // Tiempo máximo de preflight (12 horas)
	}))

	// Rutas de la API
	r.GET("/", Handler.Welcome)                               // Ruta de prueba para verificar el servidor
	r.POST("/insertdb", Handler.InsertComprasExtraordinarias) // Ruta para insertar datos en la base de datos
	r.GET("/viewdb", Handler.GetComprasExtraordinarias)       // ruta para visualizar todas las compras
	r.POST("/loginuser", Handler.LoginUser)                   //login user
	r.GET("/servicio-centros", Handler.GetServicioCentros)

	return r
}
