package model

import (
	"time"
)

type Adminka struct {
	ID     int64
	Import string
	File   string
}

type Order struct {
	ID             int64
	Quantity       int64
	DateTime       time.Time
	Price          int64
	Status         string
	DeliveryAdress []string
}

type Product struct {
	ID         int64
	Name       string
	PictureUrl string
	Price      int64
	Quantity   int64
	Descripion string
}

type Category struct {
	ID        int64
	Name      string
	Categorys map[int64]*Category
	Products  map[int64]*Product
}

type User struct {
	ID    int64
	name  string
	phone []string
}

type ID interface {
	GetID() int64
}

func (o Order) GetID() int64 {
	return o.ID
}

func (p Product) GetID() int64 {
	return p.ID
}

func (c Category) GetID() int64 {
	return c.ID
}

func (u User) GetID() int64 {
	return u.ID
}

func (a Adminka) GetID() int64 {
	return a.ID
}
