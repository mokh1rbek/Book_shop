package storage

import (
	"context"

	"github.com/mokh1rbek/Book_CRUD/models"
)

type StorageI interface {
	CloseDB()
	Book() BookRepoI
}

type BookRepoI interface {
	Create(ctx context.Context, req *models.CreateBook) (string, error)
	GetByPKey(ctx context.Context, req *models.BookPrimarKey) (*models.Book, error)
	GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error)
	Update(ctx context.Context, req *models.UpdateBook) (int64, error)
	Delete(ctx context.Context, req *models.BookPrimarKey) error
}
