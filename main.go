package main

import (
	"fmt"
	"homework7/internal/logger"
	"homework7/internal/model"
	"homework7/internal/service"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	//service.Createitem()
	var wg sync.WaitGroup
	//var lwg sync.WaitGroup

	dataCh := make(chan model.ID)
	doneCh := make(chan struct{})

	//wg.Add(1)
	//go func() {
	//	defer wg.Done()
	//	service.GenerateItems(dataCh)
	//}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		service.ReceiveData(dataCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Start(doneCh)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		service.StartGeneration(doneCh, dataCh)
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	sig := <-sigChan
	fmt.Printf("Получили сигнал: %v. Заканчиваю работу...\n", sig)
	close(doneCh)
	wg.Wait()
}
