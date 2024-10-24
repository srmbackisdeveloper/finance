package database

import (
	"fmt"
	"log"
	"os"
	"users/internal/models"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Service interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserById(id int64) (*models.User, error)

	CreateUser(*models.User) error
	UpdateUser(name string) error
	DeleteUser(id int64) error
	VerifyUser(id int64) error

	StoreRefreshToken(token *models.RefreshToken) error
	GetRefreshToken(token string) (*models.RefreshToken, error)
	UpdateRefreshToken(storedToken, refreshToken string) error
}

type service struct {
	db *gorm.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	sslMode  = os.Getenv("DB_SSL")
)

func New() Service {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, database, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	err = db.AutoMigrate(&models.User{}, &models.RefreshToken{})
	return &service{db: db}
}
