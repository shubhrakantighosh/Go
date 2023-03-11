package router

import (
	"TeacherAPI/controller"
	"github.com/gin-gonic/gin"
)

func Router(c *controller.TeacherController) *gin.Engine {

	router := gin.Default()
	teacherRouter := router.Group("/teacher")
	teacherRouter.GET("/all", c.FindAll)
	teacherRouter.GET("/:Id", c.SearchById)
	teacherRouter.POST("/register", c.Register)
	teacherRouter.GET("/course/:course", c.FindTeacherByCourseName)
	teacherRouter.DELETE("/delete/:Id", c.Delete)
	teacherRouter.PUT("", c.Update)
	teacherRouter.GET("/name/:start", c.NameStartWith)

	return router
}
