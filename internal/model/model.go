package model

import (
	"time"
)

//# Админка приложения (структура)
//Административный интерфейс позволяет администратору приложения загрузить данные о товарах и категориях из файла или по ссылке, без программирования.
//Поля структуры Админка приложения:
//- Ссылка на файл/feed (string)
//- Файл //Возможность загрузки файла, не знаю какого типа поле

type Adminka struct {
	ID     int64
	Import string
	File   string
}

//# Заказ (структура)
//Поля струтктуры Заказ:
//- имя (string)
//- ID номер заказа (int)
//- количество товара в заказе (int)
//- дата и время (time.Time)
//- цена (int)
//- итоговая сумма (int)
//- наименование товара (string)
//- статус заказа (принят, отправлен, завершен) (string)
//- id пользователя (string)
//- телефон (string)
//- адрес доставки //не знаю получится сделать или нет, тк по уму там надо города списком выводить, подтягивать кладер адресов и городов (string)

type Order struct {
	ID             int64
	Quantity       int64
	DateTime       time.Time
	Price          int64
	Status         string
	DeliveryAdress []string
}

//# Товары (структура)
//Поля структуры Товары:
//- ID товара (int)
//- Поле для привязки товара к категории по ID (int)
//- название товара (string)
//- ссылка на картинку (картинка) //в бд наверное лучше сохранять ссылку на картинку (string)
//- цена (int)
//- количество товара (остаток на складе) (int)
//- описание товара (string)

type Product struct {
	ID         int64
	Name       string
	PictureUrl string
	Price      int64
	Quantity   int64
	Descripion string
}

//# Категория товара (структура)
//Поля структуры категория:
//- название категории (string)
//- ID категории (int)

type Category struct {
	ID        int64
	Name      string
	Categorys map[int64]*Category
	Products  map[int64]*Product
}

//# Структура пользователь
//Поля структуры Пользователь:
//- ID пользователя (int)
//- Имя (string)
//- Телефон (string)

type User struct {
	ID    int64
	name  string
	phone []string
}

//type Orders interface {
//	Add() string
//}

type ID interface {
	GetID() int64
}

//func (a Adminka) Add() string {
//	// return "I am the Adminka"
//	return "Админка" + a.File + a.Import
//}

func (o Order) GetID() int64 {
	// return "I am the Order stuct"
	return o.ID
}

func (p Product) GetID() int64 {
	// return "I am the Product sruct"
	return p.ID
}

func (c Category) GetID() int64 {
	// return "I am the Category sruct"
	return c.ID
}

func (u User) GetID() int64 {
	// return "I am the User sruct"
	return u.ID
}

func (a Adminka) GetID() int64 {
	// return "I am the Adminka sruct"
	return a.ID
}
