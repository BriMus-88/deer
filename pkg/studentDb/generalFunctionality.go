package studentDb

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)




type Student struct {
	ID     int
	Name   string
	Course string
	Age    int
	City   string
}




func AllStudentsList(dbConnection *sql.DB) {
	clearTerminal()
	fmt.Println("\nALL STUDENTS RECORDS:")
	
	rows, err := dbConnection.Query(`
		SELECT id, name, course, age, city
		FROM students
	`)
	if err != nil {
		log.Fatal(err)
	}

	var temp Student // A temporary struct for the current row

	for rows.Next() {
		err = rows.Scan(&temp.ID, &temp.Name, &temp.Course, &temp.Age, &temp.City)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Name: %s, Course: %s, Age: %d, City: %s\n", temp.Name, temp.Course, temp.Age, temp.City)
	}

	fmt.Print("\n Press enter key to return to main menu...")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		MainMenuDisplay(dbConnection)
	}

	MainMenuDisplay(dbConnection)
}



func AddStudentMenu(dbConnection *sql.DB) {

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

		dbConnection.QueryRow(`INSERT INTO students (name, course, age, city) VALUES ($1, $2, $3, $4)`, name, course, age, city)

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
