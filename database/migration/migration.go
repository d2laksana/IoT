package migration

import (
	"IoT/database"
	"IoT/model/entity"
	"fmt"
)

func Migrate() {
	err := database.DB.AutoMigrate(&entity.Suhu{})
	if err != nil {
		panic("failed to migrate")
	}
	fmt.Println("Database Migrated")
}
