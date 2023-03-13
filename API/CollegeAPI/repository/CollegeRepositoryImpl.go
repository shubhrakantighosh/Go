package repository

import (
	"CollegeAPI/DTO"
	"CollegeAPI/model"
	"gorm.io/gorm"
	"sort"
)

type CollegeRepositoryImpl struct {
	Db *gorm.DB
}

func NewCollegeRepositoryImpl(Db *gorm.DB) CollegeRepository {
	return &CollegeRepositoryImpl{Db: Db}
}

func (c *CollegeRepositoryImpl) FindAllStudents() []model.Student {
	students := []model.Student{}
	c.Db.Find(&students)
	return students
}

func (c *CollegeRepositoryImpl) SearchByStudentRoll(roll int) DTO.StudentDTO {
	student := model.Student{}
	c.Db.Find(&student, roll)
	return DTO.StudentDTO{student.Roll, student.FirstName, student.LastName, student.Age, student.Marks}
}

func (c *CollegeRepositoryImpl) SearchByStudentName(name string) []DTO.StudentDTO {
	students := []model.Student{}
	query := "first_name='" + name + "'"
	c.Db.Find(&students, query)
	student_dto := []DTO.StudentDTO{}

	for _, i := range students {
		student_dto = append(student_dto, DTO.StudentDTO{i.Roll, i.FirstName, i.LastName, i.Age, i.Marks})
	}
	return student_dto
}

func (c *CollegeRepositoryImpl) SortByStudentName() []DTO.StudentDTO {
	students := []model.Student{}
	c.Db.Find(&students)

	students_dto := []DTO.StudentDTO{}

	for _, i := range students {
		students_dto = append(students_dto, DTO.StudentDTO{i.Roll, i.FirstName, i.LastName, i.Age, i.Marks})
	}

	sort.SliceStable(students_dto, func(i, j int) bool {
		if students_dto[i].FistName != students_dto[j].FistName {
			return students_dto[i].FistName < students_dto[j].FistName
		}
		return students_dto[i].FistName < students_dto[j].FistName
	})

	return students_dto
}

func (c *CollegeRepositoryImpl) SortByStudentMarks() []DTO.StudentDTO {
	students := []model.Student{}
	c.Db.Find(&students)

	students_dto := []DTO.StudentDTO{}

	for _, i := range students {
		students_dto = append(students_dto, DTO.StudentDTO{i.Roll, i.FirstName, i.LastName, i.Age, i.Marks})
	}

	sort.SliceStable(students_dto, func(i, j int) bool {
		if students_dto[i].Marks != students_dto[j].Marks {
			return students_dto[i].Marks > students_dto[j].Marks
		}
		return students_dto[i].FistName < students_dto[j].FistName
	})

	return students_dto
}

func (c *CollegeRepositoryImpl) SearchByStudentByCity(city string) []DTO.StudentDTO {
	students := []model.Student{}
	query := "city= '" + city + "'"
	c.Db.Find(&students, query)
	student_dto := []DTO.StudentDTO{}

	for _, i := range students {
		student_dto = append(student_dto, DTO.StudentDTO{i.Roll, i.FirstName, i.LastName, i.Age, i.Marks})
	}
	return student_dto
}

func (c *CollegeRepositoryImpl) RegisterStudent(student model.Student) error {
	grade := GradeGenerate_MySQL(student.Marks)
	student.Grade = model.Grade(grade)
	err := c.Db.Save(&student)
	if err != nil {
		return nil
	} else {
		return err.Error
	}
}

func (c *CollegeRepositoryImpl) NameSearchStartWith(name string) []DTO.StudentDTO {
	students := []model.Student{}
	query := "first_name like '" + name + "%'"
	c.Db.Find(&students, query)

	students_dto := []DTO.StudentDTO{}

	for _, i := range students {
		students_dto = append(students_dto, DTO.StudentDTO{i.Roll, i.FirstName, i.LastName, i.Age, i.Marks})
	}
	return students_dto
}

func (c *CollegeRepositoryImpl) StudentsCheck(students []model.Student) []string {
	result := []string{}
	for _, j := range students {
		student := model.Student{}
		c.Db.Find(&student, j.Roll, j.FirstName)
		if student.Marks == 0 {
			result = append(result, j.FirstName+" : Not found")
		} else {
			result = append(result, j.FirstName+" : Verified")
		}
	}
	return result
}
