package connection

import "github.com/jinzhu/gorm"

// PG is Singleton
var PG *gorm.DB

// InitDBConn will init db and return singleton instance
func InitDBConn() *gorm.DB {
	PG, err := gorm.Open("postgres", getDBSettings())
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	PG.LogMode(true)
	PG.DB().SetMaxIdleConns(3)
	PG.DB().SetMaxOpenConns(10)

	return PG
}

//TODO  write to common database config
func getDBSettings() string {
	return ""
}
