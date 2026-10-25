package request

type UpdateProduct struct {
	Name        *string  `json:"name"        binding:"omitempty,min=2,max=255"`
	Slug        *string  `json:"slug"        binding:"omitempty,min=2,max=255"`
	Description *string  `json:"description" binding:"omitempty,max=1000"`
	Price       *float64 `json:"price"       binding:"omitempty,gt=0"`
	Stock       *int     `json:"stock"       binding:"omitempty,gte=0"`
}
