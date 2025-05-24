CREATE TABLE students (
    id SERIAL PRIMARY KEY,
    group_name VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()

);

CREATE TABLE subjects (
    subject_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()

);

CREATE TABLE grades (
    grade_id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    student_id INT NOT NULL,  
    subject_id INT NOT NULL,
    grade FLOAT NOT NULL CHECK (grade >= 0 AND grade <= 10), 
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE, 
    FOREIGN KEY (subject_id) REFERENCES subjects(subject_id) ON DELETE RESTRICT, 
    UNIQUE (student_id, subject_id)  
);

