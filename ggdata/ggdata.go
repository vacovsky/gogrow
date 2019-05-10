package ggdata

import (
	"log"
	"os"

	"github.com/vacovsky/gogrowggmodels"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// ORM is the object holding the active database connection interface
var ORM *gorm.DB

func init() {
	log.Println("ggdata initializing")
	var err error
	ORM, err = gorm.Open(os.Getenv("GG_DB_DIALECT"), os.Getenv("GG_DB_PATH"))

	if err != nil {
		log.Println(err)
	}

	ORM.LogMode(os.Getenv("GG_DB_LOGMODE") == "ON")
}

// MigrateDataSchema auto-migrates structs
func MigrateDataSchema() {

	ORM.AutoMigrate(ggmodels.DeviceState{})
	ORM.AutoMigrate(ggmodels.Light{})
	ORM.AutoMigrate(ggmodels.Fan{})
	ORM.AutoMigrate(ggmodels.Environment{})
	ORM.AutoMigrate(ggmodels.Cycle{})
	ORM.AutoMigrate(ggmodels.Weather{})
	ORM.AutoMigrate(ggmodels.Relay{})
	ORM.AutoMigrate(ggmodels.FertilizationEvent{})
	ORM.AutoMigrate(ggmodels.WateringEvent{})
	ORM.AutoMigrate(ggmodels.Filter{})

}

func Service() *gorm.DB {
	return ORM
}
