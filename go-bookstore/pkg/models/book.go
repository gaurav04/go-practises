package models

import (
	"github.com/go-bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name,omitempty"`
	Author      string `json:"author,omitempty"`
	Publication string `json:"publication,omitempty"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBookById(Id int64) Book {
	var book Book
	db.Unscoped().Delete(&book, "ID = ? ", Id) // unscoped will do permanent delete from DB
	//	db.Where("ID=?", Id).Delete(&book)
	return book
}
