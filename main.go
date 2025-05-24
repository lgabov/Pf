package main

import (
	"fmt"
	"log"
	"pf/database"
	"pf/routes"
	"pf/services"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	//  Inicializar la conexión a la base de datos
	db, err := database.NewDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Db.Close()

	//  Verificar la conexión a la base de datos
	err = db.Db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	fmt.Println("Successfully connected to database!")

	// Configurar Gin
	r := gin.Default()

	// Inicializar servicios con la conexión a la base de datos
	studentService := services.NewStudentService(db.Db)
	subjectService := services.NewSubjectService(db.Db)
	gradeService := services.NewGradeService(db.Db)

	// Configurar rutas
	routes.SetupAllRoutes(r, studentService, subjectService, gradeService)

	// Iniciar servidor
	fmt.Println("Server running on port 3000")
	if err := r.Run(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}