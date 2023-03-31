package service

import "P1/model"

type BookService interface {
	CreateBook(model.Book) (res model.Book, err error)
	ListBook() (res []model.Book, err error)
	GetBook(id int) (res model.Book, err error)
	UpdateBook(newBook model.Book) (res model.Book, err error)
	DeleteBook(newBook model.Book) (err error)
}

func (s *Service) CreateBook(newBook model.Book) (res model.Book, err error) {
	res, err = s.repo.CreateBook(newBook)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) ListBook() (res []model.Book, err error) {
	res, err = s.repo.ListBook()
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) GetBook(id int) (res model.Book, err error) {
	res, err = s.repo.GetBook(id)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) UpdateBook(newBook model.Book) (res model.Book, err error) {
	res, err = s.repo.UpdateBook(newBook)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (s *Service) DeleteBook(newBook model.Book) (err error) {
	err = s.repo.DeleteBook(newBook)
	if err != nil {
		return err
	}
	return nil
}
