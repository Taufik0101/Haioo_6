package injection

import (
	"Haioo_6/api/model"
	"Haioo_6/api/utils"
	"gorm.io/gorm"
	"log"
)

type Migration interface {
	Migrate()
}

type migrationConnection struct {
	connection *gorm.DB
}

func (m migrationConnection) Migrate() {
	if utils.EnvVar("APP_ENV", "DEVELOPMENT") == "DEVELOPMENT" {
		m.connection.Exec("CREATE DATABASE IF NOT EXISTS")

		if m.connection.Migrator().HasTable(&model.Cart{}) {
			errUser := m.connection.Migrator().DropTable(&model.Cart{})
			if errUser != nil {
				log.Println("Failed To Drop Table User")
			}
		}

		err := m.connection.AutoMigrate(
			&model.Cart{},
		)

		if err != nil {
			log.Println("Failed To Migrate Table")
		}
	}
}

func NewMigration(conn *gorm.DB) Migration {
	return &migrationConnection{
		connection: conn,
	}
}
