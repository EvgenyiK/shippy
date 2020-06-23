package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"shippy/shippy-service-consignment/proto/consignment"

	"context"

	pb "github.com/EvgenyiK/shippy/shippy-service-consignment/proto/consignment"
	"google.golang.org/grpc"
)

const(
	address	= "localhost:50051"
	defaultFilename = "consigment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data,err:=ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err	
}

func main() {
	//Коннектимся к серверу
	conn,err:= grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	client:=pb.NewShippingSrviceClient(conn)
	//Ответ от сервера
	file:=defaultFilename
	if len(os.Args)>1 {
		file = os.Args[1]
	}

	consignment, err:=parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}
	r, err:= client.CreateConsignment(context.Background(), consignment)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)
}