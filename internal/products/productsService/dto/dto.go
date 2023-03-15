package dto

type CreateProduct struct {
	Name  string  `json:"name" validate:"required"`
	Price float64 `json:"price,omitempty" validate:"required"`
	Qty   float32 `json:"qty,omitempty" validate:"required"`
}
type ProductCreateDto CreateProduct

type FindProduct struct {
	Id string
}
type DeleteProduct struct {
	Id string
}
