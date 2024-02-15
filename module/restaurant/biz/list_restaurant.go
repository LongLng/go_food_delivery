package restaurantbiz

import (
	"context"
	"fooddelivery/common"
	restaurantmodel "fooddelivery/module/restaurant/model"
)

type ListRestaurantStore interface {
	GetListWithCondition(
		context context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging) ([]restaurantmodel.Restaurant, error)
}

type listRestaurantBiz struct {
	store ListRestaurantStore
}

func NewListRestaurantBiz(store ListRestaurantStore) *listRestaurantBiz {
	return &listRestaurantBiz{store: store}
}

func (biz *listRestaurantBiz) ListRestaurant(
	context context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
) ([]restaurantmodel.Restaurant, error) {

	//Logic Business
	result, err := biz.store.GetListWithCondition(context, filter, paging)
	if err != nil {
		return nil, err
	}

	return result, nil
}
