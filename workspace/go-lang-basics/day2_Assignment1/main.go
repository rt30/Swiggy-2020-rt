package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Animal interface{
	Eat () string
	Sleep () string
	Breathe () string
}
//func (e Emp) Speak() string  {
//	return "hihiiiihihi"
//}

func (e Emp ) Eat() string{
	return "Yes"
}
func (e Emp) Sleep() string{
	return "Daily"
}
func (e Emp) Breathe() string{
	return "Always"
}

type Emp struct{

	EmpId int
	EmpName string
	Salary float64

	Address struct{
		HouseId int
		Street string
		city string
		pin int
	}
}

func main(){

	employees := make([] Emp, 5, 10)
	fmt.Println(cap(employees))
	fmt.Println(len(employees))

	println("Enter Details :")

	//11
	num_emp := 0;

	for i:= 0; i < len(employees); i++ {
	//for i,_ := range employees {
		fmt.Println("hi",i)

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter Your Name")
		employees[i].EmpName, _ = reader.ReadString('\n')

		employees[i].EmpId = i+1

		fmt.Println("Enter Your Salary")
		strSalary, _ := reader.ReadString('\n')
		employees[i].Salary, _ = strconv.ParseFloat( strings.TrimSpace(strSalary) ,64)

		print("Do you want to continue? (y/n) \n")
		ans, _ := reader.ReadString('\n')

		strings.TrimRight(ans,"\n")
		strings.TrimSpace(ans)

		if strings.TrimRight(ans, "\n") == "n" {
			num_emp = i
			break
		} else {
			if len(employees) == i+1 {
				fmt.Println("incrementing capacity ...")
				employees = employees[:cap(employees)]

				//fmt.Println(cap(employees))
				//fmt.Println(len(employees))
				//continue
			}
		}


	}
	fmt.Println("Employee Details are as follows :\n")
	fmt.Printf("Employee Id\t Employee Name\t Employee Salary\n")
	for i:=0 ; i<= num_emp ; i++ {
		fmt.Printf("\n %v\t, %v\t, %.2f", employees[i].EmpId, employees[i].EmpName, employees[i].Salary)
	}

	for _, emps := range employees{
		fmt.Println(emps.Eat())
		fmt.Println(emps.Sleep())
		fmt.Println(emps.Breathe())
	}

}

