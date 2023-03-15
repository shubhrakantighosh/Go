package controller

import (
	"CollegeAPI/DTO"
	"CollegeAPI/model"
	"CollegeAPI/service"
	"github.com/gin-gonic/gin"
	"gopkg.in/validator.v2"
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
	students_details := []DTO.Student_Details_DTO{}
	for _, j := range students {
		students_details = append(students_details, DTO.Student_Details_DTO{j.Roll, j.FistName, j.LastName, j.Age, "A", j.PinCode, j.City, j.State})
	}
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
	if err := c.ShouldBindJSON(&student); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := validator.Validate(student); err != nil {
		c.IndentedJSON(http.StatusBadGateway, "Wrong Grade")
		return
	}
	if error := college.service.RegisterStudent(student); error != nil {
		c.IndentedJSON(http.StatusBadRequest, error.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "Registered successfully.")
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

func (college *CollegeController) RegisterTeacher(c *gin.Context) {
	teacher := model.Teacher{}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	if err := college.service.RegisterTeacher(teacher); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, "Registered successfully.")
}

func (college *CollegeController) ValidationData(c *gin.Context) {
	teacher := model.Teacher{}
	if err := c.ShouldBindJSON(&teacher); err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error())
		return
	}

	result := college.service.ValidationData(teacher)

	if result == nil {
		c.IndentedJSON(http.StatusBadRequest, "Not Found")
		return
	}

	c.IndentedJSON(http.StatusOK, result)
}
