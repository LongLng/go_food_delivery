package main

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// id                  int auto_increment
// primary key,
// owner_id            int                                 null,
// name                varchar(50)                         not null,
// addr                varchar(255)                        not null,
// city_id             int                                 null,
// lat                 double                              null,
// lng                 double                              null,
// cover               json                                null,
// logo                json                                null,
// shipping_fee_per_km double    default 0                 null,
// status              int       default 1                 not null,
// created_at          timestamp default CURRENT_TIMESTAMP null,
// updated_at          timestamp default CURRENT_TIMESTAMP null on update CURRENT_TIMESTAMP
type Restaurant struct {
	Id   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
	Addr string `json:"addr" gorm:"column:addr"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}
func main() {
	err_env := godotenv.Load(".env")
	if err_env != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := os.Getenv("MYSQL_CONN_STRING")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	newRestaurant := Restaurant{Name: "Long", Addr: "32 Do Duc Duc"}
	db.Create(&newRestaurant)
}
