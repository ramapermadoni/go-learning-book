package user

import "time"

// User struct to represent a user.
type User struct {
	ID         int64     `json:"id" gorm:"primaryKey;autoIncrement"`        // Tambahkan validate
	Username   string    `json:"username" validate:"required,min=3,max=50"` // Tambahkan validate
	Password   string    `json:"password" validate:"required,min=6"`        // Tambahkan validate
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`          // Tambahkan validate jika perlu
	CreatedBy  string    `json:"created_by" validate:"required"`            // Tambahkan validate
	ModifiedAt time.Time `json:"modified_at" gorm:"autoCreateTime"`         // Tambahkan validate jika perlu
	ModifiedBy string    `json:"modified_by"`                               // Tambahkan validate
}
