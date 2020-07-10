package main

import (
	"context"

	pb "github.com/EvgenyiK/shippy/shippy-service-consignment/proto/consignment"
	"go.mongodb.org/mongo-driver/mongo"
)

type Consignment struct {
	ID          string     `json:"id"`
	Weight      int32      `json:"weight"`
	Description string     `json:"description"`
	Containers  Containers `json:"containers"`
	VesselID    string     `json:"vessel_id"`
}

type Container struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	UserID     string `json:"user_id"`
}

type Containers []*Container

func MarshalContainerCollection(containers *pb.Container) []*Container {
	collection := make([]*Container, 0)
	for _, container := range containers {
		collection = append(collection, MarshalContainer(container))
	}
	return collection
}

func UnmarshalContainerCollection(containers *pb.Container) []*Container {
	collection := make([]*pb.Container, 0)
	for _, container := range containers {
		collection = append(collection, UnmarshalContainer(container))
	}
	return collection
}

func UnmarshalConsignmentCollection(consigments []*Consignment) []*pb.Consignment {
	collection := make([]*pb.Consignment, 0)
	for _, container := range consigments {
		collection = append(collection, UnmarshalConsiggnment(container))
	}
	return collection
}

func UnmarshalContainer(container *Container) *pb.Container {
	return &pb.Container{
		Id:         container.ID,
		CustomerId: container.CustomerID,
		UserId:     container.UserID,
	}
}

func MarshalContainer(container *pb.Container) *Container {
	return &Container{
		ID:         container.Id,
		CustomerID: container.CustomerId,
		UserID:     container.UserId,
	}
}

func MarshalConsignment(consignment *pb.Consignment) *Consignment {
	containers := MarshalContainerCollection(consignment.Containers)
	return &Consignment{
		ID:          consignment.Id,
		Weight:      consignment.Weight,
		Description: consignment.Description,
		Containers:  containers,
		VesselID:    consignment.VesselID,
	}
}

type repository interface {
	Create(ctx context.Context, consignment *Consignment) error
	GetAll(ctx context.Context) ([]*Consignment, error)
}

//Реализация Mongo
type MongoRepository struct {
	collection *mongo.Collection
}

//Создание
func (repository *MongoRepository) Create(ctx context.Context, consignment *Consignment) error {
	_, err := repository.collection.InsertOne(ctx, consignment)
	return err
}

//Выборка всех
func (repository *MongoRepository) GetAll(ctx context.Context) ([]*Consignment, error){
	cur,err:= repository.collection.Find(ctx, nil, nil)
	var consignments []*Consignment
	for cur.Next(ctx){
		var consignment *Consignment
		if err:= cur.Decode(&consignment); err!=nil{
			return nil, err
		}
		consignments = append(consignments, consignment)
	}
	return consignments, err
}
