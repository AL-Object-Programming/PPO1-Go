package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	names := []string{"Andrew", "William", "James", "Harper", "Mason", "Evelyn", "Ella", "Avery"}
	lastNames := []string{"Smith", "Jones", "Williams", "Brown"}
	var students []Student

	for index := 0; index < 20; index++ {

		students = append(students, Student{
			name:     names[random.Intn(len(names))],
			lastName: lastNames[random.Intn(len(lastNames))],
			index: random.Intn(999) + 39000,
			status:   true,
		})

		for _, student := range students {
			fmt.Println(student.lastName + " " + student.name + " (" + strconv.Itoa(student.index) + ") ")
		}
	}
}
// go run main.go student.go