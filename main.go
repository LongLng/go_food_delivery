package main

import (
	"fooddelivery/component/appctx"
	"fooddelivery/middleware"
	"fooddelivery/module/restaurant/transport/ginrestaurant"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
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
	Id     int    `json:"id" gorm:"column:id"`
	Name   string `json:"name" gorm:"column:name"`
	Addr   string `json:"addr" gorm:"column:addr"`
	CityId int    `json:"city_id" gorm:"column:city_id"`
	//Lat    float64 `json:"lat" gorm:"column:lat"`
	//Lng    float64 `json:"lng" gorm:"column:lng"`
	//Cover  float64 `json:"cover" gorm:"column:cover"`
	//Logo   float64 `json:"logo" gorm:"column:logo"`
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
	log.Println(db)

	r := gin.Default()
	appContext := appctx.NewAppContext(db)

	r.Use(middleware.Recover(appContext))
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	//POST /restaurants
	v1 := r.Group("/api/v1")

	restaurants := v1.Group("/restaurants")
	restaurants.POST("", ginrestaurant.CreateRestaurant(appContext))

	restaurants.DELETE("/:id", ginrestaurant.DeleteRestaurant(appContext))
	r.Run()

}
