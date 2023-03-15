package router

import (
	"CollegeAPI/controller"
	"github.com/gin-gonic/gin"
)

func Router(controller *controller.CollegeController) *gin.Engine {

	router := gin.Default()

	collegeRouter := router.Group("/college")
	collegeRouter.GET("/students", controller.FindAllStudents)
	collegeRouter.GET("/roll/:roll", controller.SearchByStudentRoll)
	collegeRouter.GET("/name/:name", controller.SearchByStudentName)
	collegeRouter.GET("/sort/name", controller.SortByStudentName)
	collegeRouter.GET("/sort/marks", controller.SortByStudentMarks)
	collegeRouter.GET("/city/:city", controller.SearchByStudentByCity)
	collegeRouter.POST("/student", controller.RegisterStudent)
	collegeRouter.GET("/start/:name", controller.NameSearchStartWith)
	collegeRouter.POST("/test", controller.StudentsCheck)
	collegeRouter.POST("/teacher", controller.RegisterTeacher)
	collegeRouter.POST("/teacher/Id", controller.ValidationData)

	return router
}
