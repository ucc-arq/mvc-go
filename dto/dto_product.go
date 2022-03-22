package dto

type ProductDto struct {
	Name string `json:"name"`
	UniversalCode string `json:"universal_code"`
	Picture string `json:"picture_url"`
	Price float32 `json:"base_price"`
}

type ProductsDto []ProductDto

