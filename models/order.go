package models

type OrderPrimarKey struct {
	Id string `json:"order_id"`
}

type CreateOrderSwagger struct {
	UserId string `json:"user_id"`
	BookId string `json:"book_id"`
}

type CreateOrder struct {
	UserId string  `json:"user_id"`
	BookId string  `json:"book_id"`
	Payed   float64 `json:"payed"`
}

type Order struct {
	Id        string  `json:"order_id"`
	UserId   string  `json:"user_id"`
	BookId   string  `json:"book_id"`
	Payed     float64 `json:"payed"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type OrderGroup struct {
	FullName   string  `json:"fullname"`
	TotalPayed float64 `json:"total_payed"`
}

type UpdateOrderSwagger struct {
	UserId string `json:"user_id"`
	BookId string `json:"book_id"`
}

type UpdateOrder struct {
	Id        string  `json:"order_id"`
	UserId   string  `json:"user_id"`
	BookId   string  `json:"book_id"`
	Payed     float64 `json:"payed"`
	UpdatedAt string  `json:"updated_at"`
}

type GetListOrderRequest struct {
	Limit  int32
	Offset int32
}

type GetListOrderResponse struct {
	Count  int32         `json:"count"`
	Orders []*OrderGroup `json:"orders"`
}
