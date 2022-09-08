// Trying to build without receiver functions and get practice in with syntax

package main

import (
	"fmt"
)

type Book struct {
	Title, Author string
	HasBeenRead   bool
}

func main() {

	library := []Book{}
	status := "question"

	for status != "exit" {
		s, l := handleSelection(getUserSelection(), library)
		library = l
		status = s
	}

	return

}

func handleSelection(status string, lib []Book) (s string, l []Book) {
	switch status {
	case "add":
		book := getBookInfo()
		lib = append(lib, book)
	case "list":
		for _, v := range lib {
			fmt.Println(v)
		}
	case "remove":
		fmt.Println("functionality TBD")
	case "edit":
		fmt.Println("functionality TBD")
	case "seed":
		book1 := Book{"Sample Book 1", "John Smith", true}
		book2 := Book{"Money", "Adam Smith", true}
		book3 := Book{"Waterfalls", "Calvin Harris", false}
		lib = append(lib, book1, book2, book3)
	case "exit":
		return "exit", lib
	}
	return "question", lib

}

func getUserSelection() string {
	var choice string
	var choiceMap = map[string]string{
		"1": "add",
		"2": "remove",
		"3": "edit",
		"4": "seed",
		"5": "list",
		"6": "exit",
	}

	fmt.Println("What would you like to do?")
	fmt.Println("1. Add a book")
	fmt.Println("2. Remove a book")
	fmt.Println("3. Edit a book")
	fmt.Println("4. Seed library")
	fmt.Println("5. List Library")
	fmt.Println("6. Exit Library")
	fmt.Scan(&choice)

	return choiceMap[choice]
}

func getBookInfo() Book {
	var title, author string
	var haveRead bool

	fmt.Println("What is the title of the book?")
	fmt.Scan(&title)
	fmt.Println("Who is the author of the book?")
	fmt.Scan(&author)
	fmt.Println("Have you read this book? (true/false)")
	fmt.Scan(&haveRead)

	return Book{title, author, haveRead}
}
