package injection

import (
	"Haioo_6/api/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Gagal untuk keluar koneksi database")
	}
	_ = dbSQL.Close()
}

func CreateDatabase() *gorm.DB {
	var err error
	dbHost := utils.EnvVar("DB_HOST", "")
	dbPort := utils.EnvVar("DB_PORT", "")
	dbDatabase := utils.EnvVar("DB_DATABASE", "")
	dbUsername := utils.EnvVar("DB_USERNAME", "")
	dbPassword := utils.EnvVar("DB_PASSWORD", "")

	dataSourceName := dbUsername + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbDatabase + "?parseTime=True"
	dialect := mysql.Open(dataSourceName)

	db, err := gorm.Open(dialect, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		fmt.Println(err)
		panic(any("Failed to connect database!"))
	}

	return db
}
