package main

import "fmt"

type Engineer struct {
	Name        string
	Age         int
	ProjectName Project
}

type Project struct {
	Name         string
	Priority     string
	Technologies []string
}

func (e Engineer) Print() {
	fmt.Println("Engineer Information")
	fmt.Printf("Name: %s\n", e.Name)
	fmt.Printf("Age: %d\n", e.Age)
	fmt.Printf("Current Project handled are: \n - %s \n - %s \n", e.ProjectName.Technologies[0], e.ProjectName.Technologies[1])
}

func (e *Engineer) UpdateAge(age int) {
	e.Age = age
}

func main() {
	fmt.Println("Structs Tutorial")

	engineer := Engineer{Name: "Elliot", Age: 33, ProjectName: Project{
		Name:         "Learning Go programming",
		Priority:     "High",
		Technologies: []string{"Go", "OperataorSDK"},
	},
	}

	//fmt.Printf("%+v\n", engineer)

	//fmt.Println(engineer.Name)
	// fmt.Println(engineer.ProjectName.Technologies[0])
	engineer.Print()
	engineer.UpdateAge(44)
	engineer.Print()

}
