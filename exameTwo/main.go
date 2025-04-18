package main

import (
	"exameTwo/model"
	"fmt"
)

func main() {
	fmt.Println("Testando funcionalidades...")

	collectionBooks := model.MediaCollection[model.Book]{}

	collectionBooks.AddItem(model.Book{
		Title: "Harry Potter 1",
		Autor: "JK",
		Pages: 230,
		Year:  1998,
	})

	collectionBooks.AddItem(model.Book{
		Title: "Harry Potter 3",
		Autor: "JK",
		Pages: 234,
		Year:  2004,
	})

	collectionBooks.AddItem(model.Book{
		Title: "Harry Potter 8",
		Autor: "JK",
		Pages: 400,
		Year:  2008,
	})

	collectionBooks.AddItem(model.Book{
		Title: "Lords of the Rings",
		Autor: "KL",
		Pages: 506,
		Year:  2008,
	})

	fmt.Println("Books from 2008: ", collectionBooks.ListByYear(2008))
}
