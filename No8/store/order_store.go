package store

import (
	"learn-go/No8/model"
)

func CreatOrder(user int64, title string, price float64) *model.Order {
	o := &model.Order{
		UserID: user,
		Title:  title,
		Price:  price,
	}

	return o
}

func GetOrderByUser(user int64) []*model.Order {

	var result []*model.Order
	//TODO: query DB get order list
	return result
}
