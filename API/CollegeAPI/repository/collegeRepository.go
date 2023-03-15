package repository

import (
	"CollegeAPI/DTO"
	"CollegeAPI/model"
)

type CollegeRepository interface {
	FindAllStudents() []model.Student
	SearchByStudentRoll(roll int) DTO.StudentDTO
	SearchByStudentName(name string) []DTO.StudentDTO
	SortByStudentName() []DTO.StudentDTO
	SortByStudentMarks() []DTO.StudentDTO
	SearchByStudentByCity(city string) []DTO.StudentDTO
	RegisterStudent(student model.Student) error
	NameSearchStartWith(name string) []DTO.StudentDTO
	StudentsCheck(students []model.Student) []string
	RegisterTeacher(teacher model.Teacher) error
	ValidationData(teacher model.Teacher) []model.Student
}
