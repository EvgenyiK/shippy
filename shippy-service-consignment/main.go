package main

import (
	"context"
	"log"
	

	//Импортировать сгенерированный код protobuf
	pb "github.com/EvgenyiK/shippy/shippy-service-consignment/proto/consignment"
	   "github.com/micro/go-micro/v2"
)


type repository interface {
	Create(*pb.Consignment) (*pb.Consignment, error)
	GetAll() []*pb.Consignment
}

//Репозиторий имитирующий использование хранилища данных
type Repository struct {
	consignments []*pb.Consignment
}

// Создать новую партию
func (repo *Repository) Create(consignment *pb.Consignment) (*pb.Consignment, error) {
	updated := append(repo.consignments, consignment)
	repo.consignments = updated
	return consignment, nil
}

//покажи все партии
func (repo *Repository) GetAll() []*pb.Consignment {
	return repo.consignments
}

//Сервис должен реализовывать все методы
type consignmentService struct {
	repo repository
}

//CreateConsignment метод в сервисе , который является методом create
//обрабатываются сервером gRPC
func (s *consignmentService) CreateConsignment(ctx context.Context,
	req *pb.Consignment, res *pb.Response) error {
	//Сохраним партию
	consignment, err := s.repo.Create(req)
	if err != nil {
		return  err
	}

	//Возврат ответа на сообщение `Response`, которое мы создали
	// в нашем protobuf
	res.Created = true
	res.Consignment = consignment
	return nil
}

//GetConsignments
func (s *consignmentService) GetConsignments(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {
	consignments := s.repo.GetAll()
	res.Consigments = consignments
	return nil
}

func main() {
	repo := &Repository{}
	//Создаем новый сервис
	service:=micro.NewService(
		//Это имя должно соответствовать имени пакета, указанному в вашем определении protobuf
		micro.Name("shippy.service.consignment"),
	)
	// Проинициализируем командную строку
	service.Init()

	//Регистрация сервиса
	if err:=pb.RegisterShippingServiceHandler(service.Server(), 
				&consignmentService{repo});err!=nil{
					log.Panic(err)
				}
	
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
