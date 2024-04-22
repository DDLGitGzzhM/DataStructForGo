package main

import "fmt"

type Person struct {
	Name string
}

/*
UpdateName
Go语言没有引用传递，但是为什么这里能够传递过去，是因为间接进行了引用传递
*/
func UpdateName(p *Person, newName string) {
	p.Name = newName
}

func main() {
	person := Person{Name: "Alice"}
	fmt.Println("Before:", person.Name)
	UpdateName(&person, "Bob")
	fmt.Println("After:", person.Name)
}
