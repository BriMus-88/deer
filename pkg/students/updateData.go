package students

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

var temp Student

func AddStudent() {

	scanner := bufio.NewScanner(os.Stdin)

	var name string
	var course string
	var age int
	var city string

	for {

		fmt.Print("Enter student name:")
		if scanner.Scan() {
			name = scanner.Text()
		}
		fmt.Print("Enter student course:")
		if scanner.Scan() {
			course = scanner.Text()
		}
		fmt.Print("Enter student age:")
		if scanner.Scan() {
			age, _ = strconv.Atoi(scanner.Text())
		}
		fmt.Print("Enter student city:")
		if scanner.Scan() {
			city = scanner.Text()
		}

		temp = Student{
			Name:   name,
			Course: course,
			Age:    age,
			City:   city,
		}

		Db.QueryRow(`INSERT INTO students (name, course, age, city) VALUES ($1, $2, $3, $4)`, temp.Name, temp.Course, temp.Age, temp.City)

		fmt.Println("Student added successfully")

		fmt.Print("Do you want to add another student? (y/n) :")
		var choice string
		scanner.Scan()
		choice = scanner.Text()
		//fmt.Scan(&choice)
		if choice != "y" {
			return
		}

	}
}
