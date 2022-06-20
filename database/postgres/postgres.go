package postgres

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"three/utils"
	userdomain "three/v1/user/domain"

	"gorm.io/gorm"
)

func InitDatabase(config utils.PostgresConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", config.DbHost, config.DbUser, config.DbPassword, config.DbName, config.DbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&userdomain.User{})
	if err != nil {
		return nil, err
	}
	err = createUserData(db)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func createUserData(db *gorm.DB) error {
	var userData userdomain.User

	dbRes := db.Find(&userData)

	if dbRes.RowsAffected > 0 {
		return nil
	}

	userPassword, err := bcrypt.GenerateFromPassword([]byte("inibebek"), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userData.Username = "bebek"
	userData.Password = string(userPassword)

	db.Create(&userData)
	return nil
}
