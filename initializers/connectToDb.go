package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
    // connect to the database
    var err error
    dsn := os.Getenv("DB")
    DB , err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    fmt.Print()
    if err != nil {
        panic("Failed to connect to database!")
    }

    fmt.Println(DB)

}