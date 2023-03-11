package repository

import (
	"TeacherAPI/DTO"
	"TeacherAPI/model"
)

type TeacherRepository interface {
	GetAll() []model.Teacher
	Save(teacher model.Teacher)
	SearchById(Id int) model.Teacher
	SearchByCourseName(courseName string) []DTO.Teacher
	Update(teacher model.Teacher)
	Delete(teacher model.Teacher)
	NameStartWith(nameStart string) []model.Teacher
}
