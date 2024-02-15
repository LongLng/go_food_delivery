package restaurantstorage

import (
	"context"
	"fooddelivery/common"
	restaurantmodel "fooddelivery/module/restaurant/model"
)

func (s *sqlStore) GetListWithCondition(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {

	var result []restaurantmodel.Restaurant
	db := s.db.Where("status in (1)")

	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("ower_id=?", f.OwnerId)
		}
	}
	if err := db.Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
