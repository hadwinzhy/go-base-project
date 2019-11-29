package connection

import (
	"cw-app/hotpot-backend/configs"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// PG is Singleton
var PG *gorm.DB

// InitPGConnection will init db and return singleton instance
func InitPGConnection() *gorm.DB {
	PG, err := gorm.Open("postgres", getPGDBString())
	if err != nil {
		panic("failed to connect database" + err.Error())
	}
	PG.LogMode(true)
	PG.DB().SetMaxIdleConns(3)
	PG.DB().SetMaxOpenConns(10)

	return PG
}

//TODO  write to common database config
func getPGDBString() string {
	dbConfig := configs.DatabaseConfig
	dbString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.DBname,
		dbConfig.SslMode,
	)
	fmt.Println(dbString)
	return dbString
}
