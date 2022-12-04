package dto

type FoodDetailViewModel struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FoodImage   string `json:"food_image"`
	UserName    string `json:"user_name"`
}

type FoodViewModel struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FoodImage   string `json:"food_image"`
}
