package repository

import (
	"errors"
	"metonode-golang/task3/sql_task/models"

	"github.com/jinzhu/gorm"
)

type StudentRepository interface {
	Insert(students *models.Students) error
	Update(students *models.Students) error
	SelectById(id uint) (*models.Students, error)
}

type studentRepository struct {
	db *gorm.DB
}

func NewStudentRepository(db *gorm.DB) StudentRepository {
	return &studentRepository{db: db}
}

func (s studentRepository) Insert(students *models.Students) error {
	return s.db.Create(students).Error
}

func (s studentRepository) Update(students *models.Students) error {
	if _, err := s.SelectById(students.Id); err != nil {
		return errors.New("student does not exist")
	}

	return s.db.Update(students).Error
}

func (s studentRepository) SelectById(id uint) (*models.Students, error) {
	var students models.Students
	if err := s.db.First(&students, id).Error; err != nil {
		return nil, errors.New("student does not exist")
	}
	return &students, nil
}
