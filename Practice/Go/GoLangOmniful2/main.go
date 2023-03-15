package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	user1 := User{"Rahul", 20}
	fmt.Println(user1)

	person_enc, err := json.Marshal(user1)

	if err != nil {
		println(err)
	}

	fmt.Println(string(person_enc))

	users := []User{
		{"Rahul", 30},
		{"Rony", 60},
		{"Doglesh", 27},
	}

	user_exe, err1 := json.Marshal(users)

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(user_exe))

	var humans []User

	user2 := []byte(`  [{"Name":"Rahul","Age":30},
	{"Name":"Rony","Age":60},
	{"Name":"Doglesh","Age":27}  ]`)

	datas := json.Unmarshal(user2, &humans)

	println("Started ")

	if datas != nil {
		fmt.Println(datas)
	}

	for i := range humans {
		fmt.Println(humans[i])
	}

}

type User struct {
	Name string
	Age  int
}
