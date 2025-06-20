package logger

import (
	"fmt"
	"homework7/internal/model"
	"homework7/internal/repository"
	"time"
)

func check(sliceLen, prevLen int, slice any) {
	if sliceLen > prevLen {
		switch items := slice.(type) {
		case []*model.Order:
			for _, o := range items[prevLen:] {
				fmt.Printf("[LOG] В slice добавлен Order: %+v\n", *o)
			}
		case []*model.Product:
			for _, p := range items[prevLen:] {
				fmt.Printf("[LOG] В slice добавлен Product: %+v\n", *p)
			}
		case []*model.Category:
			for _, c := range items[prevLen:] {
				fmt.Printf("[LOG] В slice добавлен Category: %+v\n", *c)
			}
		case []*model.User:
			for _, u := range items[prevLen:] {
				fmt.Printf("[LOG] В slice добавлен User: %+v\n", *u)
			}
		case []*model.Adminka:
			for _, a := range items[prevLen:] {
				fmt.Printf("[LOG] В slice добавлен Adminka: %+v\n", *a)
			}
		}
	}
}

// Start каждые 200 миллисекунд проверяет, сколько элементов в слайсах,
// и выводит только что добавленные
func Start(done <-chan struct{}) {
	ticker := time.NewTicker(200 * time.Millisecond)
	defer ticker.Stop()

	var prevO, prevP, prevC, prevU, prevA int

	for {
		select {
		case <-ticker.C:
			os, ps, cs, us, as := repository.Snapshot()
			check(len(os), prevO, os)
			check(len(ps), prevP, ps)
			check(len(cs), prevC, cs)
			check(len(us), prevU, us)
			check(len(as), prevA, as)
			prevO, prevP, prevC, prevU, prevA = len(os), len(ps), len(cs), len(us), len(as)
		case <-done:
			break
		}
	}
}
