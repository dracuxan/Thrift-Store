package models

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password []byte `json:"-"`
}

type Admin struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Password []byte `json:"-"`
}

type Product struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	Img       string  `json:"img"`
	Title     string  `json:"title"`
	Reviews   string  `json:"reviews"`   // Number of reviews as a string (e.g., "(3 reviews)")
	PrevPrice string  `json:"prevPrice"` // Previous price as a string
	NewPrice  float64 `json:"newPrice"`  // New price as a string
	Company   string  `json:"company"`
	Color     string  `json:"color"`
	Category  string  `json:"category"`
}

type CartItem struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	Img        string  `json:"img"`
	Title      string  `json:"title"`
	ProductID  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	TotalPrice float64 `json:"total_price"`
}

type CartResponse struct {
	Cart []CartItem `json:"cart"` // Use the appropriate type for CartItem
}
