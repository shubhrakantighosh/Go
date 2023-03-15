package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//const portNumber = ":8080"
//
//func homeHandler(w http.ResponseWriter, r *http.Request) {
//	bio := `I'm Shubhra`
//	w.Header().Set("Content-Type", "text/html; charset=utf-8 ")
//	fmt.Fprint(w, "<h1>"+bio+"</h1>")
//}

type student struct {
	Roll      int
	FirstName string
	LastName  string
	Age       int
}

type address struct {
	Id      int
	PinCode string
	City    string
	State   string
}

// `gorm:"primary_key;auto_increment;"`
// `gorm:"embedded"`
func main() {

	//students := []student{
	//	student{Roll: 11, FirstName: "Shubhra", LastName: "Ghosh", Age: 26},
	//	student{Roll: 12, FirstName: "Chirag", LastName: "Gupta", Age: 26},
	//}
	//
	//address := []address{
	//	address{PinCode: "700102", City: "Kolkata", State: "WB"},
	//	address{PinCode: "713246", City: "Durgapur", State: "WB"},
	//}

	dsn := "root:shubhrakanti@tcp(localhost:3306)/db2"

	db, err := gorm.Open("mysql", dsn)
	db.LogMode(true)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	db.DropTableIfExists(&student{}, &address{})
	db.CreateTable(&student{}, &address{})

	//db.CreateTable(&Person{})
	//
	//db.Create(&Person{
	//	student: student{11, "Shubhra", "Ghosh", 26},
	//	address: []address{
	//		{PinCode: "7", City: "8", State: "1"},
	//		{PinCode: "1", City: "8", State: "1"},
	//	},
	//})

	//db.AutoMigrate(&student{}, &address{})

	db.Create(&student{FirstName: "Shubhra", LastName: "Ghosh", Age: 26})
	db.Create(&student{FirstName: "Chirag", LastName: "Gupta", Age: 26})

	//router := chi.NewRouter()
	//router.Get("/", homeHandler)
	//router.NotFound(func(w http.ResponseWriter, request *http.Request) {
	//	http.Error(w, "Page Not Found ", http.StatusNotFound)
	//})
	//http.ListenAndServe(portNumber, router)

	//t, err := template.ParseFiles("goland.html")
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//user := User{"Shubhra", 12, "shubhra@gmail.com"}
	//
	//_ = t.Execute(os.Stdout, user)

	//mysql, err := sql.Open("mysql", "root:shubhrakanti@tcp(localhost:3306)/db2")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//defer mysql.Close()

	//
	//row, err := mysql.Query("select * from user where Id=?", 12)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//defer row.Close()
	//
	//users := []User{}
	//
	//for row.Next() {
	//
	//	user := User{}
	//	err := row.Scan(&user.Id, &user.Name, &user.age, &user.Address)
	//
	//	if err != nil {
	//		fmt.Println(err)
	//	}
	//
	//	users = append(users, user)
	//
	//}
	//
	//fmt.Println(users)

	//result, err := mysql.Exec("insert into user values (19,'Roshan Gupta',26,'Delhi')")
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(result.LastInsertId())
	//fmt.Println(result.RowsAffected())

	//result1, err := mysql.Exec("delete from user where id=?", 18)
	//
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//fmt.Println(result1.LastInsertId())
	//fmt.Println(result1.RowsAffected())

}

//type User struct {
//	Id      int
//	Name    string
//	age     int
//	Address string
//}
