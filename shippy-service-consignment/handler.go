package main

import (
	"context"

	pb "github.com/EvgenyiK/shippy/shippy-service-consignment/proto/consignment"
	vesselProto "github.com/EvgenyiK/shippy/shippy-service-vessel/proto/vessel"
	"github.com/pkg/errors"
)

type handler struct{
	repository
	vesselClient vesselProto.VesselService
}
//CreateConsignment - мы создали только один метод в нашем сервисе

func (s *handler) CreateConsignment(ctx context.Context, req *pb.Consignment, res *pb.Response) error{
	//Создаем экземпляр сервиса судна и устанавливаем вместимость еонтейнеров.
	vesselResponse, err:= s.vesselClient.FindAvailable(ctx, &vesselProto.Specification{
		MaxWeight: req.Weight,
		Capacity: int32(len(req.Containers)),
	})
	if vesselResponse == nil {
		return errors.New("error fetching vessel, returned nil")
	}
	if err != nil {
		return  err
	}
	//Устанавливаем ID у судна после обслуживания сервисом
	req.VesselId = vesselResponse.Vessel.Id
	//Сохраняем партию грузов
	if err = s.repository.Create(ctx, MarshalConsignment(req)); err!=nil{
		return err
	}
	res.Created = true
	res.Consignment = req
	return nil
}

// GetConsignments
func (s *handler) GetConsigments(ctx context.Context, req *pb.GetRequest, res *pb.Response)error {
	consignments,err:= s.repository.GetAll(ctx)
	if err != nil {
		return  err
	}
	res.Consignments = UnmarshalConsignmentCollection(consignments)
	return nil
}