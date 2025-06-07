package main

import (
	"homework7/internal/logger"
	"homework7/internal/model"
	"homework7/internal/repository"
	"homework7/internal/service"
)

func main() {
	dataCh := make(chan model.ID)
	doneCh := make(chan struct{})

	go service.GenerateItems(dataCh)
	go repository.ReceiveData(dataCh, doneCh)
	logger.Start(doneCh)
}
