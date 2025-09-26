package models

import (
	"github.com/hardikn13/book-management-system/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(Book{})
}

func FindAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func FindBook(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func (b *Book) AddBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func RemoveBook(Id int64) Book {
	var getBook Book
	db.Where("ID=?", Id).Delete(getBook)
	return getBook
}
