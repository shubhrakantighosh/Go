package service

import (
	"CollegeAPI/DTO"
	"CollegeAPI/model"
	"CollegeAPI/repository"
	"errors"
)

type CollegeServiceImpl struct {
	CollegeRepository repository.CollegeRepository
}

func NewCollegeServiceImpl(college repository.CollegeRepository) CollegeService {
	return &CollegeServiceImpl{CollegeRepository: college}
}

func (college *CollegeServiceImpl) FindAllStudents() []DTO.Student_Details_DTO {
	students := college.CollegeRepository.FindAllStudents()
	students_DTO := []DTO.Student_Details_DTO{}
	for _, i := range students {
		students_DTO = append(students_DTO, DTO.Student_Details_DTO{i.Roll, i.FirstName, i.LastName, i.Age, i.Grade, i.Address.PinCode, i.Address.City, i.Address.State})
	}
	return students_DTO
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
	s := college.SearchByStudentRoll(student.Roll)
	if s.Roll != 0 {
		return errors.New("Already Exists")
	}
	student.Grade = model.GradeCheck(student.Grade)
	err := college.CollegeRepository.RegisterStudent(student)
	return err
}

func (college *CollegeServiceImpl) NameSearchStartWith(name string) []DTO.StudentDTO {
	return college.CollegeRepository.NameSearchStartWith(name)
}

func (college *CollegeServiceImpl) StudentsCheck(students []model.Student) []string {
	return college.CollegeRepository.StudentsCheck(students)
}

func (college *CollegeServiceImpl) RegisterTeacher(teacher model.Teacher) error {
	err := college.CollegeRepository.RegisterTeacher(teacher)
	return err
}

func (college *CollegeServiceImpl) ValidationData(teacher model.Teacher) []string {

	result := []string{}

	students := college.CollegeRepository.ValidationData(teacher)

	if students == nil {
		return nil
	}

	for i, j := range students {
		if j.ID == 0 {
			result = append(result, teacher.Student[i].FirstName+": Not Found")
		} else {
			result = append(result, j.FirstName+": Verified")
		}
	}
	return result

}
