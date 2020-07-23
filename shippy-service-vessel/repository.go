package main

import(
	"context"

	pb "github.com/EvgenyiK/shippy/shippy-service-vessel/proto/vessel"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository interface{
	FindAvailable(ctx context.Context, spec *Specification)(*Vessel, error)
	Create(ctx context.Context, vessel *Vessel) error
}

type MongoRepository struct{
	collection *mongo.Collection
}

type Specification struct{
	Capacity int32
	MaxWeight int32
}

func MarshalSpecification(spec *pb.Specification) *Specification{
	return &Specification{
		Capacity: spec.Capasity,
		MaxWeight: spec.MaxWeight,
	}
}

func UnmarshalVessel(vessel *Vessel) *pb.Vessel{
	return &Vessel{
		ID:	vessel.ID,
		Capacity: vessel.Capacity,
		MaxWeight:	vessel.MaxWeight,
		Name:	vessel.Name,
		Available:	vessel.Available,
		OwnerID:	vessel.OwnerID,
	}
}

type Vessel struct{
	ID	string
	Capacity int
	Name string
	Available bool
	OwnerID string
	MaxWeight int32
}

//Если вместимость и максимальный вес ниже вместимости судна,то возвращаем его
func (repository *MongoRepository)FindAvailable(ctx context.Context, spec *Specification) (*Vessel, error) {
	filter:=bson.D{{
		"capacity",
		bson.D{{
			"$lte",
			spec.Capacity,
		},{
			"$lte",
			spec.MaxWeight,
		}},
	}}
	vessel:=&Vessel{}
	if err := repository.collection.FindOne(ctx,filter).Decode(vessel); err != nil {
		return nil, err
	}
	return vessel, nil
}

func (repository *MongoRepository) Create(ctx context.Context, vessel *Vessel) error {
	_, err:= repository.collection.InsertOne(ctx,vessel)
	return err
}