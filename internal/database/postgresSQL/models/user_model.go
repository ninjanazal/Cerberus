package postgres_models

import "time"

// User represents a user entity in the application.
//
// This struct is used for object-relational mapping (ORM) with GORM,
// defining the structure and constraints of the user data in the database.
//
// Fields:
//
//	ID: Unique identifier for the user (primary key in the database).
//	Name: The user's name (cannot be null).
//	Email: The user's email address (must be unique and cannot be null).
//	Password: The user's hashed password (cannot be null).
//	CreatedAt: Timestamp of when the user account was created (automatically set).
//
// GORM Tags:
//
//	"primaryKey": Designates the field as the primary key.
//	"not null": Specifies that the field cannot contain a null value.
//	"unique": Ensures the field value is unique across all records.
//	"autoCreateTime": Automatically sets the time when the record is created.
type User struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string    `gorm:"not null"`
	Email     string    `gorm:"unique;not null"`
	Password  string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
