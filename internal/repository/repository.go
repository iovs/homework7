package repository

import (
	"fmt"
	"homework6/internal/model"
)

// var adminka []model.Adminka
var order []*model.Order
var product []*model.Product
var category []*model.Category
var user []*model.User
var adminka []*model.Adminka

// func AddData(data model.Orders) {
func AddData(data model.ID) {
	switch v := data.(type) {
	case model.Order:
		order = append(order, &v)
		fmt.Println("add Order", v)
	case model.Product:
		product = append(product, &v)
		fmt.Println("add Product", v)
	case model.Category:
		category = append(category, &v)
		fmt.Println("add Category", v)
	case model.User:
		user = append(user, &v)
		fmt.Println("add User", v)
	case model.Adminka:
		adminka = append(adminka, &v)
		fmt.Println("add Adminka", v)
	default:
		fmt.Printf("Неопределенный тип данных: %T/n", v)
	}

}

//var OrderSlice []model.Order // слайс структур model.Order

// Описание подходящего нам интерфейса
//type ID interface {
//	GetID() int64
//}

// функция принимает любые структуры, если они имплементируют (предоставляют) метод GetID() int64
//func Store(input ID) {
//	switch typed := input.(type) { // с помощью .(type) получаем реальный тип input, по сути снимаем мантию интерфейса

//	case model.Order:
//
// append в нужный слайс
//
//	OrderSlice = append(OrderSlice, typed)
//
// и так далее
//
//	}
//}
