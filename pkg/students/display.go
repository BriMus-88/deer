package students

import (
	"fmt"
	"log"
	"os"
)

type User struct {
	userName string
	password string
	fullName string
	role     string
}

var validUsers = make(map[int]User)

func init() {
	validUsers[1] = User{userName: "admin", password: "admin", fullName: "Admin User", role: "admin"}
	validUsers[2] = User{userName: "user", password: "user", fullName: "User User", role: "user"}
	validUsers[3] = User{userName: "guest", password: "guest", fullName: "Guest User", role: "guest"}
}

func MainMenuLogic(allStudents []Student) {

	var choice int
	for {
		clearTerminal()
		fmt.Print("MAIN MENU:")
		fmt.Print("\n")
		fmt.Println("1. Update Student Database")
		fmt.Println("2. Empty")
		fmt.Println("3. Display all students")
		fmt.Println("4. Search for a student Menu")
		fmt.Println("5. Export students to file")
		fmt.Println("6. Import students from file")
		fmt.Println("7. Group students by course")
		fmt.Println("8. Switch modes")
		fmt.Println("9. Delete all students")
		fmt.Println("0. Exit")

		count, err := CountStudents()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("Number of students: ", count)
		fmt.Print("\nEnter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 0:
			fmt.Println("Bye Bye")
			os.Exit(0)
		case 3:
			DisplayAllStudents(allStudents)
		case 5:
			fmt.Println("Export All students to file")

		case 6:
			fmt.Println("Import students from file")

		case 7:
			fmt.Println("Group students by course")

		case 8:
			fmt.Println("Switch modes")

		case 9:
			fmt.Println("Delete all students")
		default:
			fmt.Println("Invalid choice")
		}

	}
}

func DisplaySearchMenu(allStudents []Student) {

	for {
		clearTerminal()
		fmt.Println("\nSEARCH FOR STUDENT BY:")
		fmt.Println("\t 1. Name")
		fmt.Println("\t 2. Age")
		fmt.Println("\t 3. Course")
		fmt.Println("\t 0. Return to Main Menu")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 0:
			MainMenuLogic(allStudents)
		case 1:
			fmt.Println("Search for a student by name")
		case 2:
			fmt.Println("Search for a student by age")
		case 3:
			fmt.Println("Search for a student by course")
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func DisplayEditMenu(allStudents []Student) {

	clearTerminal()
	fmt.Println("\nEDIT STUDENT DATABASE:")
	fmt.Println("\t 1. Add a student")
	fmt.Println("\t 2. Remove a student")
	fmt.Println("\t 0. Return to Main Menu")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 0:
		MainMenuLogic(allStudents)
	case 1:
		AddStudent()
	case 2:
		fmt.Println("Remove a student")
	default:
		fmt.Println("Invalid choice")
	}
}

func DisplayIntro() {
	clearTerminal()
	var userName, password string

	for {
		fmt.Println("\nWelcome to the student management database (SAD)")
		fmt.Print("\n")
		fmt.Print("Enter User Name : ")
		fmt.Scan(&userName)
		fmt.Print("Enter Password : ")
		fmt.Scan(&password)
		fmt.Print("\n")

		for k := range validUsers {
			if validUsers[k].userName == userName {
				if validUsers[k].password == password {
					fmt.Println("Welcome", validUsers[k].fullName)
					fmt.Println("Its running on file mode")

					return
				}
			}

		}
		fmt.Println("Invalid User Name or Password")

	}
}

func clearTerminal() {
	fmt.Println("\033[H")
	fmt.Println("\033[2J")
}


func DisplayAllStudents(allStudents []Student) {
	clearTerminal()
	fmt.Println("\nALL STUDENTS RECORDS:")
	
	rows, err := Db.Query(`
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

	var choice int

	fmt.Print("\n Select any key to return to main menu...")
	fmt.Scan(&choice)

	MainMenuLogic(allStudents)
}