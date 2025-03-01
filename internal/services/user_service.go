package services

import (
	"cerberus/internal/dto/auth_dto"
	"cerberus/internal/dto/session_dto"
	"cerberus/internal/models"
	"cerberus/internal/repository"
	logger "cerberus/internal/tools/logger"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
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
	usr, err := repository.FindUserByEmail(p_dg, p_email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	} else if err != nil {
		logger.Log(fmt.Sprintf("Something went wrong - %s", err.Error()), logger.ERROR)
		return false, err
	}

	return usr != nil, nil
}

// RegisterUser creates a new user account in the database.
//
// Parameters:
//   - p_dg: A pointer to the GORM database connection.
//   - p_email: The email address of the user to be registered.
//   - p_pwd: The password for the new user account.
//   - p_name: The name of the user to be registered.
//
// Returns:
//   - A pointer to the newly created User object if registration is successful.
//   - An error if registration fails (e.g., duplicate email, password hashing error, or database error).
//
// If any step fails, an appropriate error is logged and returned.
func RegisterUser(p_dg *gorm.DB, p_register_dto *auth_dto.RegisterRequest) (*models.User, error) {
	r, _ := IsUserRegistered(p_dg, p_register_dto.Email)
	if r {
		msg := fmt.Sprintf("Failed to registered, duplication - %s", p_register_dto.Email)
		logger.Log(msg, logger.ERROR)
		return nil, errors.New(msg)
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(p_register_dto.Password), bcrypt.DefaultCost)
	if err != nil {
		logger.Log("Failed to hash password - "+err.Error(), logger.ERROR)
		return nil, err
	}

	var user *models.User = &models.User{
		Name:     p_register_dto.Name,
		Email:    p_register_dto.Email,
		Password: string(hashedPwd),
	}

	err = repository.CreateUser(p_dg, user)
	if err != nil {
		logger.Log("Failed to Create user - "+err.Error(), logger.ERROR)
		return nil, err
	}

	return user, nil
}

// ChangePassword updates a user's password in the database.
//
// Parameters:
//   - p_db: A pointer to a gorm.DB instance representing the database connection.
//   - p_change_pwd_dto: A pointer to an auth_dto.ChangePasswordRequest struct containing:
//   - Email: The user's email address
//   - CurrentPassword: The user's current password
//   - NewPassword: The desired new password
//
// Returns:
//   - error: An error if any step fails, or nil if the password change is successful.
//     Possible error messages include:
//   - "user not found"
//   - "invalid password"
//   - "password must be different"
//   - "failed to hash the new password - [error details]"
//   - "failed to update password - [error details]"
//
// Note: This function uses bcrypt for password hashing and comparison.
func ChangePassword(p_db *gorm.DB, p_change_pwd_dto *auth_dto.ChangePasswordRequest) error {
	usr, err := repository.FindUserByEmail(p_db, p_change_pwd_dto.Email)
	if err != nil {
		return errors.New("user not found")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(p_change_pwd_dto.CurrentPassword)); err != nil {
		return errors.New("invalid password")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(p_change_pwd_dto.NewPassword)); err == nil {
		return errors.New("password must be different")
	}

	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(p_change_pwd_dto.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Log("Failed to hash password - "+err.Error(), logger.ERROR)
		return errors.New("failed to hash the new password - " + err.Error())
	}

	err = repository.UpdatePassword(p_db, usr, string(hashedPwd))
	if err != nil {
		logger.Log("Failed to update password - "+err.Error(), logger.ERROR)
		return errors.New("failed to update password - " + err.Error())
	}

	return nil
}

// LoginUserPasswordCheck verifies a user's login credentials against the database.
//
// Parameters:
//   - p_db: A pointer to the gorm.DB instance for database operations.
//   - p_login_dto: A pointer to the LoginRequest containing the user's email and password.
//
// Returns:
//   - *postgres_models.User: A pointer to the User model if authentication is successful.
//   - error: An error if authentication fails, nil otherwise.
//
// TODO: Maybe set a group of generic dtos, packs that will be shared between services
func LoginUserPasswordCheck(p_db *gorm.DB, p_login_dto *session_dto.LoginRequest) (*models.User, error) {
	usr, err := repository.FindUserByEmail(p_db, p_login_dto.Email)
	if err != nil {
		logger.Log("User not found - "+err.Error(), logger.ERROR)
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(p_login_dto.Password)); err != nil {
		logger.Log("Invalid credentials - "+err.Error(), logger.ERROR)
		return nil, err
	}

	return usr, nil
}
