package struct_pack

type person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *person {
	person := new(person)
	person.name = name
	person.age = age
	return person
}

func (person *person) SetPersonName(name string) {
	person.name = name
}

func (person *person) GetPersonName() string {
	return person.name
}

func (person *person) SetPersonAge(age int) {
	person.age = age
}

func (person *person) GetPersonAge() int {
	return person.age
}
