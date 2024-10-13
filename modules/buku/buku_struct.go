package buku

import "time"

// Buku struct to represent a book.
type Buku struct {
	ID          int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	ReleaseYear int       `json:"release_year" validate:"required,min=1980,max=2024"`
	Price       int       `json:"price" validate:"required"`
	TotalPage   int       `json:"total_page" validate:"required"`
	Thickness   string    `json:"thickness" validate:"required"`
	CategoryID  int       `json:"category_id" validate:"required"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy   string    `json:"created_by" validate:"required"`
	ModifiedAt  time.Time `json:"modified_at"`
	ModifiedBy  string    `json:"modified_by"`
}
