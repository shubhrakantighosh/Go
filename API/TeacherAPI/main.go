package main

import (
	"TeacherAPI/config"
	"TeacherAPI/controller"
	"TeacherAPI/model"
	"TeacherAPI/repository"
	"TeacherAPI/router"
	"TeacherAPI/service"
	"net/http"
)

func main() {

	Db := config.DataBaseConnection()
	Db.AutoMigrate(&model.Teacher{})

	tRepo := repository.NewTeacherRepositoryImpl(Db)
	tSer := service.NewTeacherServiceImpl(tRepo)
	tCon := controller.NewTeacherController(tSer)

	router := router.Router(tCon)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}
