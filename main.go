package main

import "fmt"

type Student struct {
	id    uint64
	name  string
	class string
}

var StudentList []Student

func main() {
	ListMenu()
}

func ListMenu() {
	for {
		var menuOption int

		fmt.Println("===")
		fmt.Println("1.Get All Data")
		fmt.Println("2.Create Data")
		fmt.Println("3.Update Data")
		fmt.Println("4.Delete Data")
		fmt.Println("5.Exit")
		fmt.Println("===")

		fmt.Print("select menu : ")
		fmt.Scan(&menuOption)

		switch menuOption {
		case 1:
			GetDataStudent()

		case 2:
			if CreateDataStudent() {
				fmt.Println("Success Create Data")
			} else {
				fmt.Println("Failed Create Data")
			}
		case 3:
			if UpdateDataStudent() {
				fmt.Println("Success Edit Data")
			} else {
				fmt.Println("Failed Edit Data")
			}
			break
		case 4:
			if DeleteDataStudent() {
				fmt.Println("Success Delete Data")
			} else {
				fmt.Println("Failed Delete Data")
			}

		}

		if menuOption == 5 {
			break
		}
	}
}

func GetDataStudent() {
	for _, studentList := range StudentList {
		fmt.Println("===")
		fmt.Println("id : ", studentList.id)
		fmt.Println("name : ", studentList.name)
		fmt.Println("class : ", studentList.class)
		fmt.Println("===")
	}
}

func CreateDataStudent() bool {
	var name string
	var class string

	fmt.Print("name : ")
	fmt.Scan(&name)
	fmt.Print("class : ")
	fmt.Scan(&class)

	createSutdent := Student{
		id:    uint64(len(StudentList)) + 1,
		name:  name,
		class: class,
	}

	StudentList = append(StudentList, createSutdent)

	return true
}

func UpdateDataStudent() bool {
	var student Student
	fmt.Print("id : ")
	fmt.Scan(&student.id)
	fmt.Print("name : ")
	fmt.Scan(&student.name)
	fmt.Print("class : ")
	fmt.Scan(&student.class)

	for i, studentList := range StudentList {
		if studentList.id == student.id {

			StudentList[i] = student

			return true

		}
	}

	return false

}

func DeleteDataStudent() bool {
	var id uint64
	fmt.Print("id : ")
	fmt.Scan(&id)

	for i, studentList := range StudentList {
		if studentList.id == id {
			StudentList = append(StudentList[:i], StudentList[i+1:]...)
			return true
		}
	}
	return false
}
