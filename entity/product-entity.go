package entity

type Product struct {
	Code  string `json:"code" binding:"required"`
	Price string `json:"price" binding:"required"`
}
type ProductUpdate struct {
	Code  string `json:"code"`
	Price string `json:"price"`
}
