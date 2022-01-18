package commons

import (
	"log"

	"github.com/infortech07/crudApi/src/models"
	"github.com/jinzhu/gorm"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/test?charset=utf8")

	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("migrando")
	db.AutoMigrate(&models.Person{})
}
