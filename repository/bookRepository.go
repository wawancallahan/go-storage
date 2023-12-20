package repository

import (
	"pinjammodal/go-storage/database"
	"pinjammodal/go-storage/model"
)

type BookRepository struct {
	db *database.Database
}

func New(db *database.Database) interface{} {
	return &BookRepository{db}
}

func (r *BookRepository) FindAll() ([]model.Book, error) {
	var book []model.Book

	err := r.db.Find(&book).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *BookRepository) Find(id string) (model.Book, error) {
	var book model.Book

	err := r.db.Find(&book, id).Error

	if err != nil {
		return book, err
	}

	return book, nil
}

func (r *BookRepository) Destroy(id string) error {
	var book model.Book
	err := r.db.Where("id = ?", id).Delete(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Create(book model.Book) error {
	err := r.db.Create(&book).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *BookRepository) Update(book model.Book) error {
	err := r.db.Save(&book).Error

	if err != nil {
		return err
	}

	return nil
}
