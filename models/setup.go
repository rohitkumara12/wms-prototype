package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectedDB() {

	conectionString := "host=localhost user=postgres password=Sempakbau12 dbname=wms port=5432 sslmode=disable TimeZone=Asia/Jakarta"

	db, err := gorm.Open(postgres.Open(conectionString), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("failed connected database : %s", err.Error()))
	}
	var migrations []error

	migrations = append(migrations, db.AutoMigrate(&User{}))
	migrations = append(migrations, db.AutoMigrate(&Product{}))
	migrations = append(migrations, db.AutoMigrate(&Order{}))
	migrations = append(migrations, db.AutoMigrate(&OrderItem{}))
	migrations = append(migrations, db.AutoMigrate(&StockMovement{}))
	migrations = append(migrations, db.AutoMigrate(
		&Wave{},
		&WaveItem{},
	))
	migrations = append(migrations, db.AutoMigrate(&WaveOrder{}))

	for _, err := range migrations {
		if err != nil {
			panic(fmt.Sprintf("failed imigrations : %s", err.Error()))
		}
	}

	DB = db

	fmt.Println("Database Connected Successfully")
}
