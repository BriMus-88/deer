package main

import (
	"fmt"
	"log"
	// "os"
	// "encoding/json"
	// "sort"
	"github.com/brian/deer/pkg/students"
)



func main() {

 allStudents, err := students.LoadData()
    if err != nil {
        log.Fatal(err)
    }


	count, err := students.CountStudents()
    if err != nil {
        log.Fatal(err)
    }

	fmt.Println("Number of students:", count)

//	students.DisplayIntro()

	students.MainMenuLogic()


	// myMap := make(map[string]student)

	// inputfile, err := os.Open("data.json")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer inputfile.Close()

	// json.NewDecoder(inputfile).Decode(&myMap)	
	

	// for k, v := range myMap {
	// 	fmt.Printf("Name: %s, Course: %s, Age: %d, City: %s\n", k, v.course, v.age, v.city)
	// }



	// if _, ok := myMap["Dan"]; ok {
	// 	delete(myMap, "Dan")
	// } else {
	// 	fmt.Println("Dan is not in the map")
	// }

	// for k, v := range myMap {
	// 	fmt.Printf("Name: %s, Course: %s, Age: %d, City: %s\n", k, v.course, v.age, v.city)
	// }

	// if _, ok := myMap["Dan"]; ok {
	// 	delete(myMap, "Dan")
	// } else {
	// 	fmt.Println("Dan is not in the map")
	// }


	// var keys []string
	// for k := range myMap {
	// 	keys = append(keys, k)
	// }

	// sort.Strings(keys)


	// fmt.Println("\nSorted map:")
	// for _, k := range keys {
	// 	fmt.Printf("Name: %s, Course: %s, Age: %d, City: %s\n", k, myMap[k].course, myMap[k].age, myMap[k].city)
	// }


	// 	var name string

	// fmt.Println("\n Enter students name :")

	// fmt.Scan(&name)

	// if _, ok := myMap[name]; ok {
	// 	fmt.Println("Student is in the map")
	// 	fmt.Printf("Name: %s, Course: %s, Age: %d, City: %s\n", name, myMap[name].course, myMap[name].age, myMap[name].city)
	// } else {
	// 	fmt.Println("Student is not in the map")
	// }

	fmt.Println(allStudents)
	students.Db.Close() 

}
