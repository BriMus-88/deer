package studentDb

import (
	"database/sql"
	"fmt"
	// "log"
	// "os"
	// "bufio"
)

func (s *StudentStore) CountStudents() int {

	var count int

	err := s.DB.QueryRow(`SELECT COUNT(*) FROM students`).Scan(&count)
	if err != nil {
		fmt.Println("Error counting students:", err)
		return 0
	}

	return count
}

func NewStudentStore(db *sql.DB) *StudentStore {
	return &StudentStore{
		DB: db,
	}
}

type StudentStore struct {
	DB *sql.DB
}
