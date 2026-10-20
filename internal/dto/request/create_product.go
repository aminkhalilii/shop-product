package request

type CreateProduct struct {
	Name        string  `json:"name" binding:"required,min=2,max=255"`
	Slug        string  `json:"slug" binding:"required,min=2,max=255"`
	Description string  `json:"description" binding:"max=1000"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	Stock       int     `json:"stock" binding:"gte=0"`
}
