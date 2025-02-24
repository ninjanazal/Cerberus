package user_service

import (
	user_repository "cerberus/internal/database/postgresSQL/repository"
	logger "cerberus/internal/tools"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// IsUserRegistered checks if a user with the given email is already registered in the database.
//
// Parameters:
//   - p_dg: A pointer to a gorm.DB instance representing the database connection.
//   - p_email: A string containing the email address to check.
//
// Returns:
//   - bool: true if the user is registered, false otherwise.
//   - error: An error if the database query fails, or nil if successful.
func IsUserRegistered(p_dg *gorm.DB, p_email string) (bool, error) {
	usr, err := user_repository.FindUserByEmail(*p_dg, p_email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		logger.Log(fmt.Sprintf("Something went wrong - %s", err.Error()), logger.ERROR)
		return false, err
	}

	return usr != nil, nil
}
