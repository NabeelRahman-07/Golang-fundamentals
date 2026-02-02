package main

import "fmt"

type Student struct {
	Id   int
	Name string
	Age  int
}

var idCounter = 1
var students []Student

func AddStudent(name string, age int) {
	st := Student{
		Id:   idCounter,
		Name: name,
		Age:  age,
	}
	students = append(students, st)
	fmt.Println("New student added")
	idCounter++
}

func EditStudent(id int, name string, age int) {
	for i := range students {
		if students[i].Id == id {
			students[i].Name = name
			students[i].Age = age
			fmt.Println("Student updated")
			return
		}
	}
	fmt.Println("Student not found")
}
func Deletestudent(id int) {
	for i := range students {
		if students[i].Id == id {
			students = append(students[:i], students[i+1:]...)
			fmt.Println("Student deleted")
			return
		}
	}
	fmt.Println("Student not found")
}
func ListStudent() {
	if len(students) == 0 {
		fmt.Println("No students found")
		return
	}
	for _, v := range students {
		fmt.Printf("Id: %d | Name: %s | Age: %d \n", v.Id, v.Name, v.Age)
	}
}

func main() {
	for {
		fmt.Println("\n 1.Add student")
		fmt.Println("2.View students")
		fmt.Println("3.Update student")
		fmt.Println("4.Delete student")
		fmt.Println("5.Exit")

		var choice int
		fmt.Println("Enter your choice:")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var age int
			fmt.Println("Enter student name:")
			fmt.Scan(&name)
			fmt.Println("Enter student age:")
			fmt.Scan(&age)
			AddStudent(name, age)

		case 2:
			ListStudent()

		case 3:
			var id int
			var newname string
			var newage int
			fmt.Println("Enter student Id:")
			fmt.Scan(&id)
			fmt.Println("Enter new name:")
			fmt.Scan(&newname)
			fmt.Println("Enter new age")
			fmt.Scan(&newage)
			EditStudent(id, newname, newage)

		case 4:
			var id int
			fmt.Println("Enter student Id:")
			fmt.Scan(&id)
			Deletestudent(id)

		case 5:
			fmt.Println("Good bye")
			return

		default:
			fmt.Println("Invalid option!")
		}
	}
}
