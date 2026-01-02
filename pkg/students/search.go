package students

import (
		"fmt"
		"os"
		"bufio"
		"log"
)

	
func SearchStudentByName(){

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name to search : ")
	var name string
	if scanner.Scan() {
		name = scanner.Text()
	}

	rows, err := Db.Query(`
			SELECT id, name, course, age, city
			FROM students
			WHERE name = $1
		`, name)
		if err != nil {
			log.Fatal(err)
		}

		for rows.Next() {
			var temp Student // A temporary struct for the current row
			err = rows.Scan(&temp.ID, &temp.Name, &temp.Course, &temp.Age, &temp.City)
			if err != nil {
				fmt.Println("Error scanning row:", err)
			}
			// Add this student to our list
			studentsList = append(studentsList, temp)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

	}


func SearchByAge() {
	
}

func SearchByCourse() {
	
}
