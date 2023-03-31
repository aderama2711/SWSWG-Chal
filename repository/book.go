package repository

import (
	"P1/model"
	"time"
)

type BookRepo interface {
	CreateBook(model.Book) (res model.Book, err error)
	ListBook() (res []model.Book, err error)
	GetBook(id int) (res model.Book, err error)
	UpdateBook(newBook model.Book) (res model.Book, err error)
	DeleteBook(newBook model.Book) (err error)
}

func (r Repo) CreateBook(newBook model.Book) (res model.Book, err error) {
	err = r.db.Create(&newBook).Scan(&res).Error
	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) ListBook() (res []model.Book, err error) {
	err = r.db.Find(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) GetBook(id int) (res model.Book, err error) {
	err = r.db.First(&res, id).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) UpdateBook(newBook model.Book) (res model.Book, err error) {
	err = r.db.First(&res, newBook.ID).Error

	if err != nil {
		return res, err
	}

	err = r.db.Model(&res).Where("id = ?", newBook.ID).Updates(map[string]interface{}{"name_book": newBook.Name_book, "author": newBook.Author, "updated_at": time.Now()}).Scan(&res).Error

	if err != nil {
		return res, err
	}

	return res, nil
}

func (r Repo) DeleteBook(newBook model.Book) (err error) {
	err = r.db.First(&newBook, newBook.ID).Error

	if err != nil {
		return err
	}

	r.db.Delete(&newBook, newBook.ID)

	if err != nil {
		return err
	}

	return nil
}
