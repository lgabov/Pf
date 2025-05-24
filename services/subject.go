package services

import (
	"database/sql"
	"fmt"
)

type Subject struct {
	SubjectID int    `json:"subject_id"`
	Name      string `json:"name"`
}

type SubjectService struct {
	db *sql.DB
}

func NewSubjectService(db *sql.DB) *SubjectService {
	return &SubjectService{
		db: db,
	}
}

func (s *SubjectService) CreateSubject(subject Subject) (Subject, error) {
	if subject.Name == "" {
		return Subject{}, fmt.Errorf("subject name is required")
	}

	query := `
		INSERT INTO subjects (name) 
		VALUES ($1) 
		RETURNING subject_id, name
	`
	
	var createdSubject Subject
	err := s.db.QueryRow(query, subject.Name).Scan(
		&createdSubject.SubjectID,
		&createdSubject.Name,
	)
	
	if err != nil {
		return Subject{}, fmt.Errorf("error creating subject: %v", err)
	}

	return createdSubject, nil
}

func (s *SubjectService) UpdateSubject(id int, updatedSubject Subject) (Subject, error) {
	if updatedSubject.Name == "" {
		return Subject{}, fmt.Errorf("subject name is required")
	}

	query := `
		UPDATE subjects 
		SET name = $1 
		WHERE subject_id = $2
		RETURNING subject_id, name
	`
	
	var subject Subject
	err := s.db.QueryRow(query, updatedSubject.Name, id).Scan(
		&subject.SubjectID,
		&subject.Name,
	)
	
	if err != nil {
		if err == sql.ErrNoRows {
			return Subject{}, fmt.Errorf("subject not found")
		}
		return Subject{}, fmt.Errorf("error updating subject: %v", err)
	}

	return subject, nil
}

func (s *SubjectService) GetSubject(id int) (Subject, error) {
	query := "SELECT subject_id, name FROM subjects WHERE subject_id = $1"
	row := s.db.QueryRow(query, id)

	var subject Subject
	err := row.Scan(&subject.SubjectID, &subject.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return Subject{}, fmt.Errorf("subject not found")
		}
		return Subject{}, fmt.Errorf("error fetching subject: %v", err)
	}

	return subject, nil
}

func (s *SubjectService) DeleteSubject(id int) error {
	query := "DELETE FROM subjects WHERE subject_id = $1"
	result, err := s.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error deleting subject: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("subject not found")
	}

	return nil
}

func (s *SubjectService) GetAllSubjects() ([]Subject, error) {
	query := "SELECT subject_id, name FROM subjects"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching subjects: %v", err)
	}
	defer rows.Close()

	var subjects []Subject
	for rows.Next() {
		var subject Subject
		err := rows.Scan(&subject.SubjectID, &subject.Name)
		if err != nil {
			return nil, fmt.Errorf("error scanning subject: %v", err)
		}
		subjects = append(subjects, subject)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating subjects: %v", err)
	}

	return subjects, nil
}