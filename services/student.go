package services

import "fmt"

type Student struct {  
    ID    int    `json:"id"`
    Group string `json:"group"`
    Name  string `json:"name"` 
    Email string `json:"email"`
}

type StudentService struct {
    students []Student
}

func NewStudentService() *StudentService {  
    return &StudentService{
        students: []Student{},
    }
}

func (s *StudentService) GetStudents() []Student {
    return s.students
}

func (s *StudentService) GetStudent(id int) (Student, error) {
    for _, student := range s.students {
        if student.ID == id {
            return student, nil
        }
    }
    return Student{}, fmt.Errorf("student not found")
}

func (s *StudentService) CreateStudent(student Student) (Student, error) {
    if student.Name == "" || student.Email == "" {
        return Student{}, fmt.Errorf("name and email are required")
    }
    student.ID = len(s.students) + 1
    s.students = append(s.students, student)
    return student, nil
}

func (s *StudentService) UpdateStudent(id int, updatedStudent Student) (Student, error) {
    for i, student := range s.students {  
        if student.ID == id {
            s.students[i].Group = updatedStudent.Group
            s.students[i].Name = updatedStudent.Name
            s.students[i].Email = updatedStudent.Email
            return s.students[i], nil  
        }
    }
    return Student{}, fmt.Errorf("student not found")  
}

func (s *StudentService) DeleteStudent(id int) error {
    for i, student := range s.students {
        if student.ID == id {
            s.students = append(s.students[:i], s.students[i+1:]...)
            return nil  
        }
    }
    return fmt.Errorf("student not found")
}
