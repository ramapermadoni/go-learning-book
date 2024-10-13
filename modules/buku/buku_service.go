package buku

import (
	"errors"
	"go-learning-book/modules/kategori"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Service interface for buku operations
type Service interface {
	CreateBukuService(ctx *gin.Context) (Buku, error)
	GetAllBukuService(ctx *gin.Context) ([]Buku, int64, error)
	GetBukuByIDService(ctx *gin.Context) (Buku, error)
	UpdateBukuService(ctx *gin.Context) error
	DeleteBukuService(ctx *gin.Context) error
}

type bukuService struct {
	repo      Repository
	validator *validator.Validate
}

// NewBukuService initializes the service with the repository
func NewBukuService(repo Repository) Service {
	return &bukuService{repo: repo,
		validator: validator.New()}
}

// CreateBukuService handles the creation of a new buku
func (s *bukuService) CreateBukuService(ctx *gin.Context) (Buku, error) {
	var buku Buku
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		return Buku{}, err
	}

	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return Buku{}, errors.New("username not found in context")
	}

	// Set CreatedBy field
	buku.CreatedBy = username
	if buku.TotalPage > 100 {
		buku.Thickness = "Tebal"
	} else {
		buku.Thickness = "Tipis"
	}

	if err := s.validator.Struct(buku); err != nil {
		return Buku{}, err
	}

	err := s.repo.InsertBuku(&buku)
	return buku, err
}

// GetAllBukuService retrieves all bukus with pagination and search
func (s *bukuService) GetAllBukuService(ctx *gin.Context) ([]Buku, int64, error) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	search := ctx.Query("search")

	// Parse page and limit to integers (implement error handling as needed)
	// Example: page, limit = parseQueryParams(page, limit)

	return s.repo.GetAllBuku(page, limit, search)
}

// GetBukuByIDService retrieves a buku by ID
func (s *bukuService) GetBukuByIDService(ctx *gin.Context) (Buku, error) {
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return Buku{}, errors.New("invalid ID")
	}
	return s.repo.GetBukuByID(id)
}

// UpdateBukuService updates an existing buku
func (s *bukuService) UpdateBukuService(ctx *gin.Context) error {
	var buku Buku

	// Parse ID from URL parameter
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the buku exists in the database
	existingBuku, err := s.repo.GetBukuByID(id)
	if err != nil {
		return errors.New("buku not found")
	}

	// Bind JSON input to the buku struct
	if err := ctx.ShouldBindJSON(&buku); err != nil {
		return errors.New("invalid input data")
	}

	// Retrieve the username from context and check if it exists
	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return errors.New("username not found in context")
	}

	// Set Modified field
	buku.ModifiedAt = time.Now()
	buku.ModifiedBy = username
	// Use the existing buku ID
	buku.ID = existingBuku.ID

	// Update the buku in the database
	if err := s.repo.UpdateBuku(buku); err != nil {
		return errors.New("failed to update buku")
	}

	return nil
}

// DeleteBukuService deletes a buku by ID
func (s *bukuService) DeleteBukuService(ctx *gin.Context) error {
	// Parse ID from URL parameter
	id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the buku exists in the database
	if _, err := s.repo.GetBukuByID(id); err != nil {
		return errors.New("buku not found")
	}

	// Create a buku struct with the given ID
	buku := Buku{ID: int(id)}

	// Delete the buku from the database
	if err := s.repo.DeleteBuku(buku); err != nil {
		return errors.New("failed to delete buku")
	}

	return nil
}
func ValidateCategoryID(db *gorm.DB, categoryID int) error {
	var kategori kategori.Kategori
	result := db.First(&kategori, categoryID)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("invalid category ID")
	}
	return nil
}
