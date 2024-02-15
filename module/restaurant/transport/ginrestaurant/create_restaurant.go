package ginrestaurant

import (
	"fooddelivery/common"
	"fooddelivery/component/appctx"
	restaurantbiz "fooddelivery/module/restaurant/biz"
	restaurantmodel "fooddelivery/module/restaurant/model"
	restaurantstorage "fooddelivery/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		db := appCtx.GetMailDBConnection()
		var data restaurantmodel.RestaurantCreate
		if err := c.ShouldBind(&data); err != nil {
			panic(err)
		}

		store := restaurantstorage.NewSqlStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSucessResponse(data.Id))
	}

}
