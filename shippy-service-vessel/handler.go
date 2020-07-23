package main

import(
	"context"
	pb "github.com/EvgenyiK/shippy/shippy-service-vessel/proto/vessel"
)

type handler struct{
	repository
}

//найти все доступные судна
func (s *handler) FindAvailable(ctx context.Context, req *pb.Specifications, res *pb.Response)error{
	//найти доступное судно
	vessel, err:= s.repository.FindAvailable(ctx, MarshalSpecification(req))
	if err != nil {
		return err
	}
	//записать ответ
	res.Vessel = UnmarshalVessel(vessel)
	return nil
}

func (s *handler) Create(ctx context.Context, req *pb.Vessel, res *pb.Response) error {
	if err:= s.repository.Create(ctx, MarshalVessel(req));err != nil {
		return  err
	}
	res.Vessel = req
	return nil
}