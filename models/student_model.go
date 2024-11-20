package models

import (
	"database/sql"
	"fiqri_muhammad/student-api/config"
)

type Student struct {
	ID       int    `json:"id,omitempty"`
	NIM      string `json:"nim"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Faculty  string `json:"faculty"`
	Semester int    `json:"semester"`
}

func (s *Student) Create() error {
	stmt, err := config.DB.Prepare("INSERT INTO students(nim, name, email, gender, faculty, semester) VALUES(?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(s.NIM, s.Name, s.Email, s.Gender, s.Faculty, s.Semester)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	s.ID = int(id)
	return nil
}

func GetAllStudents() ([]Student, error) {
	var students []Student

	rows, err := config.DB.Query("SELECT id, nim, name, email, gender, faculty, semester FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var student Student
		err := rows.Scan(&student.ID, &student.NIM, &student.Name, &student.Email,
			&student.Gender, &student.Faculty, &student.Semester)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}

	return students, nil
}

func GetStudentByID(id int) (Student, error) {
	var student Student

	err := config.DB.QueryRow("SELECT id, nim, name, email, gender, faculty, semester FROM students WHERE id = ?",
		id).Scan(&student.ID, &student.NIM, &student.Name, &student.Email,
		&student.Gender, &student.Faculty, &student.Semester)
	if err == sql.ErrNoRows {
		return student, err
	}
	if err != nil {
		return student, err
	}

	return student, nil
}

func (s *Student) Update() error {
	stmt, err := config.DB.Prepare("UPDATE students SET nim=?, name=?, email=?, gender=?, faculty=?, semester=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(s.NIM, s.Name, s.Email, s.Gender, s.Faculty, s.Semester, s.ID)
	return err
}

func DeleteStudent(id int) error {
	stmt, err := config.DB.Prepare("DELETE FROM students WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	return err
}
