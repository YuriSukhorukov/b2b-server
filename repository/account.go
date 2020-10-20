package repository

import models "github.com/b2b-server/model"

type AccountRepository interface {
	Fetch(cursor string, num int64) ([]*models.Article, error)
	GetByID(id int64) (*models.Article, error)
	Update(article *models.Article) (*models.Article, error)
	Store(a *models.Article) (int64, error)
	Delete(id int64) (bool, error)
}