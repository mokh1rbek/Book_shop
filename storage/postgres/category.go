package postgres

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"

// 	"github.com/google/uuid"
// 	"github.com/jackc/pgx/v4/pgxpool"

// 	"github.com/mokh1rbek/Book_CRUD/models"
// 	"github.com/mokh1rbek/Book_CRUD/pkg/helper"
// )

// type BookRepo struct {
// 	db *pgxpool.Pool
// }

// func NewBookRepo(db *pgxpool.Pool) *BookRepo {
// 	return &BookRepo{
// 		db: db,
// 	}
// }

// func (f *BookRepo) Create(ctx context.Context, books *models.CreateBook) (string, error) {

// 	var (
// 		id    = uuid.New().String()
// 		query string
// 	)

// 	query = `
// 		INSERT INTO book (
// 			book_id,
// 			book_name,
// 			author_name,
// 			prise,
// 			date,
// 			update_at
// 		) VALUES ( $1, $2, $3, $4, $5, now())
// 	`

// 	_, err := f.db.Exec(ctx, query,
// 		id,
// 		books.BookName,
// 		books.AuthorName,
// 		books.Prise,
// 		books.Date,
// 	)

// 	if err != nil {
// 		return "", err
// 	}

// 	return id, nil
// }

// func (f *BookRepo) GetByPKey(ctx context.Context, pkey *models.BookPrimarKey) (*models.Book, error) {

// 	var (
// 		id          sql.NullString
// 		book_name   sql.NullString
// 		author_name sql.NullString
// 		prise       sql.NullString
// 		date        sql.NullString
// 		created_at  sql.NullString
// 		updated_at  sql.NullString
// 	)

// 	query := `
// 		SELECT
// 			book_id,
// 			book_name,
// 			author_name,
// 			prise,
// 			date,
// 			created_at,
// 			updated_at
// 		FROM
// 			book
// 		WHERE book_id = $1
// 	`

// 	err := f.db.QueryRow(ctx, query, pkey.Id).
// 		Scan(
// 			&id,
// 			&name,
// 			&parent_uuid,
// 			&created_at,
// 			&cpdated_at,
// 		)

// 	if err != nil {
// 		return nil, err
// 	}

// 	return &models.Book{
// 		ID:         id.String,
// 		Name:       name.String,
// 		ParentUUID: parent_uuid.String,
// 		CreatedAt:  created_at.String,
// 		UpdatedAt:  cpdated_at.String,
// 	}, nil
// }

// func (f *BookRepo) GetList(ctx context.Context, req *models.GetListBookRequest) (*models.GetListBookResponse, error) {

// 	var (
// 		resp   = models.GetListBookResponse{}
// 		offset = " OFFSET 0"
// 		limit  = " LIMIT 5"
// 	)

// 	if req.Limit > 0 {
// 		limit = fmt.Sprintf(" LIMIT %d", req.Limit)
// 	}

// 	if req.Offset > 0 {
// 		offset = fmt.Sprintf(" OFFSET %d", req.Offset)
// 	}

// 	query := `
// 		SELECT
// 			COUNT(*) OVER(),
// 			Book_id,
// 			name,
// 			parent_uuid,
// 			created_at,
// 			updated_at
// 		FROM
// 			Book
// 	`

// 	query += offset + limit

// 	rows, err := f.db.Query(ctx, query)

// 	for rows.Next() {

// 		var (
// 			id          sql.NullString
// 			name        sql.NullString
// 			parent_uuid sql.NullString
// 			created_at  sql.NullString
// 			cpdated_at  sql.NullString
// 		)

// 		err := rows.Scan(
// 			&resp.Count,
// 			&id,
// 			&name,
// 			&parent_uuid,
// 			&created_at,
// 			&cpdated_at,
// 		)

// 		if err != nil {
// 			return nil, err
// 		}

// 		resp.Book = append(resp.Book, &models.Book{
// 			ID:         id.String,
// 			Name:       name.String,
// 			ParentUUID: parent_uuid.String,
// 			CreatedAt:  created_at.String,
// 			UpdatedAt:  cpdated_at.String,
// 		})

// 	}

// 	return &resp, err
// }

// func (f *BookRepo) Update(ctx context.Context, req *models.UpdateBook) (int64, error) {

// 	var (
// 		query  = ""
// 		params map[string]interface{}
// 	)

// 	query = `
// 		UPDATE
// 			Book
// 		SET
// 			name = :name,
// 			parent_uuid = :parent_uuid,
// 			updated_at = now()
// 		WHERE Book_id = :Book_id
// 	`
// 	params = map[string]interface{}{
// 		"name":        req.Name,
// 		"parent_uuid": req.ParentUUID,
// 	}

// 	query, args := helper.ReplaceQueryParams(query, params)

// 	rowsAffected, err := f.db.Exec(ctx, query, args...)
// 	if err != nil {
// 		return 0, err
// 	}

// 	return rowsAffected.RowsAffected(), nil
// }

// func (f *BookRepo) Delete(ctx context.Context, req *models.BookPrimarKey) error {

// 	_, err := f.db.Exec(ctx, "DELETE FROM Book WHERE Book_id = $1", req.Id)
// 	if err != nil {
// 		return err
// 	}

// 	return err
// }
