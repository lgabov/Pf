package services

import "fmt"

type Grade struct {
    GradeID   int     `json:"grade_id" gorm:"primaryKey"`
    StudentID int     `json:"student_id"`
    SubjectID int     `json:"subject_id"`
    Grade     float64 `json:"grade"`
}

type GradeService struct {
    grades []Grade
}

func NewGradeService() *GradeService {
    return &GradeService{
        grades: []Grade{},
    }
}

func (s *GradeService) CreateGrade(grade Grade) (Grade, error) {
    grade.GradeID = len(s.grades) + 1
    s.grades = append(s.grades, grade)
    return grade, nil
}

func (s *GradeService) UpdateGrade(id int, updatedGrade Grade) (Grade, error) {
    for i, grade := range s.grades {
        if grade.GradeID == id {
            s.grades[i] = updatedGrade
            return s.grades[i], nil
        }
    }
    return Grade{}, fmt.Errorf("grade not found")
}

func (s *GradeService) GetGrade(id int) (Grade, error) {
    for _, grade := range s.grades {
        if grade.GradeID == id {
            return grade, nil
        }
    }
    return Grade{}, fmt.Errorf("grade not found")
}

func (s *GradeService) GetGradesByStudent(studentID int) ([]Grade, error) {
    var result []Grade
    for _, grade := range s.grades {
        if grade.StudentID == studentID {
            result = append(result, grade)
        }
    }
    if len(result) == 0 {
        return nil, fmt.Errorf("no grades found for student")
    }
    return result, nil
}

func (s *GradeService) DeleteGrade(id int) error {
    for i, grade := range s.grades {
        if grade.GradeID == id {
            s.grades = append(s.grades[:i], s.grades[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("grade not found")
}