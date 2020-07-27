package main

import (
	"context"
	"log"
	"os"

	pb "github.com/EvgenyiK/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")

	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{vesselCollection}

	h := &handler{repository}

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), h); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}

/*
Здесь мы создали экземпляр клиента для нашего сервиса судна, что позволяет нам использовать название сервиса, 
то есть shipy.service.vesselвызывать сервис судна как клиент и взаимодействовать с его методами. 
В этом случае только один метод ( FindAvailable). 
Мы отправляем вес отправления вместе с количеством контейнеров, которые мы хотим отправить,
 в качестве спецификации в судовую службу. Который затем возвращает соответствующее судно.
*/