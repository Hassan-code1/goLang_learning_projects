package main

import (
	"fmt"

)

type Student struct {
	name	 string
	id 		 int
	gpa		 int
}

type Registry struct {
	mp map[int]*Student;
	students []*Student;
}

func NewRegistry() *Registry {
	return &Registry{
		mp : make(map[int]*Student),
		students: make([]*Student, 0),
	}
}
func AddStudent(name string, id int, gpa int, register *Registry) bool{
	if register == nil {
		register = NewRegistry()
	}
	stud := Student{name, id, gpa}
	if _, ok := register.mp[id]; ok == true {
		return false
	}
	register.students = append(register.students, &stud)
	register.mp[id] = &stud
	return true
}
func PrintAllStudents(register *Registry) bool{
	if register == nil {
		return false
	}
	fmt.Println("Here is List of students Recorded in this Register:");
	if len(register.students) == 0 {
		fmt.Println("Their are O students registered currently")
	}
	for i := 0; i < len(register.students); i++ {
		fmt.Printf("%s student with id:%d have a gpa of %d\n", register.students[i].name, register.students[i].id, register.students[i].gpa)
	}
	return true
}

func findStudentWithId(register *Registry, id int) (*Student, bool){
	if register == nil {
		return nil, false
	}
	stud, ok := register.mp[id]
	if  ok == false {
		return nil, false
	}
	return stud, true
}

