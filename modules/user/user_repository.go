package user

import (
	"errors"

	"gorm.io/gorm"
)

// UserRepository interface for user operations
type UserRepository interface {
	InsertUser(user *User) error
	GetAllUser(page, limit int, search string) ([]User, int64, error)
	GetUserByID(id int64) (User, error)
	UpdateUser(user User) error
	DeleteUser(user User) error
}

// userRepository struct implementing UserRepository interface
type userRepository struct {
	db *gorm.DB
}

// NewRepository initializes the user repository with a DB connection
func NewRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// InsertUser adds a new user to the database
func (r *userRepository) InsertUser(user *User) error {
	return r.db.Table("user").Create(user).Error
}

// GetAllUser retrieves users from the database with pagination and search
func (r *userRepository) GetAllUser(page, limit int, search string) ([]User, int64, error) {
	var users []User
	var total int64

	query := r.db.Table("user")

	if search != "" {
		query = query.Where("username LIKE ?", "%"+search+"%")
	}

	err := query.Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * limit
	err = query.Limit(limit).Offset(offset).Find(&users).Error
	return users, total, err
}

// GetUserByID retrieves a user by ID from the database
func (r *userRepository) GetUserByID(id int64) (User, error) {
	var user User
	err := r.db.Table("user").Where("id = ?", id).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return User{}, errors.New("user not found")
	}
	return user, err
}

// UpdateUser updates an existing user in the database
func (r *userRepository) UpdateUser(user User) error {
	return r.db.Table("user").Save(&user).Error
}

// DeleteUser deletes a user from the database
func (r *userRepository) DeleteUser(user User) error {
	return r.db.Table("user").Delete(&user).Error
}

// CheckUserExists checks if a user exists by ID
func CheckUserExists(db *gorm.DB, id int64) (bool, error) {
	var user User
	if err := db.Table("user").Where("id = ?", id).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil // User does not exist
		}
		return false, err // An error occurred while querying
	}
	return true, nil // User exists
}
