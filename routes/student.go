package routes

import (
	"github.com/gin-gonic/gin"
	"pf/controllers"
	"pf/services"
)

func SetupStudentRoutes(r *gin.Engine, studentService *services.StudentService) {
	studentController := controllers.NewStudentController(studentService)
	
	// Grupo de rutas para estudiantes
	studentRoutes := r.Group("/students")
	{
		studentRoutes.GET("/:id", studentController.GetStudent)
		studentRoutes.POST("/", studentController.CreateStudent)
		studentRoutes.PUT("/:id", studentController.UpdateStudent)
		studentRoutes.DELETE("/:id", studentController.DeleteStudent)
	}
}