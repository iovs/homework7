package service

import (
	"homework6/internal/model"
	"homework6/internal/repository"
//	"math/rand"
//	"time"
)

func Createitem() {
//	rand.NewSource(time.Now().UnixNano())
//	if rand.Intn(5) == 0 {
//		//	d := model.Adminka{File: "file text"}
//		p := model.Product{ID: 0}
//		repository.AddData(p)
//	} else if rand.Intn(5) == 1 {
//		c := model.Category{ID: 1}
//		repository.AddData(c)
//	} else if rand.Intn(5) == 2 {
//		o := model.Order{ID: 2}
//		repository.AddData(o)
//	} else if rand.Intn(5) == 3 {
//		u := model.User{ID: 3}
//		repository.AddData(u)
//	} else if rand.Intn(5) == 4 {
//		a := model.Adminka{ID: 4}
//		repository.AddData(a)
//	}
for i :=1; i<6; i++ {
	if i == 1 {
	p := model.Product{ID: 1}
	repository.AddData(p)}
    if i == 2 {
	c := model.Category{ID: 2}
    repository.AddData(c)}
	if i == 3 {
    o := model.Order{ID: 2}
    repository.AddData(o)}
	if i == 4 {
    u := model.User{ID: 3}
    repository.AddData(u)}
	if i == 5 {
		a := model.Adminka{ID: 4}
    repository.AddData(a)}
}
}

//type Orders interface {
//	Add() string
//}

//func Func1() {
//	Adminka := model.Adminka{}
//	add(Adminka)
//}

//func add(orders Orders) {
//	fmt.Println(orders.Add())
//}

//type Adminka struct {
//	Import string
//	File   string
//}
//type Adminka interface {
//	Savefile() string
//}
//type Order struct {
//	ID             int64
//	Quantity       int64
//	DateTime       time.Time
//	Price          int64
//	Status         string
//	DeliveryAdress []string
//}

//type Orders interface {
//	Addorder() string
//}

//func NewImport(Import, File string){
//	var NewFunc1 = (aaaa)
//	fmt.Println(NewFunc1)
//}

//var Adminka []model.Adminka //? model?
//var Orders []model.Order
//var Product []model.Product
//var Category []model.Category
//var Users []model.User
//func Createitem1 (Import, File Adminka interface{Import; File; string}) //Где определяется Createitem?
//func Createitem2 (context, Orders interface{})
//func Createitem3 (context, Product interface{})
//func Createitem4 (context, Category interface{})
//func Createitem5 (context, User interface{})

//{
//проверяем тип item
//добавляем в слайс
//}
//инициализировать слайс проверить создан или нет

//type Order struct {
//	ID             int64
//	Quantity       int64
//	DateTime       time.Time
//	Price          int64
//	Status         string
//	DeliveryAdress []string
//}

//func NewOrder(ID, Quantity, Price int64 DateTime time.Time, Status string, DeliveryAdress []string)

//	type Product struct {
//		ID         int64
//		Name       string
//		PictureUrl string
//		Price      int64
//		Quantity   int64
//		Descripion string
//	}
//func AddProduct(ID, Price, Quantity int64, Name, PictureUrl, Description string) {
//	var newFunc= {1111, 200, 3, "Ivan", "с:/picture", "Описание"}
//	fmt.Println(newFunc)
//}

//	type User struct {
//		ID    int64
//		name  []string
//		phone []string
//	}
//func AddUser(ID int64, name, phone string)
