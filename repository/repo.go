package repository

import "gorm.io/gorm"

type Repo struct {
	db *gorm.DB
}

type RepoInterface interface {
	BookRepo
}

func NewRepo(db *gorm.DB) *Repo {
	return &Repo{db: db}
}
