package initializers

import "github.com/tpSpace/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}