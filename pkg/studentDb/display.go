package studentDb

import (
	"bufio"
	"fmt"

	"log"
	// 	"time"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func (s *StudentStore) MainMenuDisplay() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	for {
		clearTerminal()
		fmt.Print("MAIN MENU:")
		fmt.Print("\n\n")
		fmt.Println("\t1. Manually Updating Student Records")
		fmt.Println("\t2. Empty")
		fmt.Println("\t3. Display all students")
		fmt.Println("\t4. Search for a student Menu")
		fmt.Println("\t5. Export students to file")
		fmt.Println("\t6. Import students from file")
		fmt.Println("\t7. Add student record Manually")
		fmt.Println("\t8. Delete student record Manually")
		fmt.Println("\t9. Edit student record Manually")
		fmt.Println("\t0. Exit")

		count := s.CountStudents()

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
			s.UpdateDbMenu()
		case 3:
			s.AllStudentsList()
		case 4:
			s.SearchMenuDisplay()
		case 5:
			//			fmt.Println("Export All students to file")

		case 6:
			//			fmt.Println("Import students from file")

		case 7:
			s.AddStudentMenu()
		case 8:
			s.DeleteStudentMenu()

		case 9:
			fmt.Println("to be edited")
		default:
			fmt.Println("Invalid choice")
		}

	}
}

func (s *StudentStore) SearchMenuDisplay() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	for {
		clearTerminal()
		fmt.Println("\nSEARCH FOR STUDENT BY:")
		fmt.Println("\t 1. Name")
		fmt.Println("\t 2. Age")
		fmt.Println("\t 3. Course")
		fmt.Println("\t 0. Return to Main Menu")

		fmt.Print("\nEnter Selection : ")

		if scanner.Scan() {
			choice, _ = strconv.Atoi(scanner.Text())
		}

		switch choice {
		case 0:
			return
		case 1:
			s.SearchStudentByName()
		case 2:
			//
		case 3:
			//
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (s *StudentStore) UpdateDbMenu() {

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
			return
		case 1:
			s.AddStudentMenu()
		case 2:
			fmt.Println("Remove a student")
			// DeleteStudentRecord()
		default:
			fmt.Println("Invalid choice")
		}
	}
}

func (s *StudentStore) SearchStudentByName() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearTerminal()
		fmt.Println("\nSEARCH BY STUDENT NAME:")
		fmt.Print("\n\t Enter Student Name : ")

		var name2Search string
		if scanner.Scan() {
			name2Search = scanner.Text()
		}

		rows, err := s.DB.Query(`
			SELECT id, name, course, age, city
			FROM students
			WHERE name ILIKE $1
		`, "%"+name2Search+"%")
		if err != nil {
			log.Fatal(err)
		}

		foundCount := 0
		var temp Student // A temporary struct for the current row
		for rows.Next() {
			err = rows.Scan(&temp.ID, &temp.Name, &temp.Course, &temp.Age, &temp.City)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("ID: %d\n", temp.ID)
			fmt.Printf("Name: %s\n", temp.Name)
			fmt.Printf("Course: %s\n", temp.Course)
			fmt.Printf("Age: %d\n", temp.Age)
			fmt.Printf("City: %s\n", temp.City)
			fmt.Println("---------------------------")
			foundCount++
		}

		if foundCount == 0 {
			fmt.Println("No results found for:", name2Search)
		} else {
			fmt.Printf("Total results found: %d\n", foundCount)
		}

		fmt.Print("\nSearch again? (y/n): ")
		if scanner.Scan() {
			choice := scanner.Text()
			if choice != "y" {
				return
			}

		}

	}
}

func (s *StudentStore) DisplayIntro() {

	type User struct {
		userName string
		password string
		fullName string
		role     string
	}

	var validUsers = make(map[int]User)

	validUsers[1] = User{userName: "admin", password: "admin", fullName: "Admin User", role: "admin"}
	validUsers[2] = User{userName: "user", password: "user", fullName: "User User", role: "user"}
	validUsers[3] = User{userName: "guest", password: "guest", fullName: "Guest User", role: "guest"}

	clearTerminal()
	var userName, password string
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nWelcome to the student management database (SAD)")
		fmt.Print("\n")
		fmt.Print("Enter User Name : ")
		scanner.Scan()
		userName = scanner.Text()
		fmt.Print("Enter Password : ")
		scanner.Scan()
		password = scanner.Text()
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

func (s *StudentStore) DeleteStudentMenu() {

	scanner := bufio.NewScanner(os.Stdin)

	for {
		clearTerminal()
		fmt.Println("\nDELETE STUDENT RECORD BY STUDENT NAME:")
		fmt.Print("\n\t Enter Student Name : ")

		var name2Search string
		if scanner.Scan() {
			name2Search = scanner.Text()
		}

		rows, err := s.DB.Query(`
			SELECT id, name, course, age, city
			FROM students
			WHERE name ILIKE $1
		`, "%"+name2Search+"%")
		if err != nil {
			log.Fatal(err)
		}

		foundCount := 0
		var temp Student // A temporary struct for the current row
		for rows.Next() {
			err = rows.Scan(&temp.ID, &temp.Name, &temp.Course, &temp.Age, &temp.City)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("ID: %d\n", temp.ID)
			fmt.Printf("Name: %s\n", temp.Name)
			fmt.Printf("Course: %s\n", temp.Course)
			fmt.Printf("Age: %d\n", temp.Age)
			fmt.Printf("City: %s\n", temp.City)
			fmt.Println("---------------------------")
			foundCount++
		}

		if foundCount == 0 {
			fmt.Println("No results found for:", name2Search)
		} else {
			fmt.Printf("Total results found: %d\n", foundCount)
		}

		id2Delete := 0
		fmt.Print("\nEnter ID of Student Record to delete: ")
		if scanner.Scan() {
			id2Delete, _ = strconv.Atoi(scanner.Text())
		}

		rows, err = s.DB.Query(`
			SELECT id, name, course, age, city
			FROM students
			WHERE id = $1
		`, id2Delete)
		if err != nil {
			log.Fatal(err)
		}

		clearTerminal()
		fmt.Println("---------------------------")
		fmt.Printf("ID: %d\n", temp.ID)
		fmt.Printf("Name: %s\n", temp.Name)
		fmt.Printf("Course: %s\n", temp.Course)
		fmt.Printf("Age: %d\n", temp.Age)
		fmt.Printf("City: %s\n", temp.City)
		fmt.Println("---------------------------")

		fmt.Print("\nConfirm this is the Student Record to delete? (y/n): ")
		if scanner.Scan() {
			choice := scanner.Text()
			if choice == "y" {
				
				_, err := s.DB.Exec(`
					DELETE FROM students
					WHERE id = $1	
				`, id2Delete)
				if err != nil {
					log.Println(err)
					fmt.Println("\nStudent Record not deleted due to error", err)
				}else{
					fmt.Println("\nStudent Record deleted")
				}
				

			} else {
				fmt.Println("\nStudent Record not deleted")
			}
		}

		fmt.Println("\n\t Please select")
		fmt.Println("\t\t 1. Search another student to delete")
		fmt.Println("\t\t 2. Return to Main Menu")
		fmt.Println("\t\t 0. Exit the program")
		for {
			if scanner.Scan() {
				choice := scanner.Text()
				switch choice {
				case "1":
					s.DeleteStudentMenu()
				case "2":
					return
				case "0":
					os.Exit(0)
				default:
					fmt.Println("Invalid choice")
				}
			}
		}

	}

}
