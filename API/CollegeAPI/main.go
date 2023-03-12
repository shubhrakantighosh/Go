package main

import (
	"CollegeAPI/config"
	"CollegeAPI/controller"
	"CollegeAPI/model"
	"CollegeAPI/repository"
	"CollegeAPI/router"
	"CollegeAPI/service"
	"net/http"
)

func main() {

	Db := config.DataBaseConnection()
	Db.AutoMigrate(&model.Student{})

	collegeRepository := repository.NewCollegeRepositoryImpl(Db)
	collegeService := service.NewCollegeServiceImpl(collegeRepository)
	collegeController := controller.NewCollegeController(collegeService)

	router := router.Router(collegeController)
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.ListenAndServe()

}
