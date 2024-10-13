package kategori

import (
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// KategoriService defines the service interface
type KategoriService interface {
	CreateKategoriService(ctx *gin.Context) (Kategori, error)
	GetAllKategoriService(ctx *gin.Context) ([]Kategori, int64, error)
	GetKategoriByIDService(ctx *gin.Context) (Kategori, error)
	UpdateKategoriService(ctx *gin.Context) error
	DeleteKategoriService(ctx *gin.Context) error
}

type kategoriService struct {
	repo      Repository
	validator *validator.Validate
}

// NewKategoriService initializes the service with a repository
func NewKategoriService(repo Repository) KategoriService {
	return &kategoriService{
		repo:      repo,
		validator: validator.New(),
	}
}

// CreateKategoriService handles the creation of a new kategori
func (s *kategoriService) CreateKategoriService(ctx *gin.Context) (Kategori, error) {
	var kategori Kategori
	if err := ctx.ShouldBindJSON(&kategori); err != nil {
		return Kategori{}, err
	}

	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return Kategori{}, errors.New("username not found in context")
	}

	// Set CreatedBy field
	kategori.CreatedBy = username

	if err := s.validator.Struct(kategori); err != nil {
		return Kategori{}, err
	}

	if err := s.repo.InsertKategori(&kategori); err != nil {
		return Kategori{}, errors.New("failed to add new kategori")
	}
	return kategori, nil
}

// GetAllKategoriService retrieves all categories with pagination and search
func (s *kategoriService) GetAllKategoriService(ctx *gin.Context) ([]Kategori, int64, error) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	search := ctx.Query("search")

	return s.repo.GetAllKategori(page, limit, search)
}

// GetKategoriByIDService retrieves a kategori by its ID
func (s *kategoriService) GetKategoriByIDService(ctx *gin.Context) (Kategori, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return Kategori{}, errors.New("invalid ID")
	}

	return s.repo.GetKategoriByID(id)
}

// UpdateKategoriService updates an existing kategori
func (s *kategoriService) UpdateKategoriService(ctx *gin.Context) error {
	var kategori Kategori

	// Parse the ID from the URL parameter
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the kategori exists in the database
	existingKategori, err := s.repo.GetKategoriByID(id)
	if err != nil {
		return errors.New("kategori not found")
	}

	// Bind the JSON request to the kategori struct
	if err := ctx.ShouldBindJSON(&kategori); err != nil {
		return errors.New("invalid input data")
	}

	// Retrieve the username from context and check if it exists
	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return errors.New("username not found in context")
	}

	// Set Modified field
	kategori.ModifiedAt = time.Now()
	kategori.ModifiedBy = username
	// Use the existing kategori ID
	kategori.ID = existingKategori.ID

	// Update the kategori in the database
	if err := s.repo.UpdateKategori(kategori); err != nil {
		return errors.New("failed to update kategori")
	}

	return nil
}

// DeleteKategoriService deletes a kategori by ID
func (s *kategoriService) DeleteKategoriService(ctx *gin.Context) error {
	// Parse the ID from the URL parameter
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the kategori exists in the database
	if _, err := s.repo.GetKategoriByID(id); err != nil {
		return errors.New("kategori not found")
	}

	// Create a kategori struct with the given ID
	kategori := Kategori{ID: int(id)}

	// Delete the kategori from the database
	if err := s.repo.DeleteKategori(kategori); err != nil {
		return errors.New("failed to delete kategori")
	}

	return nil
}
