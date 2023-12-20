package service

import (
	"pinjammodal/go-storage/dto"
	"pinjammodal/go-storage/model"
	"pinjammodal/go-storage/repository"
)

type BookService struct {
	repository *repository.BookRepository
}

func New(repository *repository.BookRepository) interface{} {
	return &BookService{repository: repository}
}

func (s *BookService) Create(book dto.Book) (*model.Book, error) {
	bookModel := model.Book{
		Name: book.Name,
	}

	err := s.repository.Create(bookModel)

	if err != nil {
		return nil, err
	}

	return &bookModel, nil
}
