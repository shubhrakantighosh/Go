package DTO

import "CollegeAPI/model"

type StudentDTO struct {
	Roll     int
	FistName string
	LastName string
	Age      int
	Marks    int
}

type Student_Details_DTO struct {
	Roll     int
	FistName string
	LastName string
	Age      int
	Grade    model.Grade
	PinCode  string
	City     string
	State    string
}
