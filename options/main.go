package main

import (
	"fmt"
	"github.com/jinzhu/copier"
	"log"
)

type User struct {
	Name        string
	Role        string
	Age         int32
	EmployeCode int64 `copier:"EmployeNum"` // specify field name

	// Explicitly ignored in the destination struct.
	Salary   int
}

type Employee struct {
	// Tell copier.Copy to panic if this field is not copied.
	Name      string `copier:"must"`

	// Tell copier.Copy to return an error if this field is not copied.
	Age       int32  `copier:"must,nopanic"`

	// Tell copier.Copy to explicitly ignore copying this field.
	Salary    int    `copier:"-"`

	DoubleAge int32
	EmployeId int64 `copier:"EmployeNum"` // specify field name
	SuperRole string
}

func main() {
	var (
		user      = User{Name: "Jinzhu", Age: 18, Role: "Admin", Salary: 200000}
		employee  = Employee{Salary: 150000}
	)

	// 将 user 的字段赋值给 employee，除明确指名不复制的字段，其余字段全部复制
	//err := copier.Copy(&employee, &user)
	err := copier.CopyWithOption(&employee, &user, copier.Option{DeepCopy: true, IgnoreEmpty: true})
	if err != nil {
		log.Fatal("copy data failed: ", err)
	}

	fmt.Printf("%#v \n", employee)
}
