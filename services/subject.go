package services

import "fmt"

type Subject struct {
    SubjectID int    `json:"subject_id" gorm:"primaryKey"`
    Name      string `json:"name"`
}

type SubjectService struct {
    subjects []Subject
}

func NewSubjectService() *SubjectService {
    return &SubjectService{
        subjects: []Subject{},
    }
}

func (s *SubjectService) CreateSubject(subject Subject) (Subject, error) {
    subject.SubjectID = len(s.subjects) + 1
    s.subjects = append(s.subjects, subject)
    return subject, nil
}

func (s *SubjectService) UpdateSubject(id int, updatedSubject Subject) (Subject, error) {
    for i, subject := range s.subjects {
        if subject.SubjectID == id {
            s.subjects[i].Name = updatedSubject.Name
            return s.subjects[i], nil
        }
    }
    return Subject{}, fmt.Errorf("subject not found")
}

func (s *SubjectService) GetSubject(id int) (Subject, error) {
    for _, subject := range s.subjects {
        if subject.SubjectID == id {
            return subject, nil
        }
    }
    return Subject{}, fmt.Errorf("subject not found")
}

func (s *SubjectService) DeleteSubject(id int) error {
    for i, subject := range s.subjects {
        if subject.SubjectID == id {
            s.subjects = append(s.subjects[:i], s.subjects[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("subject not found")
}