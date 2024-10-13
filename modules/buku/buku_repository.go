package buku

import (
	"errors"

	"gorm.io/gorm"
)

// Repository interface for buku operations
type Repository interface {
	InsertBuku(buku *Buku) error
	GetAllBuku(page, limit int, search string) ([]Buku, int64, error)
	GetBukuByID(id int64) (Buku, error)
	UpdateBuku(buku Buku) error
	DeleteBuku(buku Buku) error
}

// bukuRepository struct implementing Repository interface
type bukuRepository struct {
	db *gorm.DB
}

// NewRepository initializes the repository with a DB connection
func NewRepository(db *gorm.DB) Repository {
	return &bukuRepository{db: db}
}

// InsertBuku adds a new buku to the database
func (r *bukuRepository) InsertBuku(buku *Buku) error {
	return r.db.Table("buku").Create(buku).Error
}

// GetAllBuku retrieves buku data with pagination and search
func (r *bukuRepository) GetAllBuku(page, limit int, search string) ([]Buku, int64, error) {
	var bukus []Buku
	var total int64

	query := r.db.Table("buku").Model(&Buku{})

	if search != "" {
		query = query.Where("title LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&bukus).Error; err != nil {
		return nil, 0, err
	}

	return bukus, total, nil
}

// GetBukuByID retrieves a buku by its ID
func (r *bukuRepository) GetBukuByID(id int64) (Buku, error) {
	var buku Buku
	err := r.db.Table("buku").First(&buku, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Buku{}, errors.New("buku not found")
	}
	return buku, err
}

// UpdateBuku updates an existing buku
func (r *bukuRepository) UpdateBuku(buku Buku) error {
	return r.db.Table("buku").Save(&buku).Error
}

// DeleteBuku removes a buku by ID
func (r *bukuRepository) DeleteBuku(buku Buku) error {
	return r.db.Table("buku").Delete(&buku).Error
}
