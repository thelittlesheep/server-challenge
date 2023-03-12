package models

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB create a new instance of GORM
func DB() (*gorm.DB, error) {
	var myEnv map[string]string
	myEnv, _ = godotenv.Read()
	// read environment variables
	dsn := myEnv["DATABASE_DRIVER"] + "://" +
		myEnv["DATABASE_USERNAME"] + ":" +
		myEnv["DATABASE_PASSWORD"] + "@" +
		myEnv["DATABASE_HOST"] + ":" +
		myEnv["DATABASE_PORT"] + "/" +
		myEnv["DATABASE_NAME"]
	// open a new database connection
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
	// return the GORM instance
	return db, nil
}
