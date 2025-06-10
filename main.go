package main

import (
	"homework7/internal/logger"
	"homework7/internal/model"
	"homework7/internal/service"
	"sync"
)

func main() {
	//service.Createitem()
	var wg sync.WaitGroup
	var lwg sync.WaitGroup

	dataCh := make(chan model.ID)
	doneCh := make(chan struct{})
	wg.Add(1)
	go service.GenerateItems(dataCh, &wg)
	wg.Add(1)
	go service.ReceiveData(dataCh, &wg)
	lwg.Add(1)
	go logger.Start(doneCh, &lwg)
	wg.Wait()
	close(doneCh)
	lwg.Wait()
	//lwg.Done()
}
