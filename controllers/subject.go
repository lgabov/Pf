package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pf/services"
)

type SubjectController struct {
	subjectService *services.SubjectService
}

func NewSubjectController(subjectService *services.SubjectService) *SubjectController {
	return &SubjectController{
		subjectService: subjectService,
	}
}

func (s *SubjectController) CreateSubject(c *gin.Context) {
	var subject services.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	createdSubject, err := s.subjectService.CreateSubject(subject)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error to create subject",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, createdSubject)
}

func (s *SubjectController) GetSubject(c *gin.Context) {
	idStr := c.Param("subject_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid subject ID",
			"error":   err.Error(),
		})
		return
	}

	subject, err := s.subjectService.GetSubject(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "subject not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, subject)
}

func (s *SubjectController) UpdateSubject(c *gin.Context) {
	idStr := c.Param("subject_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid subject ID",
			"error":   err.Error(),
		})
		return
	}

	var subject services.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request body",
			"error":   err.Error(),
		})
		return
	}

	updatedSubject, err := s.subjectService.UpdateSubject(id, subject)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "subject not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, updatedSubject)
}

func (s *SubjectController) DeleteSubject(c *gin.Context) {
	idStr := c.Param("subject_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid subject ID",
			"error":   err.Error(),
		})
		return
	}

	err = s.subjectService.DeleteSubject(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "subject not found",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "subject deleted successfully",
		"id":      id,
	})
}