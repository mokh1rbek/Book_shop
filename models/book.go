package models

type BookPrimarKey struct {
	Id string `json:"book_id"`
}

type CreateBook struct {
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	Prise      string `json:"prise"`
	Date       string `json:"date"`
}
type Book struct {
	Id         string `json:"book_id"`
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	Prise      string `json:"prise"`
	Date       string `json:"date"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

type UpdateBook struct {
	Id         string `json:"book_id"`
	BookName   string `json:"book_name"`
	AuthorName string `json:"author_name"`
	Prise      string `json:"prise"`
	Date       string `json:"date"`
}

type GetListBookRequest struct {
	Limit  int32
	Offset int32
}

type GetListBookResponse struct {
	Count int32   `json:"count"`
	Books []*Book `json:"book"`
}
