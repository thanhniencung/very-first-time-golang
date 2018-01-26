package main 

import (
	"app/model"
	"fmt"
)

func main() {
	model.InitDB();

	students, _ := model.FindAll()
	
	for i := range(students) {
        student := students[i]
        fmt.Println(student.FirstName)
    }
}