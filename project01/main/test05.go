package main

import (
	"fmt"
)
//结构体
type Student struct{
	name string
	sex byte
	number int
	score int
}


func main() {
	var student1 Student
	var student2 Student

	student1.name="xindewin"
	student1.sex='f'
	student2.score=12
	student2.number=12204202
    printstudent(student1)
	printstudent(student2)

}

func printstudent(student Student) {
	fmt.Printf("name:%s\n",student.name)
	fmt.Printf("sex:%q\n",student.sex)
	fmt.Printf("score:%d\n",student.score)
	fmt.Printf("number:%d\n",student.number)
}
