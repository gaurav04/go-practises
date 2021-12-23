package models

type Book struct {
	ID     string  `json:"id,omitempty"`
	Isbn   string  `json:"isbn,omitempty"`
	Title  string  `json:"title,omitempty"`
	Author *Author `json:"author,omitempty"`
}

type Author struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var Books = []Book{
	{ID: "1", Isbn: "123", Title: "Book One", Author: &Author{Firstname: "Gaurav", Lastname: "Harsola"}},
	{ID: "2", Isbn: "456", Title: "Book Two", Author: &Author{Firstname: "Golu", Lastname: "Jain"}},
}
