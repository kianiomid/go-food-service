package dto

type FoodDetailViewModel struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	FoodImage   string `json:"food_image"`
	UserName    string `json:"user_name"`
}

type FoodViewModel struct {
	ID          int64  `json:"id"`
	UserID      int64  `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FoodImage   string `json:"food_image"`
}
