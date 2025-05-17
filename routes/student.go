package routes

import (
	"github.com/gin-gonic/gin"
	"pf/controllers"
	"pf/services"
)

func SetupAllRoutes(r *gin.Engine, studentService *services.StudentService,
	 subjectService *services.SubjectService, gradeService *services.GradeService,) {
	// Configurar rutas 
	setupStudentRoutes(r, studentService)
	
	setupSubjectRoutes(r, subjectService)

	SetupGradeRoutes(r, gradeService)
}

func setupStudentRoutes(r *gin.Engine, studentService *services.StudentService) {
	studentController := controllers.NewStudentController(studentService)
	
	studentRoutes := r.Group("/students")
	{
		studentRoutes.GET("/:id", studentController.GetStudent)
		studentRoutes.POST("/", studentController.CreateStudent)
		studentRoutes.PUT("/:id", studentController.UpdateStudent)
		studentRoutes.DELETE("/:id", studentController.DeleteStudent)
	}
}

func setupSubjectRoutes(r *gin.Engine, subjectService *services.SubjectService) {
	subjectController := controllers.NewSubjectController(subjectService)
	
	subjectRoutes := r.Group("/api/subjects")
	{
		subjectRoutes.POST("/", subjectController.CreateSubject)
		subjectRoutes.PUT("/:subject_id", subjectController.UpdateSubject)
		subjectRoutes.GET("/:subject_id", subjectController.GetSubject)
		subjectRoutes.DELETE("/:subject_id", subjectController.DeleteSubject)
	}
}
func SetupGradeRoutes(r *gin.Engine, gradeService *services.GradeService) {
    gradeController := controllers.NewGradeController(gradeService)
	gradeRoutes := r.Group("/api/grades")
    {
        gradeRoutes.POST("/", gradeController.CreateGrade)
        gradeRoutes.PUT("/:grade_id", gradeController.UpdateGrade)
        gradeRoutes.DELETE("/:grade_id", gradeController.DeleteGrade)
        gradeRoutes.GET("/:grade_id/student/:student_id", gradeController.GetGrade)
        gradeRoutes.GET("/student/:student_id", gradeController.GetGradesByStudent)
    }
}

