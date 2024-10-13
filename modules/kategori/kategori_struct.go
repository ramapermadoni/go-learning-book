package kategori

import (
	"time"
)

type Kategori struct {
	ID         int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string    `json:"name" validate:"required,min=3,max=100"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy  string    `json:"created_by" validate:"required"`
	ModifiedAt time.Time `json:"modified_at" gorm:"autoCreateTime"`
	ModifiedBy string    `json:"modified_by"`
}
