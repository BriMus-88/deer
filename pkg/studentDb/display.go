package studentDb

import (
	"bufio"
	"database/sql"
	"fmt"

	// 	"log"
	// 	"time"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// type User struct {
// 	userName string
// 	password string
// 	fullName string
// 	role     string
// }

// var validUsers = make(map[int]User)

// func init() {
// 	validUsers[1] = User{userName: "admin", password: "admin", fullName: "Admin User", role: "admin"}
// 	validUsers[2] = User{userName: "user", password: "user", fullName: "User User", role: "user"}
// 	validUsers[3] = User{userName: "guest", password: "guest", fullName: "Guest User", role: "guest"}
// }

func MainMenuDisplay(dbConnection *sql.DB) {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	for {
		clearTerminal()
		fmt.Print("MAIN MENU:")
		fmt.Print("\n\n")
		fmt.Println("\t1. Update Student Records")
		fmt.Println("\t2. Empty")
		fmt.Println("\t3. Display all students")
		fmt.Println("\t4. Search for a student Menu")
		fmt.Println("\t5. Export students to file")
		fmt.Println("\t6. Import students from file")
		fmt.Println("\t7. Group students by course")
		fmt.Println("\t8. Switch modes")
		fmt.Println("\t9. Delete all students")
		fmt.Println("\t0. Exit")

		count := CountStudents(dbConnection)

		fmt.Print("\nNumber of students: ", count)

		fmt.Print("\nEnter selection: ")
		if scanner.Scan() {
			choice, _ = strconv.Atoi(scanner.Text())
		}

		switch choice {
		case 0:
			fmt.Println("Bye Bye")
			os.Exit(0)
		case 1:
			UpdateDbMenu(dbConnection)
		case 3:
			AllStudentsList(dbConnection)
		case 5:
			//			fmt.Println("Export All students to file")

		case 6:
			//			fmt.Println("Import students from file")

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

// func DisplaySearchMenu(allStudents []Student) {

// 	for {
// 		clearTerminal()
// 		fmt.Println("\nSEARCH FOR STUDENT BY:")
// 		fmt.Println("\t 1. Name")
// 		fmt.Println("\t 2. Age")
// 		fmt.Println("\t 3. Course")
// 		fmt.Println("\t 0. Return to Main Menu")

// 		var choice int
// 		fmt.Scan(&choice)

// 		switch choice {
// 		case 0:
// 			MainMenuLogic(allStudents)
// 		case 1:
// 			SearchStudentByName()
// 		case 2:
// 			fmt.Println("Search for a student by age")
// 		case 3:
// 			fmt.Println("Search for a student by course")
// 		default:
// 			fmt.Println("Invalid choice")
// 		}
// 	}
// }

func UpdateDbMenu(dbConnection *sql.DB) {
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		clearTerminal()
		fmt.Println("\nEDIT STUDENT DATABASE:")
		fmt.Println("\t 1. Add a student")
		fmt.Println("\t 2. Delete a student")
		fmt.Println("\t 0. Return to Main Menu")

		fmt.Println("\n\t Enter Selection")

		var choice int
		if scanner.Scan() {
			choice, _ = strconv.Atoi(scanner.Text())
		}

		switch choice {
		case 0:
			MainMenuDisplay(dbConnection)
		case 1:
			AddStudentMenu(dbConnection)
		case 2:
			fmt.Println("Remove a student")
			// DeleteStudentRecord()
		default:
			fmt.Println("Invalid choice")
		}
	}
}

// func DisplayIntro() {
// 	clearTerminal()
// 	var userName, password string

// 	for {
// 		fmt.Println("\nWelcome to the student management database (SAD)")
// 		fmt.Print("\n")
// 		fmt.Print("Enter User Name : ")
// 		fmt.Scan(&userName)
// 		fmt.Print("Enter Password : ")
// 		fmt.Scan(&password)
// 		fmt.Print("\n")

// 		for k := range validUsers {
// 			if validUsers[k].userName == userName {
// 				if validUsers[k].password == password {
// 					fmt.Println("Welcome", validUsers[k].fullName)
// 					fmt.Println("Its running on file mode")

// 					return
// 				}
// 			}

// 		}
// 		fmt.Println("Invalid User Name or Password")

// 	}
// }

// func clearTerminal() {
// 	fmt.Println("\033[H")
// 	fmt.Println("\033[2J")
// }

// func DisplayAllStudents(allStudents []Student) {
// 	clearTerminal()
// 	fmt.Println("\nALL STUDENTS RECORDS:")
// 	time.Sleep(2 * time.Second)
// 	rows, err := Db.Query(`
// 		SELECT id, name, course, age, city
// 		FROM students
// 	`)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var temp Student // A temporary struct for the current row

// 	for rows.Next() {
// 		err = rows.Scan(&temp.ID, &temp.Name, &temp.Course, &temp.Age, &temp.City)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		fmt.Printf("Name: %s, Course: %s, Age: %d, City: %s\n", temp.Name, temp.Course, temp.Age, temp.City)
// 	}

// 	fmt.Print("\n Press enter key to return to main menu...")
// 	scanner := bufio.NewScanner(os.Stdin)
// 	if scanner.Scan() {
// 		MainMenuLogic(allStudents)
// 	}

// 	MainMenuLogic(allStudents)
// }

func clearTerminal() {
	fmt.Println("\033[H")
	fmt.Println("\033[2J")
}
