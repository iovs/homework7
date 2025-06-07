package service

import (
	"homework7/internal/model"
)

// GenerateItems отправляет в канал по одному элементу.
func GenerateItems(ch chan<- model.ID) {
	defer close(ch)
	for i := 1; i < 6; i++ {
		switch i {
		case 1:
			ch <- model.Product{ID: 1}
		case 2:
			ch <- model.Category{ID: 2}
		case 3:
			ch <- model.Order{ID: 3}
		case 4:
			ch <- model.User{ID: 4}
		case 5:
			ch <- model.Adminka{ID: 5}
		}
	}
}
