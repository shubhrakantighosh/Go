package service

import (
	"CollegeAPI/DTO"
	"CollegeAPI/model"
	"CollegeAPI/repository"
)

type CollegeServiceImpl struct {
	CollegeRepository repository.CollegeRepository
}

func NewCollegeServiceImpl(college repository.CollegeRepository) CollegeService {
	return &CollegeServiceImpl{CollegeRepository: college}
}

func (college *CollegeServiceImpl) FindAllStudents() []model.Student {
	return college.CollegeRepository.FindAllStudents()
}

func (college *CollegeServiceImpl) SearchByStudentRoll(roll int) DTO.StudentDTO {
	return college.CollegeRepository.SearchByStudentRoll(roll)
}

func (college *CollegeServiceImpl) SearchByStudentName(name string) []DTO.StudentDTO {
	return college.CollegeRepository.SearchByStudentName(name)
}

func (college *CollegeServiceImpl) SortByStudentName() []DTO.StudentDTO {
	return college.CollegeRepository.SortByStudentName()
}

func (college *CollegeServiceImpl) SortByStudentMarks() []DTO.StudentDTO {
	return college.CollegeRepository.SortByStudentMarks()
}

func (college *CollegeServiceImpl) SearchByStudentByCity(city string) []DTO.StudentDTO {
	return college.CollegeRepository.SearchByStudentByCity(city)
}

func (college *CollegeServiceImpl) RegisterStudent(student model.Student) error {
	return college.CollegeRepository.RegisterStudent(student)
}

func (college *CollegeServiceImpl) NameSearchStartWith(name string) []DTO.StudentDTO {
	return college.CollegeRepository.NameSearchStartWith(name)
}

func (college *CollegeServiceImpl) StudentsCheck(students []model.Student) []string {
	return college.CollegeRepository.StudentsCheck(students)
}
