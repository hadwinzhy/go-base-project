package configs

import "fmt"

// DBConfiguration like postgres or mysql config
type DBConfiguration struct {
	Driver   string
	DBname   string
	Username string
	Password string
	Host     string
	Port     string
}

func setDBConfig(baseKey string, db *DBConfiguration) {
	fmt.Println(readYAMLStringValue(baseKey + ".driver"))
	fmt.Println(readYAMLStringValue("env"))
	fmt.Println(ENV)

	// defaultDB := &DBConfiguration{
	// 	Driver:   "",
	// 	DBname:   "",
	// 	Username: "",
	// 	Password: "",
	// 	Host:     "",
	// 	Port:     "",
	// }
}
