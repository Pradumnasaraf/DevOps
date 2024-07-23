package main

import "fmt"

type Printer interface {
	printInfo()
}


type Book struct {
	Title string
	Price float32
}

type Drink struct {
	Name  string
	Price float32
}

func (b Book) printInfo() {
	fmt.Println("Book Title:", b.Title, "Price:", b.Price)
}

func (d Drink) printInfo() {
	fmt.Println("Drink Name:", d.Name, "Price:", d.Price)
}

func main() {
	book := Book{Title: "The Alchemist", Price: 9.99}
	drink := Drink{Name: "Coke", Price: 1.99}

	// book.printInfo()
	// drink.printInfo()

	info := []Printer{book, drink}

	info[0].printInfo()
	info[1].printInfo()
}
