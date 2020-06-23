package main

import (
	"context"
	"log"
	"net"
	"shippy/shippy-service-consignment/proto/consignment"
	"sync"

	//Импортировать сгенерированный код protobuf
	pb "github.com/EvgenyiK/shippy/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


const (
	port=":50051"
)

type repository interface{
	Create(*pb.Consignment) (*pb.Consignment, error)
}

//Репозиторий имитирующий использование хранилища данных
type Repository struct{
	mu sync.RWMutex
	consignments []*pb.Consignment
}

// Создать новую партию
func (repo *Repository) Create(consignment *pb.Consigment)(*pb.Consignment, error) {
	repo.mu.Lock()
	updated:= append(repo.consignments, consignment)
	repo.consignments = updated
	repo.mu.Unlock()
	return consignment, nil
}

//Сервис должен реализовывать все методы
type service struct{
	repo repository
}

//CreateConsignment метод в сервисе , который является методом create
//обрабатываются сервером gRPC
func (s *service) CreateConsignment(ctx context.Context, 
	req *pb.Consigment)(*pb.Response, error) {
	//Сохраним партию
	consignment,err:= s.repo.Create(req)
	if err != nil {
		return nil, err
	}

	//Возврат ответа на сообщение `Response`, которое мы создали
	// в нашем protobuf

	return &pb.Response{Created: true, Consignment: consignment}, nil
}

func main() {
	repo:= &Repository{}
	//Настройка нашего сервера gRPC
	lis,err:=net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s:= grpc.NewServer()
	//Зарегистрируем наш сервер при помощи gRPC, это свяжет наш код
	pb.RegisterShippingServiceServer(s, &service{repo})
	reflection.Register(s)

	log.Println("Running on port:", port)
	if err:=s.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}