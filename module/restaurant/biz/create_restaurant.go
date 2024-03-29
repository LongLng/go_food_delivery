package restaurantbiz

import (
	"context"
	"fooddelivery/common"
	restaurantmodel "fooddelivery/module/restaurant/model"
)

type CreateRestaurantStore interface {
	Create(context context.Context, data *restaurantmodel.RestaurantCreate) error
}

type createRestaurantBiz struct {
	store CreateRestaurantStore
}

func NewCreateRestaurantBiz(store CreateRestaurantStore) *createRestaurantBiz {
	return &createRestaurantBiz{store: store}
}

func (biz *createRestaurantBiz) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {

	if err := data.Validate(); err != nil {
		return common.ErrInvalidRequest(err)
	}

	//Logic Business
	if err := biz.store.Create(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(restaurantmodel.EntityName, err)
	}

	return nil
}
