package students

import (
	"fmt"
)

var temp Student

func AddStudent() {

	var name string
	var course string
	var age int
	var city string

	for {

		fmt.Print("Enter student name:")
		fmt.Scan(&name)
		fmt.Print("Enter student course:")
		fmt.Scan(&course)
		fmt.Print("Enter student age:")
		fmt.Scan(&age)
		fmt.Print("Enter student city:")
		fmt.Scan(&city)

		temp = Student{
			Name:   name,
			Course: course,
			Age:    age,
			City:   city,
		}

		Db.QueryRow(`INSERT INTO students (name, course, age, city) VALUES ($1, $2, $3, $4)`, temp.Name, temp.Course, temp.Age, temp.City)

		fmt.Println("Student added successfully")

		fmt.Println("Do you want to add another student? (y/n)")
		var choice string
		fmt.Scan(&choice)
		if choice != "y" {
			return
		}

	}
}
