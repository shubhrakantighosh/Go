package repository

import (
	"TeacherAPI/DTO"
	"TeacherAPI/model"
	"fmt"
	"gorm.io/gorm"
)

type TeacherRepositoryImpl struct {
	Db *gorm.DB
}

func NewTeacherRepositoryImpl(Db *gorm.DB) TeacherRepository {
	return &TeacherRepositoryImpl{Db: Db}
}

func (t *TeacherRepositoryImpl) GetAll() []model.Teacher {
	teachers := []model.Teacher{}
	t.Db.Find(&teachers)
	return teachers
}

func (t *TeacherRepositoryImpl) Save(teacher model.Teacher) {
	t.Db.Save(&teacher)
}

func (t *TeacherRepositoryImpl) SearchById(Id int) model.Teacher {
	teacher := model.Teacher{}
	t.Db.Find(&teacher, Id)
	fmt.Print(teacher)
	return teacher
}

func (t *TeacherRepositoryImpl) SearchByCourseName(courseName string) []DTO.Teacher {
	teachers := []DTO.Teacher{}
	for _, i := range t.GetAll() {
		if i.Department == courseName {
			teacher := DTO.Teacher{Name: i.Name, Age: i.Age, Department: i.Department}
			teachers = append(teachers, teacher)
		}
	}
	return teachers
}

func (t *TeacherRepositoryImpl) Update(teacher model.Teacher) {
	t.Db.Updates(&teacher)
}

func (t *TeacherRepositoryImpl) Delete(teacher model.Teacher) {
	err := t.Db.Delete(teacher)
	fmt.Print(err)
}

func (t *TeacherRepositoryImpl) NameStartWith(nameStart string) []model.Teacher {
	teachers := []model.Teacher{}
	query := "name like" + "'" + nameStart + "%'"
	t.Db.Find(&teachers, query)
	return teachers
}
