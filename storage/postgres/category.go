package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"

	"github.com/mokh1rbek/Book_CRUD/models"
	"github.com/mokh1rbek/Book_CRUD/pkg/helper"
)

type BookRepo struct {
	db *pgxpool.Pool
}

func NewBookRepo(db *pgxpool.Pool) *BookRepo {
	return &BookRepo{
		db: db,
	}
}

func (f *BookRepo) Create(ctx context.Context, books *models.CreateBook) (string, error) {

	var (
		id    = uuid.New().String()
		query string
	)

	query = `
		INSERT INTO book (
			book_id,
			book_name,
			author_name,
			prise,
			date,
			update_at
		) VALUES ( $1, $2, $3, $4, $5, now())
	`

	_, err := f.db.Exec(ctx, query,
		id,
		books.BookName,
		books.AuthorName,
		books.Prise,
		books.Date,
	)

	if err != nil {
		return "", err
	}

	return id, nil
}

func (f *BookRepo) GetByPKey(ctx context.Context, pkey *models.BookPrimarKey) (*models.Book, error) {

	var (
		id          sql.NullString
		book_name   sql.NullString
		author_name sql.NullString
		prise       sql.NullString
		date        sql.NullString
		created_at  sql.NullString
		updated_at  sql.NullString
	)

	query := `
		SELECT
			book_id,
			book_name,
			author_name,
			prise,
			date,
			created_at,
			updated_at
		FROM
			book
		WHERE book_id = $1
	`

	err := f.db.QueryRow(ctx, query, pkey.Id).
		Scan(
			&id,
			&book_name,
			&author_name,
			&prise,
			&date,
			&created_at,
			&updated_at,
		)

	if err != nil {
		return nil, err
	}

	return &models.Book{
		Id:         id.String,
		BookName:   book_name.String,
		AuthorName: author_name.String,
		Prise:      prise.String,
		Date:       date.String,
		CreatedAt:  created_at.String,
		UpdatedAt:  updated_at.String,
	}, nil
}

func (f *BookRepo) GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error) {

	var (
		resp   = models.GetListBookResponse{}
		offset = " OFFSET 0"
		limit  = " LIMIT 5"
	)

	if req.Limit > 0 {
		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
	}

	if req.Offset > 0 {
		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
	}

	query := `
		SELECT
			COUNT(*) OVER(),
			book_id,
			name,
			parent_uuid,
			created_at,
			updated_at
		FROM
			book
	`

	query += offset + limit

	rows, err := f.db.Query(ctx, query)

	for rows.Next() {

		var (
			id          sql.NullString
			name        sql.NullString
			parent_uuid sql.NullString
			created_at  sql.NullString
			cpdated_at  sql.NullString
		)

		err := rows.Scan(
			&resp.Count,
			&id,
			&name,
			&parent_uuid,
			&created_at,
			&cpdated_at,
		)

		if err != nil {
			return nil, err
		}

		resp.Books = append(resp.Books, &models.Book{
			Id:         id.String,
			BookName:   "",
			AuthorName: "",
			Prise:      "",
			Date:       "",
			CreatedAt:  "",
			UpdatedAt:  "",
		})

	}

	return &resp, err
}

func (f *BookRepo) Update(ctx context.Context, req *models.UpdateBook) (int64, error) {

	var (
		query  = ""
		params map[string]interface{}
	)

	query = `
		UPDATE
			book
		SET
			name = :name,
			parent_uuid = :parent_uuid,
			updated_at = now()
		WHERE Book_id = :Book_id
	`
	params = map[string]interface{}{
		"name":        req.AuthorName,
		"parent_uuid": req.BookName,
	}

	query, args := helper.ReplaceQueryParams(query, params)

	rowsAffected, err := f.db.Exec(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	return rowsAffected.RowsAffected(), nil
}

func (f *BookRepo) Delete(ctx context.Context, req *models.BookPrimarKey) error {

	_, err := f.db.Exec(ctx, "DELETE FROM Book WHERE Book_id = $1", req.Id)
	if err != nil {
		return err
	}

	return err
}
