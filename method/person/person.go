package person

type Person struct {
	FirstName string
	LastName  string
	Age       int
	Address   string
}

func (person *Person) SetFirstName(firstName string) {
	person.FirstName = firstName
}

func (person *Person) SetLastName(lastName string) {
	person.LastName = lastName
}
func (person *Person) SetAge(age int) {
	person.Age = age
}
func (person *Person) SetAddress(address string) {
	person.Address = address
}

func (person *Person) SetAll(firstName string, lastName string, age int, address string) {
	person.FirstName = firstName
	person.FirstName = lastName
	person.Age = age
	person.Address = address
}

func (person Person) GetFirstName() string {
	return person.FirstName
}
func (person Person) GetLastName() string {
	return person.LastName
}
func (person Person) GetAge() int {
	return person.Age
}
func (person Person) GetAddress() string {
	return person.Address
}

func (person Person) GetAll() Person {
	return Person{person.FirstName, person.LastName, person.Age, person.Address}
}
