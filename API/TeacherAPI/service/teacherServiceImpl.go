package service

import (
	"TeacherAPI/DTO"
	"TeacherAPI/model"
	"TeacherAPI/repository"
)

type TeacherServiceImpl struct {
	TeacherRepository repository.TeacherRepository
}

func NewTeacherServiceImpl(teacher repository.TeacherRepository) TeacherService {
	return &TeacherServiceImpl{TeacherRepository: teacher}
}

func (t *TeacherServiceImpl) GetAll() []model.Teacher {
	return t.TeacherRepository.GetAll()
}

func (t *TeacherServiceImpl) Save(teacher model.Teacher) {
	t.TeacherRepository.Save(teacher)
}

func (t *TeacherServiceImpl) SearchById(Id int) model.Teacher {
	return t.TeacherRepository.SearchById(Id)
}

func (t *TeacherServiceImpl) SearchByName(courseName string) []DTO.Teacher {
	return t.TeacherRepository.SearchByCourseName(courseName)
}

func (t *TeacherServiceImpl) Update(teacher model.Teacher) {
	t.TeacherRepository.Update(teacher)
}

func (t *TeacherServiceImpl) Delete(teacher model.Teacher) {
	t.TeacherRepository.Delete(teacher)
}

func (t *TeacherServiceImpl) NameStartWith(nameStart string) []model.Teacher {
	return t.TeacherRepository.NameStartWith(nameStart)
}
