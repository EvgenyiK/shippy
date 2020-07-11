package main

import (
	"context"
	"log"
	"os"

	//Импортировать сгенерированный код protobuf
	pb "github.com/EvgenyiK/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/EvgenyiK/shippy/shippy-service-vessel/proto/"
	"github.com/micro/go-micro/v2"
)


const(
	defaultHost = "datastore:27017"
)

func main() {
	//Создаем новый сервис
	service:=micro.NewService(
		//Это имя должно соответствовать имени пакета, указанному в вашем определении protobuf
		micro.Name("shippy.service.consignment"),
	)
	// Проинициализируем командную строку
	service.Init()

	uri:=os.Getenv("DB_HOST")
	if uri = "" {
		uri = defaultHost
	}

	client,err:= CreateClient(context.Background(), uri, 0)
	if err!=nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection:= client.Database("shippy").Collection("Consigments")

	repository:=&MongoRepository{consignmentCollection}
	vesselClient:= vesselProto.NewVesselService("shippy.service.client", service.Client())
	h:=&handler{repository, vesselClient}

	//Регистрируем обработчики
	pb.RegisterShippingServiceHandler(service.Server(), h)
	
	//Запускаем сервер
	if err:=service.Run();err!=nil {
		log.Panic(err)
	}
	//Теперь нам нужно передать переменную окружения при запуске нашего контейнера, 
	//чтобы определить, на каком порту работать.
	/*
	docker run -p 50051:50051 \
      -e MICRO_SERVER_ADDRESS=:50051 \
      shippy-service-consignment
	*/			
}
