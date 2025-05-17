package controllers

import (
    "net/http"
    "strconv"
    "github.com/gin-gonic/gin"
    "pf/services"
)

type GradeController struct {
    gradeService *services.GradeService
}

func NewGradeController(gradeService *services.GradeService) *GradeController {
    return &GradeController{
        gradeService: gradeService,
    }
}

func (g *GradeController) CreateGrade(c *gin.Context) {
    var grade services.Grade
    if err := c.ShouldBindJSON(&grade); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid request body",
            "error":   err.Error(),
        })
        return
    }

    createdGrade, err := g.gradeService.CreateGrade(grade)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "message": "Error creating grade",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(http.StatusCreated, createdGrade)
}

func (g *GradeController) UpdateGrade(c *gin.Context) {
    idStr := c.Param("grade_id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid grade ID",
            "error":   err.Error(),
        })
        return
    }

    var grade services.Grade
    if err := c.ShouldBindJSON(&grade); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid request body",
            "error":   err.Error(),
        })
        return
    }

    updatedGrade, err := g.gradeService.UpdateGrade(id, grade)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "grade not found",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, updatedGrade)
}

func (g *GradeController) GetGrade(c *gin.Context) {
    gradeIDStr := c.Param("grade_id")
    studentIDStr := c.Param("student_id")

    gradeID, err := strconv.Atoi(gradeIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid grade ID",
            "error":   err.Error(),
        })
        return
    }

    _, err = strconv.Atoi(studentIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid student ID",
            "error":   err.Error(),
        })
        return
    }

    grade, err := g.gradeService.GetGrade(gradeID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "grade not found",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, grade)
}

func (g *GradeController) GetGradesByStudent(c *gin.Context) {
    studentIDStr := c.Param("student_id")
    studentID, err := strconv.Atoi(studentIDStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid student ID",
            "error":   err.Error(),
        })
        return
    }

    grades, err := g.gradeService.GetGradesByStudent(studentID)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "no grades found for student",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, grades)
}

func (g *GradeController) DeleteGrade(c *gin.Context) {
    idStr := c.Param("grade_id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "message": "invalid grade ID",
            "error":   err.Error(),
        })
        return
    }

    err = g.gradeService.DeleteGrade(id)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{
            "message": "grade not found",
            "error":   err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "grade deleted successfully",
        "id":      id,
    })
}