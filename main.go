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
	subjectService := services.NewSubjectService()
	gradeService := services.NewGradeService()
 
	// Configurar rutas
	routes.SetupAllRoutes(r, studentService, subjectService, gradeService)

	// Iniciar servidor
	fmt.Println("Server running on port 3000")
	r.Run(":3000")
}
