package database

import (
	"Project_2/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect(connect_to string) {
	DB, err = gorm.Open(postgres.Open(connect_to),&gorm.Config{})
	if err!=nil {
		panic("cannot connect to the database")
	}
	fmt.Println("connected to the database")
}
//migrating users struct to users table
func Migrator()  {
	DB.AutoMigrate(&models.Users{})
	fmt.Println("tables created successfully")
}

//migrating admin struct to admin table
func Migrateadmnm(){
	DB.AutoMigrate(&models.Admins{})
	fmt.Println("admin tables created successfully...")
}

//migrating todos struct to todos table
func Migratetodos()  {
	DB.AutoMigrate(&models.Todos{})
	fmt.Println("todos table created succesfully...")
}