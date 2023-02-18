package entity

type InputBook struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageUrl    string `json:"image_url" binding:"url"`
	ReleaseYear int    `json:"release_year" binding:"min=1980,max=2021"`
	Price       string `json:"price"`
	TotalPage   int    `json:"total_page"`
	CategoryId  int    `json:"category_id"`
}
