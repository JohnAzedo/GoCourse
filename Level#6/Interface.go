package main

import (
	"fmt"
)

type Article struct {
	title  string
	author string
}

func (article Article) String() string {
	return fmt.Sprintf("The %q article was written by %v.", article.title, article.author)
}

type Book struct {
	title    string
	author   string
	numPages int
}

func (book Book) String() string {
	return fmt.Sprintf("The %q book was written by %v.", book.title, book.author)
}

type Stringer interface {
	String() string
}

func main() {
	book := Book{
		title:  "Harry Potter",
		author: "J. K. Rowling",
	}

	article := Article{
		title:  "Understanding Interfaces in Go",
		author: "Sammy Shark",
	}

	Print(book)
	Print(article)

}

func Print(stringer Stringer) {
	fmt.Println(stringer.String())
}
