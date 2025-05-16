package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"pf/middleware"
	"pf/routes"
	"pf/services"
)

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerMiddleware())

	// Crear servicio
	studentService := services.NewStudentService()

	// Configurar rutas
	routes.SetupStudentRoutes(r, studentService)

	// Iniciar servidor
	fmt.Println("Server running on port 3000")
	r.Run(":3000")
}
