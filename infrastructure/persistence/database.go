package persistence

import (
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

func ConnectToDatabase(config *DatabaseConfig) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	// 	config.User,
	// 	config.Password,
	// 	config.Host,
	// 	config.Port,
	// 	config.DbName,
	// )

	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to connect to database: %v", err)
	// 	return nil, err
	// }

	return nil, nil
}
