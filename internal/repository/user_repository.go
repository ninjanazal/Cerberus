package repository

import (
	"cerberus/internal/models"

	"gorm.io/gorm"
)

// region Public

// FindUserByEmail retrieves a user from the database by their email address.
//
// Parameters:
//   - p_db: A gorm.DB instance representing the database connection.
//   - p_email: A string containing the email address to search for.
//
// Returns:
//   - *postgres_models.User: A pointer to the User struct if found, or nil if not found.
//   - error: An error if the database query fails, or nil if successful.
func FindUserByEmail(p_db *gorm.DB, p_email string) (*models.User, error) {
	var u models.User
	res := p_db.Where("email = ?", p_email).First(&u)
	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}

// FindUserById retrieves a user from the database by their unique ID.
// It uses the provided GORM database connection to query the user table
// and returns the corresponding user model if found.
//
// Parameters:
//   - p_db: A pointer to the GORM database connection (*gorm.DB) used to execute the query.
//   - p_usrId: The unique ID (string) of the user to be retrieved.
//
// Returns:
//   - *models.User: A pointer to the user model if the user is found; otherwise, nil.
//   - error: An error object if the query fails or the user is not found; otherwise, nil.
func FindUserById(p_db *gorm.DB, p_usrId string) (*models.User, error) {
	var u models.User
	res := p_db.Where("id = ?", p_usrId).First(&u)
	if res.Error != nil {
		return nil, res.Error
	}
	return &u, nil
}

// CreateUser creates a new user record in the database.
//
// Parameters:
//   - p_db: A pointer to the gorm.DB database connection.
//   - p_user: A pointer to the postgres_models.User struct containing the user data to be inserted.
//
// Returns:
//   - error: An error if the creation fails, or nil if successful.
func CreateUser(p_db *gorm.DB, p_user *models.User) error {
	return p_db.Create(p_user).Error
}

// UpdatePassword updates the password of a user in the database.
// It uses the provided GORM database connection to update the "password" field
// for the user identified by their ID.
//
// Parameters:
//   - p_db: A pointer to the GORM database connection (*gorm.DB) used to execute the query.
//   - p_user: A pointer to the user model (*models.User) containing the user's ID.
//   - p_pwd: The new password (string) to be set for the user.
//
// Returns:
//   - error: An error object if the update operation fails; otherwise, nil.
func UpdatePassword(p_db *gorm.DB, p_user *models.User, p_pwd string) error {
	return p_db.Model(&models.User{}).Where("id = ?", p_user.ID).Update("password", p_pwd).Error
}

// endregion Public
