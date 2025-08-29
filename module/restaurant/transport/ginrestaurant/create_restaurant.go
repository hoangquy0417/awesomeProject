package ginrestaurant

import (
	"awesomeProject/module/common"
	"awesomeProject/module/components/appctx"
	restaurantbiz "awesomeProject/module/restaurant/biz"
	restaurantmodel "awesomeProject/module/restaurant/model"
	restaurantstorage "awesomeProject/module/restaurant/storage"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateRestaurant(appCtx appctx.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data restaurantmodel.RestaurantCreate
		db := appCtx.GetMaiDBConnection()
		if err := c.ShouldBindJSON(&data); err != nil {
			panic(err)
		}
		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewCreateRestaurantBiz(store)
		if err := biz.CreateRestaurant(c.Request.Context(), &data); err != nil {
			panic(err)
		}
		data.Mask(false)
		c.JSON(http.StatusCreated, common.SimpleSuccessResponse(data.FakeId.String()))
	}
}
