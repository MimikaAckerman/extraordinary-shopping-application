package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

// Declarar la variable DB como una variable global
var DB *sql.DB
var DBSecondary *sql.DB

func ConnectDB() {
	// Credenciales de conexión directamente en el código
	server := "ubserveraz.database.windows.net"
	user := "ubadmin"
	password := "EZFRsuw$QowTrYoV"
	database := "ubProd"

	// Crear la cadena de conexión para SQL Server
	connectionString := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=true&trustServerCertificate=true", user, password, server, database)

	// Abrir la conexión a la base de datos
	var err error
	DB, err = sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatalf("Error creando la conexión a la base de datos: %v", err)
	}

	// Verificar la conexión
	err = DB.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos exitosa")

	//segunda conexion - base de datos secundaria
	database2 := "UbHR"
	connectionString2 := fmt.Sprintf("sqlserver://%s:%s@%s?database=%s&encrypt=true&trustServerCertificate=true", user, password, server, database2)
	DBSecondary, err = sql.Open("sqlserver", connectionString2)
	if err != nil {
		log.Fatalf("Error creando la conexión a la base de datos secundaria: %v", err)
	}

	err = DBSecondary.Ping()
	if err != nil {
		log.Fatalf("No se pudo conectar a la base de datos secundaria: %v", err)
	}

	log.Println("Conexión a la base de datos secundaria exitosa")
}

func GetPort() string {
	// Verificar primero la variable de entorno estándar de Render
	port := os.Getenv("PORT")
	if port == "" {
		// Verificar otra variable de entorno que podrías estar usando
		port = os.Getenv("HTTP_PLATFORM_PORT")
		if port == "" {
			port = "8000" // Valor predeterminado si no se proporciona un puerto
		}
	}
	return port
}
