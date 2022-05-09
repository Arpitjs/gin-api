package entity

type Product struct {
	Code  string `db:"code" json:"code" binding:"required"`
	Price string `db:"price" json:"price" binding:"required"`
}
type ProductUpdate struct {
	Code  string `json:"code"`
	Price string `json:"price"`
}
