package database

import (
	"errors"
	"gorm.io/gorm"
	"users/internal/models"
)

// GetUserByEmail retrieves a user by their email.
func (s *service) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := s.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, err // Handle other errors
	}
	return &user, nil
}

// GetUserById retrieves a user by their ID.
func (s *service) GetUserById(id int64) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // User not found
		}
		return nil, err // Handle other errors
	}
	return &user, nil
}

// CreateUser inserts a new user into the database.
func (s *service) CreateUser(user *models.User) error {
	if err := s.db.Create(user).Error; err != nil {
		return err // Handle insert error
	}
	return nil
}

// UpdateUser updates a user's name (can be extended to update other fields).
func (s *service) UpdateUser(name string) error {
	return s.db.Model(&models.User{}).Where("name = ?", name).Updates(map[string]interface{}{
		"name": name,
	}).Error
}

// DeleteUser removes a user from the database.
func (s *service) DeleteUser(id int64) error {
	if err := s.db.Delete(&models.User{}, id).Error; err != nil {
		return err // Handle delete error
	}
	return nil
}

// VerifyUser marks a user as verified.
func (s *service) VerifyUser(id int64) error {
	if err := s.db.Model(&models.User{}).Where("id = ?", id).Update("is_verified", true).Error; err != nil {
		return err // Handle verification error
	}
	return nil
}

// StoreRefreshToken inserts a new refresh token into the database.
func (s *service) StoreRefreshToken(token *models.RefreshToken) error {
	if err := s.db.Create(token).Error; err != nil {
		return err // Handle insert error
	}
	return nil
}

// GetRefreshToken retrieves a refresh token by its token string.
func (s *service) GetRefreshToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := s.db.Where("token = ?", token).First(&refreshToken).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Token not found
		}
		return nil, err // Handle other errors
	}
	return &refreshToken, nil
}

// UpdateRefreshToken updates the refresh token with a new value.
func (s *service) UpdateRefreshToken(storedToken, newToken string) error {
	return s.db.Model(&models.RefreshToken{}).Where("token = ?", storedToken).Update("token", newToken).Error
}
