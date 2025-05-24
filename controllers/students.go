package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pf/services" 
)

type StudentController struct {
	studentService *services.StudentService
}

func NewStudentController(studentService *services.StudentService) *StudentController {
	return &StudentController{
		studentService: studentService,
	}
}

func (s *StudentController) GetStudent(c *gin.Context) {
	idStr := c.Param("student_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid student ID",
			"error":   err.Error(),
		})
		return
	}

	student, err := s.studentService.GetStudentByID(c.Request.Context(), id)
	if err != nil {
		if err == services.ErrStudentNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "student not found",
				"error":   err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error getting student",
				"error":   err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, student)
}

func (s *StudentController) CreateStudent(c *gin.Context) {
	var student services.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	createdStudent, err := s.studentService.CreateStudent(c.Request.Context(), &student)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "error creating student",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdStudent)
}

func (s *StudentController) UpdateStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid student ID",
			"error":   err.Error(),
		})
		return
	}

	var student services.Student
	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	updatedStudent, err := s.studentService.UpdateStudent(c.Request.Context(), id, &student)
	if err != nil {
		if err == services.ErrStudentNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "student not found",
				"error":   err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error updating student",
				"error":   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, updatedStudent)
}

func (s *StudentController) DeleteStudent(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid student ID",
			"error":   err.Error(),
		})
		return
	}

	err = s.studentService.DeleteStudent(c.Request.Context(), id)
	if err != nil {
		if err == services.ErrStudentNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "student not found",
				"error":   err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "error deleting student",
				"error":   err.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted successfully",
		"id":      id,
	})
}