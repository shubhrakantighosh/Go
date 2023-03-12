package controller

import (
	"CollegeAPI/model"
	"CollegeAPI/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CollegeController struct {
	service service.CollegeService
}

func NewCollegeController(service service.CollegeService) *CollegeController {
	return &CollegeController{service: service}
}

func (college *CollegeController) FindAllStudents(c *gin.Context) {
	students := college.service.FindAllStudents()
	c.IndentedJSON(http.StatusOK, students)
}

func (college *CollegeController) SearchByStudentRoll(c *gin.Context) {
	id := c.Param("roll")
	roll, err := strconv.Atoi(id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, "No")
		return
	}
	student := college.service.SearchByStudentRoll(roll)
	c.IndentedJSON(http.StatusOK, student)
}

func (college *CollegeController) SearchByStudentName(c *gin.Context) {
	name := c.Param("name")
	students := college.service.SearchByStudentName(name)
	c.IndentedJSON(http.StatusOK, students)
}

func (college *CollegeController) SortByStudentName(c *gin.Context) {
	students := college.service.SortByStudentName()
	c.IndentedJSON(http.StatusOK, students)
}

func (college *CollegeController) SortByStudentMarks(c *gin.Context) {
	students := college.service.SortByStudentMarks()
	c.IndentedJSON(http.StatusOK, students)
}

func (college *CollegeController) SearchByStudentByCity(c *gin.Context) {
	city := c.Param("city")
	students := college.service.SearchByStudentByCity(city)
	c.IndentedJSON(http.StatusOK, students)
}

func (college *CollegeController) RegisterStudent(c *gin.Context) {
	student := model.Student{}
	c.ShouldBindJSON(&student)
	student.Grade = model.A
	err := college.service.RegisterStudent(student)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	c.IndentedJSON(http.StatusOK, student)
}

func (college *CollegeController) NameSearchStartWith(c *gin.Context) {
	name := c.Param("name")
	students := college.service.NameSearchStartWith(name)
	if len(students) == 0 {
		c.IndentedJSON(http.StatusBadRequest, "No students exists.")
		return
	}
	c.IndentedJSON(http.StatusOK, students)
}

func (college *CollegeController) StudentsCheck(c *gin.Context) {
	student := []model.Student{}
	c.ShouldBindJSON(&student)
	results := college.service.StudentsCheck(student)
	c.IndentedJSON(http.StatusOK, results)
}
