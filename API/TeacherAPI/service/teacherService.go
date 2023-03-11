package service

import (
	"TeacherAPI/DTO"
	"TeacherAPI/model"
)

type TeacherService interface {
	GetAll() []model.Teacher
	Save(teacher model.Teacher)
	SearchById(Id int) model.Teacher
	SearchByName(courseName string) []DTO.Teacher
	Update(teacher model.Teacher)
	Delete(teacher model.Teacher)
	NameStartWith(nameStart string) []model.Teacher
}
