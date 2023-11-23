package config

import (
	"server/entity"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var database *gorm.DB
var e error

func DatabaseInit() *gorm.DB {
	// host := os.Getenv("DB_HOST")
	// user := os.Getenv("DB_USER")
	// password := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// port, err := strconv.Atoi(os.Getenv("DB_PORT"))

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	dsn := Env().DB_DSN
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Jakarta", host, user, password, dbName, port)
	database, e = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if e != nil {
		panic(e)
	}

	database.AutoMigrate(&entity.Admin{}, &entity.Experience{}, &entity.Project{}, &entity.Skill{}, &entity.User{})

	return database
}

func DB() *gorm.DB {

	dbGorm, err := database.DB()
	if err != nil {
		panic(err)
	}
	dbGorm.Ping()

	return database
}
