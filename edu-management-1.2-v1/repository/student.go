package repository

import (
	"a21hc3NpZ25tZW50/model"
	"database/sql"
)

type StudentRepository interface {
	FetchAll() ([]model.Student, error)
	FetchByID(id int) (*model.Student, error)
	Store(s *model.Student) error
	Delete(id int) error
}

type studentRepoImpl struct {
	db *sql.DB
}

func NewStudentRepo(db *sql.DB) *studentRepoImpl {
	return &studentRepoImpl{db}
}

func (s *studentRepoImpl) FetchAll() ([]model.Student, error) {
	rows, err := s.db.Query("SELECT * FROM students")
	if err != nil {
		return nil, err
	}

	var data []model.Student

	for rows.Next() {
		var student model.Student
		err := rows.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
		if err != nil {
			return nil, err
		}

		data = append(data, student)
	}

	return data, nil // TODO: replace this
}

func (s *studentRepoImpl) FetchByID(id int) (*model.Student, error) {
	row := s.db.QueryRow("SELECT id, name, address, class FROM students WHERE id = $1", id)

	var student model.Student
	err := row.Scan(&student.ID, &student.Name, &student.Address, &student.Class)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (s *studentRepoImpl) Store(student *model.Student) error {
	_, err := s.db.Exec("INSERT INTO students(name, address, class) VALUES($1, $2, $3) RETURNING id", student.Name, student.Address, student.Class)
	if err != nil {
		return err
	}
	return nil
}

func (s *studentRepoImpl) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM students WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}
