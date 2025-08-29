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

func ListRestaurant(appCtx appctx.AppContext) func(c *gin.Context) {
	return func(c *gin.Context) {
		db := appCtx.GetMaiDBConnection()
		var pagingData common.Paging
		if err := c.ShouldBindQuery(&pagingData); err != nil {
			panic(common.ErrInvalidRequest(err))
		}
		pagingData.Fulfill()
		var filter restaurantmodel.Filter
		if err := c.ShouldBindQuery(&filter); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		store := restaurantstorage.NewSQLStore(db)
		biz := restaurantbiz.NewListRestaurantBiz(store)
		result, err := biz.ListRestaurant(c.Request.Context(), &filter, &pagingData)
		if err != nil {
			panic(err)
		}
		for i := range result {
			result[i].Mask(false)
		}
		c.JSON(http.StatusOK, common.NewSuccessResponse(result, pagingData, filter))
	}
}
