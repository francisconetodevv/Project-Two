package model

import "time"

/* Interface */
type Media interface {
	GetTitle() string
	GetYearPublished() int
}

/* Normal struct */
type Book struct {
	Title string
	Autor string
	Pages int
	Year  int
}

type Film struct {
	Title    string
	Director string
	TimeFilm time.Time
	Year     int
}

/* 1. Book: Methods for the interface Book */
func (b Book) GetTitle() string {
	return b.Title
}

func (b Book) GetYearPublished() int {
	return b.Year
}

/* 2. Film: Methods for the interface film */
func (f Film) GetTitle() string {
	return f.Title
}

func (f Film) GetYearPublished() int {
	return f.Year
}

/* Generic struct - Collection */
type MediaCollection[M Media] struct {
	Collection []M
}

/* Method to add an media into the collection of media */
/* As always if you would like to change the value of something
It is necessary to pass the memory address (*)*/
func (m *MediaCollection[M]) AddItem(media M) {
	m.Collection = append(m.Collection, media)
}

/* Method to list the medias by year */
func (m *MediaCollection[M]) ListByYear(Year int) []M {
	var result []M

	for _, item := range m.Collection {
		if item.GetYearPublished() == Year {
			result = append(result, item)
		}
	}

	return result
}
