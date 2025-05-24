package services

import (
	"database/sql"
	"fmt"
)

type Grade struct {
	GradeID   int     `json:"grade_id"`
	StudentID int     `json:"student_id"`
	SubjectID int     `json:"subject_id"`
	Grade     float64 `json:"grade"`
}

type GradeService struct {
	db *sql.DB
}

func NewGradeService(db *sql.DB) *GradeService {
	return &GradeService{
		db: db,
	}
}

func (s *GradeService) CreateGrade(grade Grade) (Grade, error) {
	// Validaciones básicas
	if grade.StudentID == 0 {
		return Grade{}, fmt.Errorf("student ID is required")
	}
	if grade.SubjectID == 0 {
		return Grade{}, fmt.Errorf("subject ID is required")
	}
	if grade.Grade < 0 || grade.Grade > 10 {
		return Grade{}, fmt.Errorf("grade must be between 0 and 10")
	}

	query := `
		INSERT INTO grades (student_id, subject_id, grade) 
		VALUES ($1, $2, $3)
		RETURNING grade_id, student_id, subject_id, grade
	`
	
	var createdGrade Grade
	err := s.db.QueryRow(
		query,
		grade.StudentID,
		grade.SubjectID,
		grade.Grade,
	).Scan(
		&createdGrade.GradeID,
		&createdGrade.StudentID,
		&createdGrade.SubjectID,
		&createdGrade.Grade,
	)
	
	if err != nil {
		return Grade{}, fmt.Errorf("error creating grade: %v", err)
	}

	return createdGrade, nil
}

func (s *GradeService) UpdateGrade(id int, updatedGrade Grade) (Grade, error) {
	// Validaciones
	if updatedGrade.Grade < 0 || updatedGrade.Grade > 10 {
		return Grade{}, fmt.Errorf("grade must be between 0 and 10")
	}

	query := `
		UPDATE grades 
		SET student_id = $1, subject_id = $2, grade = $3
		WHERE grade_id = $4
		RETURNING grade_id, student_id, subject_id, grade
	`
	
	var grade Grade
	err := s.db.QueryRow(
		query,
		updatedGrade.StudentID,
		updatedGrade.SubjectID,
		updatedGrade.Grade,
		id,
	).Scan(
		&grade.GradeID,
		&grade.StudentID,
		&grade.SubjectID,
		&grade.Grade,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return Grade{}, fmt.Errorf("grade not found")
		}
		return Grade{}, fmt.Errorf("error updating grade: %v", err)
	}

	return grade, nil
}

func (s *GradeService) GetGrade(id int) (Grade, error) {
	query := `
		SELECT grade_id, student_id, subject_id, grade 
		FROM grades 
		WHERE grade_id = $1
	`
	row := s.db.QueryRow(query, id)

	var grade Grade
	err := row.Scan(
		&grade.GradeID,
		&grade.StudentID,
		&grade.SubjectID,
		&grade.Grade,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Grade{}, fmt.Errorf("grade not found")
		}
		return Grade{}, fmt.Errorf("error fetching grade: %v", err)
	}

	return grade, nil
}

func (s *GradeService) GetGradesByStudent(studentID int) ([]Grade, error) {
	query := `
		SELECT grade_id, student_id, subject_id, grade 
		FROM grades 
		WHERE student_id = $1
		ORDER BY subject_id
	`
	rows, err := s.db.Query(query, studentID)
	if err != nil {
		return nil, fmt.Errorf("error fetching grades: %v", err)
	}
	defer rows.Close()

	var grades []Grade
	for rows.Next() {
		var grade Grade
		err := rows.Scan(
			&grade.GradeID,
			&grade.StudentID,
			&grade.SubjectID,
			&grade.Grade,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning grade: %v", err)
		}
		grades = append(grades, grade)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating grades: %v", err)
	}

	if len(grades) == 0 {
		return nil, fmt.Errorf("no grades found for student")
	}

	return grades, nil
}

func (s *GradeService) DeleteGrade(id int) error {
	query := "DELETE FROM grades WHERE grade_id = $1"
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting grade: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("grade not found")
	}

	return nil
}

// Método adicional para obtener el promedio de un estudiante
func (s *GradeService) GetStudentAverage(studentID int) (float64, error) {
	query := `
		SELECT AVG(grade) 
		FROM grades 
		WHERE student_id = $1
	`
	row := s.db.QueryRow(query, studentID)

	var average *float64 // Usamos puntero para manejar NULL
	err := row.Scan(&average)
	if err != nil {
		return 0, fmt.Errorf("error calculating average: %v", err)
	}

	if average == nil {
		return 0, fmt.Errorf("no grades found for student")
	}

	return *average, nil
}