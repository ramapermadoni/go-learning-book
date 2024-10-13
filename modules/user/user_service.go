package user

import (
	"errors"
	"go-learning-book/utils/common"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// UserService defines the service interface for users
type UserService interface {
	GetAllUsers(ctx *gin.Context) ([]User, int64, error)
	GetUserByID(ctx *gin.Context) (User, error)
	CreateUser(ctx *gin.Context) (User, error)
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
}

type userService struct {
	repo      UserRepository
	validator *validator.Validate
}

// NewUserService initializes the user service
func NewUserService(repo UserRepository) UserService {
	return &userService{
		repo:      repo,
		validator: validator.New(),
	}
}

// GetAllUsers retrieves all users with pagination and search
func (s *userService) GetAllUsers(ctx *gin.Context) ([]User, int64, error) {
	page, err := strconv.Atoi(ctx.Query("page"))
	if err != nil || page < 1 {
		page = 1 // Default to page 1
	}

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil || limit < 1 {
		limit = 10 // Default limit
	}

	search := ctx.Query("search")

	users, total, err := s.repo.GetAllUser(page, limit, search)
	if err != nil {
		return nil, 0, err
	}

	// Unset the password for each user
	for i := range users {
		users[i].Password = "*****" // Clear the password field
	}

	return users, total, nil
}

// GetUserByID retrieves a user by ID
func (s *userService) GetUserByID(ctx *gin.Context) (User, error) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return User{}, errors.New("invalid ID")
	}

	user, err := s.repo.GetUserByID(int64(id))
	if err != nil {
		return User{}, err
	}

	// Unset the password before returning the user
	user.Password = "*****" // Clear the password field

	return user, nil
}

// CreateUser handles the creation of a new user
func (s *userService) CreateUser(ctx *gin.Context) (User, error) {
	var user User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return User{}, err
	}

	// Hash the password
	hashedPassword, err := common.HashPassword(user.Password)
	if err != nil {
		return User{}, errors.New("failed to hash password")
	}
	user.Password = hashedPassword // Store hashed password
	// Retrieve the username from context and check if it exists
	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return User{}, errors.New("username not found in context")
	}

	// Set CreatedBy field
	user.CreatedBy = username

	if err := s.validator.Struct(user); err != nil {
		return User{}, err
	}
	if err := s.repo.InsertUser(&user); err != nil {
		return User{}, errors.New("failed to add new user")
	}

	return user, nil
}

// UpdateUser updates an existing user
func (s *userService) UpdateUser(ctx *gin.Context) error {
	var user User

	// Parse the ID from the URL parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the user exists in the database
	existingUser, err := s.repo.GetUserByID(int64(id))
	if err != nil {
		return errors.New("user not found")
	}

	// Bind the JSON request to the user struct
	if err := ctx.ShouldBindJSON(&user); err != nil {
		return errors.New("invalid input data")
	}

	// Hash the password if it's provided (can be optional)
	if user.Password != "" {
		hashedPassword, err := common.HashPassword(user.Password)
		if err != nil {
			return errors.New("failed to hash password")
		}
		user.Password = hashedPassword // Store hashed password
	}

	// Retrieve the username from context and check if it exists
	username, ok := ctx.Value("username").(string)
	if !ok || username == "" {
		return errors.New("username not found in context")
	}

	// Set Modified field
	user.ModifiedAt = time.Now()
	user.ModifiedBy = username
	// Assign the existing user ID to the new user object
	user.ID = existingUser.ID

	// Update the user in the database
	if err := s.repo.UpdateUser(user); err != nil {
		return errors.New("failed to update user")
	}

	return nil
}

// DeleteUser deletes a user
func (s *userService) DeleteUser(ctx *gin.Context) error {
	// Parse the user ID from the URL parameter
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errors.New("invalid ID format")
	}

	// Check if the user exists
	_, err = s.repo.GetUserByID(int64(id))
	if err != nil {
		return errors.New("user not found")
	}

	// Proceed to delete the user
	if err := s.repo.DeleteUser(User{ID: int64(id)}); err != nil {
		return err
	}

	return nil
}
