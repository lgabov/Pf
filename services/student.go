package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

var (
	ErrStudentNotFound  = errors.New("student not found")
	ErrInvalidInput     = errors.New("invalid input data")
	ErrDatabase         = errors.New("database error")
)

type Student struct {
	ID        int       `json:"id"`
	Group     string    `json:"group"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type StudentService struct {
	db *sql.DB
}

func NewStudentService(db *sql.DB) *StudentService {
	return &StudentService{db: db}
}

// GetStudents retrieves all students with pagination support
func (s *StudentService) GetStudents(ctx context.Context, limit, offset int) ([]Student, error) {
	query := `
		SELECT id, group_name, name, email, created_at
		FROM students 
		ORDER BY id 
		LIMIT $1 OFFSET $2
	`

	rows, err := s.db.QueryContext(ctx, query, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}
	defer rows.Close()

	var students []Student
	for rows.Next() {
		var student Student
		err := rows.Scan(
			&student.ID,
			&student.Group,
			&student.Name,
			&student.Email,
			&student.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
		}
		students = append(students, student)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return students, nil
}

// GetStudentByID retrieves a single student by ID
func (s *StudentService) GetStudentByID(ctx context.Context, id int) (*Student, error) {
	query := `
		SELECT id, group_name, name, email, created_at
		FROM students 
		WHERE id = $1
	`

	var student Student
	err := s.db.QueryRowContext(ctx, query, id).Scan(
		&student.ID,
		&student.Group,
		&student.Name,
		&student.Email,
		&student.CreatedAt,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, ErrStudentNotFound
	case err != nil:
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return &student, nil
}

// CreateStudent creates a new student record
func (s *StudentService) CreateStudent(ctx context.Context, student *Student) (*Student, error) {
	if student.Name == "" || student.Email == "" {
		return nil, fmt.Errorf("%w: name and email are required", ErrInvalidInput)
	}

	query := `
		INSERT INTO students (group_name, name, email) 
		VALUES ($1, $2, $3) 
		RETURNING id, group_name, name, email, created_at
	`

	err := s.db.QueryRowContext(ctx, query,
		student.Group,
		student.Name,
		student.Email,
	).Scan(
		&student.ID,
		&student.Group,
		&student.Name,
		&student.Email,
		&student.CreatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return student, nil
}

// UpdateStudent updates an existing student record
func (s *StudentService) UpdateStudent(ctx context.Context, id int, student *Student) (*Student, error) {
	query := `
		UPDATE students 
		SET group_name = $1, name = $2, email = $3 
		WHERE id = $4
		RETURNING id, group_name, name, email, created_at
	`

	var updatedStudent Student
	err := s.db.QueryRowContext(ctx, query,
		student.Group,
		student.Name,
		student.Email,
		id,
	).Scan(
		&updatedStudent.ID,
		&updatedStudent.Group,
		&updatedStudent.Name,
		&updatedStudent.Email,
		&updatedStudent.CreatedAt,
	)

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return nil, ErrStudentNotFound
	case err != nil:
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return &updatedStudent, nil
}

// DeleteStudent deletes a student record
func (s *StudentService) DeleteStudent(ctx context.Context, id int) error {
	query := "DELETE FROM students WHERE id = $1"
	result, err := s.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	if rowsAffected == 0 {
		return ErrStudentNotFound
	}

	return nil
}