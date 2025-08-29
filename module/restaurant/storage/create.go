package restaurantstorage

import (
	"awesomeProject/module/common"
	restaurantmodel "awesomeProject/module/restaurant/model"
	"context"
)

// co io thi nen dung context

func (s sqlStore) Create(context context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return common.ErrDB(err)
	}

	return nil
}
