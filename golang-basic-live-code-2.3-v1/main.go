package main

import (
	"fmt"
)

func ViewStudents(students map[string][]interface{}) {
	fmt.Println("Name\tAddress\tPhone\tScore")
	for name, info := range students {
		address, _ := info[0].(string)
		phone, _ := info[1].(string)
		score, _ := info[2].(int)
		fmt.Printf("%s\t%s\t%s\t%d\n", name, address, phone, score)
	}
}

func AddStudent(students *map[string][]interface{}) func(string, string, string, int) {
	return func(name string, address string, phone string, score int) {
		(*students)[name] = []interface{}{address, phone, score}
	}
}

func RemoveStudent(students *map[string][]interface{}) func(string) {
	return func(name string) {
		delete(*students, name)
	} // TODO: replace this
}

func HighestScore(students *map[string][]interface{}) (string, int) {
	HighestScore := 0
	highestName := ""
	for name, info := range *students {
		score, _ := info[2].(int)
		if score > HighestScore {
			HighestScore = score
			highestName = name
		}
	}

	return highestName, HighestScore // TODO: replace this
}

func main() {
	students := make(map[string][]interface{})
	add := AddStudent(&students)
	remove := RemoveStudent(&students)

	for {
		var command string
		fmt.Print("Enter command (add, remove, high-score, view): ")
		fmt.Scan(&command)

		switch command {
		case "add":
			var name, address, phone string
			var score int
			fmt.Print("Enter name: ")
			fmt.Scan(&name)
			fmt.Print("Enter address: ")
			fmt.Scan(&address)
			fmt.Print("Enter phone: ")
			fmt.Scan(&phone)
			fmt.Print("Enter score: ")
			fmt.Scan(&score)

			add(name, address, phone, score)
		case "remove":
			var name string
			fmt.Print("Enter name: ")
			fmt.Scan(&name)

			remove(name)
		case "high-score":
			name, score := HighestScore(&students)
			fmt.Println("High Score: ", name, score)
		case "view":
			fmt.Println("Student data:")
			ViewStudents(students)
		default:
			fmt.Println("Invalid command. Available commands: add, remove, high-score, view")
		}
	}
}
