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
		status := getUserSelection()
		s, l := handleSelection(status, library)
		library = l
		status = s
		fmt.Println("after handleSelection", library)
	}

	return

}

func handleSelection(status string, lib []Book) (s string, l []Book) {
	fmt.Println("handling status", status)
	switch status {
	case "add":
		book := getBookInfo()
		lib = append(lib, book)
		return "question", lib

	case "remove":

	case "edit":

	case "seed":

	case "exit":

		return "exit", lib
	default:
		return "exit", lib
	}
	return
}

func getUserSelection() string {
	var choice string
	var choiceMap = map[string]string{
		"1": "add",
		"2": "remove",
		"3": "edit",
		"4": "seed",
		"5": "exit",
	}

	fmt.Println("What would you like to do?")
	fmt.Println("1. Add a book")
	fmt.Println("2. Remove a book")
	fmt.Println("3. Edit a book")
	fmt.Println("4. Seed library")
	fmt.Println("5. Exit Library")
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
