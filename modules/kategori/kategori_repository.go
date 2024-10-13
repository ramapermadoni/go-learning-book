package kategori

import (
	"errors"

	"gorm.io/gorm"
)

// Repository interface for kategori operations
type Repository interface {
	InsertKategori(kategori *Kategori) error
	GetAllKategori(page, limit int, search string) ([]Kategori, int64, error)
	GetKategoriByID(id int64) (Kategori, error)
	UpdateKategori(kategori Kategori) error
	DeleteKategori(kategori Kategori) error
}

// kategoriRepository struct implementing Repository interface
type kategoriRepository struct {
	db *gorm.DB
}

// NewRepository initializes the repository with a DB connection
func NewRepository(db *gorm.DB) Repository {
	return &kategoriRepository{db: db}
}

// InsertKategori adds a new kategori to the database
func (r *kategoriRepository) InsertKategori(kategori *Kategori) error {
	return r.db.Table("kategori").Create(kategori).Error
}

// GetAllKategori retrieves kategori data with pagination and search
func (r *kategoriRepository) GetAllKategori(page, limit int, search string) ([]Kategori, int64, error) {
	var kategori []Kategori
	var total int64

	query := r.db.Table("kategori").Model(&Kategori{})

	if search != "" {
		query = query.Where("name LIKE ?", "%"+search+"%")
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	if err := query.Offset(offset).Limit(limit).Find(&kategori).Error; err != nil {
		return nil, 0, err
	}

	return kategori, total, nil
}

// GetKategoriByID retrieves a kategori by its ID
func (r *kategoriRepository) GetKategoriByID(id int64) (Kategori, error) {
	var kategori Kategori
	err := r.db.Table("kategori").First(&kategori, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Kategori{}, errors.New("kategori not found")
	}
	return kategori, err
}

// UpdateKategori updates an existing kategori
func (r *kategoriRepository) UpdateKategori(kategori Kategori) error {
	return r.db.Table("kategori").Save(&kategori).Error
}

// DeleteKategori removes a kategori by ID
func (r *kategoriRepository) DeleteKategori(kategori Kategori) error {
	return r.db.Table("kategori").Delete(&kategori).Error
}
