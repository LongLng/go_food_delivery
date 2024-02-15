package restaurantbiz

import (
	"context"
	"fooddelivery/common"
	restaurantmodel "fooddelivery/module/restaurant/model"
)

type DeleteRestaurantStore interface {
	Get(
		context context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	store DeleteRestaurantStore
}

func NewDeleteRestaurantBiz(store DeleteRestaurantStore) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{store: store}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.store.Get(context, map[string]interface{}{"id": id})

	if err != nil {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}

	if oldData.Status == 0 {
		return common.ErrEntityDeleted(restaurantmodel.EntityName, nil)
	}
	if err := biz.store.Delete(context, id); err != nil {
		return common.ErrCannotDeleteEntity(restaurantmodel.EntityName, nil)
	}
	return nil
}
