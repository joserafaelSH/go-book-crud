package models

import (
	"github.com/jinzhu/gorm"
	"github.com/joserafaelSH/book_management_system/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(255);unique_index"`
	Author      string `json:"author" gorm:"type:varchar(255)"`
	Publication string `json:"publication" gorm:"type:varchar(255)"`
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
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(bookId int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", bookId).Find(&getBook)
	return &getBook, db
}

func (b *Book) UpdateBook() *Book {
	db.Save(&b)
	return b
}

func DeleteBook(bookId int64) Book {
	var book Book
	db.Where("ID = ?", bookId).Delete(book)
	return book
}
