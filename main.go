package main

import (
	"bk-compras-extraordinarias/config"
	"bk-compras-extraordinarias/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Establecer el modo de Gin en 'release' para producción
	gin.SetMode(gin.ReleaseMode)

	// Conectar a la base de datos
	config.ConnectDB()

	// Registrar las rutas
	r := routes.RegisterRoutes()

	// Obtener el puerto desde la configuración
	PORT := config.GetPort()

	// Iniciar el servidor en el puerto especificado
	log.Printf("Servidor corriendo en http://localhost:%s", PORT)
	if err := r.Run(":" + PORT); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
