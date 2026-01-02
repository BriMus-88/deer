package studentDb

import (
	"database/sql"
	"fmt"
	// "log"
	// "os"
	// "bufio"
)

func CountStudents(dbConnection *sql.DB) int {

	var count int

	err := dbConnection.QueryRow(`SELECT COUNT(*) FROM students`).Scan(&count)
	if err != nil {
		fmt.Println("Error counting students:", err)
		return 0
	}

	return count
}

