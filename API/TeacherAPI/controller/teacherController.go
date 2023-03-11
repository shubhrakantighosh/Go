package controller

import (
	"TeacherAPI/model"
	"TeacherAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TeacherController struct {
	t service.TeacherService
}

func NewTeacherController(t service.TeacherService) *TeacherController {
	return &TeacherController{t: t}
}

func (t *TeacherController) FindAll(c *gin.Context) {
	teachers := t.t.GetAll()
	if len(teachers) == 0 {
		c.IndentedJSON(http.StatusOK, "No teachers exists")
		return
	}
	c.IndentedJSON(http.StatusOK, teachers)
}

func (t *TeacherController) SearchById(c *gin.Context) {
	Id := c.Param("Id")
	i, err := strconv.Atoi(Id)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "Please pass number value")
		return
	}
	teacher := t.t.SearchById(i)
	if teacher.Id == 0 {
		c.IndentedJSON(http.StatusOK, "Not Found")
		return
	} else {
		c.IndentedJSON(http.StatusOK, teacher)
	}
}

func (t *TeacherController) FindTeacherByCourseName(c *gin.Context) {
	courseName := c.Param("course")
	teachers := t.t.SearchByName(courseName)
	c.IndentedJSON(http.StatusOK, teachers)
}

func (t *TeacherController) Register(c *gin.Context) {
	var teacher model.Teacher
	c.ShouldBindJSON(&teacher)
	t.t.Save(teacher)
	c.IndentedJSON(http.StatusOK, "Registered Successfully")
}

func (t *TeacherController) Delete(c *gin.Context) {
	i := c.Param("Id")
	Id, err := strconv.Atoi(i)
	if err != nil {
		c.IndentedJSON(http.StatusOK, "Please pass number value")
		return
	}
	teacher := t.t.SearchById(Id)
	if teacher.Id == 0 {
		c.IndentedJSON(http.StatusOK, "Not Found")
		return
	}
	t.t.Delete(teacher)
	c.IndentedJSON(http.StatusOK, teacher)
}

func (t *TeacherController) Update(c *gin.Context) {
	teacher := model.Teacher{}
	c.ShouldBindJSON(&teacher)

	teacher_id := t.t.SearchById(teacher.Id)
	if teacher_id.Id == 0 {
		c.IndentedJSON(http.StatusOK, "No Found")
		return
	}
	t.t.Update(teacher)
	c.IndentedJSON(http.StatusOK, teacher)
}

func (t *TeacherController) NameStartWith(c *gin.Context) {
	query := c.Param("start")
	teachers := t.t.NameStartWith(query)
	if len(teachers) == 0 {
		c.IndentedJSON(http.StatusOK, "No Found")
		return
	}
	c.IndentedJSON(http.StatusOK, teachers)
}
