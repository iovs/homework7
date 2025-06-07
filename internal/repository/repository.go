package repository

import (
	"fmt"
	"homework7/internal/model"
	"sync"
)

var (
	order    []*model.Order
	product  []*model.Product
	category []*model.Category
	user     []*model.User
	adminka  []*model.Adminka
	mu       sync.Mutex
)

// ReceiveData читает из канала и добавляет в соответствующий слайс.
func ReceiveData(ch <-chan model.ID, done chan<- struct{}) {
	defer close(done)
	for v := range ch {
		switch elm := v.(type) {
		case model.Order:
			order = append(order, &elm)
			fmt.Println("Добавлен Order: ", elm)
		case model.Product:
			product = append(product, &elm)
			fmt.Println("Добавлен Product: ", elm)
		case model.Category:
			category = append(category, &elm)
			fmt.Println("Добавлен Category: ", elm)
		case model.User:
			user = append(user, &elm)
			fmt.Println("Добавлен User: ", elm)
		case model.Adminka:
			adminka = append(adminka, &elm)
			fmt.Println("Добавлен Adminka: ", elm)
		default:
			fmt.Println("Неизвестный тип: ", elm)
		}
	}
}

// Snapshot возвращает слайсы для безопасного чтения.
func Snapshot() ([]*model.Order, []*model.Product, []*model.Category, []*model.User, []*model.Adminka) {
	mu.Lock()
	defer mu.Unlock()
	os := append([]*model.Order(nil), order...)
	ps := append([]*model.Product(nil), product...)
	cs := append([]*model.Category(nil), category...)
	us := append([]*model.User(nil), user...)
	as := append([]*model.Adminka(nil), adminka...)
	return os, ps, cs, us, as
}
