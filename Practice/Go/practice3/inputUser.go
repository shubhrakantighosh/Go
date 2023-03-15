package main

import "fmt"

func InputUser() map[Student]Address {

	students := make(map[Student]Address)

	var firstName, lastName, email, pinCode, city, state string
	var roll, age, length int

	fmt.Println("How many students you want to add")
	fmt.Scanf("%d", &length)

	for i := 0; i < length; i++ {

		fmt.Println("Enter Your Roll ")
		fmt.Scanf("%d", &roll)
		fmt.Println("Enter Your First Name ")
		fmt.Scanf("%s", &firstName)
		fmt.Println("Enter Your Last Name ")
		fmt.Scanf("%s", &lastName)
		fmt.Println("Enter Your Age ")
		fmt.Scanf("%d", &age)
		fmt.Println("Enter Your Email ")
		fmt.Scanf("%s", &email)

		fmt.Println("Enter Your Pin Code ")
		fmt.Scanf("%s", &pinCode)
		fmt.Println("Enter Your City ")
		fmt.Scanf("%s", &city)
		fmt.Println("Enter State ")
		fmt.Scanf("%s", &state)

		student := Student{roll, firstName, lastName, age, email}
		address := Address{pinCode, city, state}

		students[student] = address

	}

	return students
}
