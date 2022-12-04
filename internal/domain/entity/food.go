package entity

type Food struct {
	ID          int    `json:"id"`
	UserID      int    `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	FoodImage   string `json:"food_image"`
}
