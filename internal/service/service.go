package service

import (
	"homework7/internal/model"
	"homework7/internal/repository"
	"sync"
)

//var (
//	order    []*model.Order
//	product  []*model.Product
//	category []*model.Category
//	user     []*model.User
//	adminka  []*model.Adminka
//	mu       sync.Mutex
//)

//func Createitem() {
//for i :=1; i<6; i++ {
//if i == 1 {
//p := model.Product{ID: 1}
//repository.AddData(p)}
//if i == 2 {
//c := model.Category{ID: 2}
//repository.AddData(c)}
//if i == 3 {
//o := model.Order{ID: 2}
//repository.AddData(o)}
//if i == 4 {
//u := model.User{ID: 3}
//repository.AddData(u)}
//if i == 5 {
//a := model.Adminka{ID: 4}
//repository.AddData(a)}
//}
//}

// GenerateItems отправляет в канал по одному элементу.
func GenerateItems(ch chan<- model.ID, wg *sync.WaitGroup) {
	defer close(ch)
	defer wg.Done()
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

func ReceiveData(ch <-chan model.ID, wg *sync.WaitGroup) {
	//func ReceiveData(ch <-chan model.ID, done chan<- struct{}, wg *sync.WaitGroup) {
	//defer close(done)
	defer wg.Done()
	for v := range ch {
		repository.AddData(v)
	}
}
