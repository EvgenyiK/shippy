package main

import (
	"context"
	"errors"
	"log"

	pb "github.com/EvgenyiK/shippy/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

type Repositiry interface{
	FindAvailable(*pb.Specification)(*pb.Vessel, error)
}

type VesselRepository struct{
	vessels []*pb.Vessel
}

//FindAvailable - проверяет спецификацию по карте судов,
//Если вместимость и вес ниже вместимости судна ,то
//возвращаем это судно.

func (repo *VesselRepository) FindAvailable(spec *pb.Specification)(*pb.Vessel, error){
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight{
			return vessel, nil
		}
	}
	return nil, errors.New("No vessel found by that spec")
}

//Сервисный обработчик grpc
type vesselService struct{
	repo Repositiry
}

func(s *vesselService)FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response)error{
	//Найти доступное судно
	vessel, err :=s.repo.FindAvailable(req)
	if err != nil {
		return err
	}
	//получить ответ
	res.Vessel = vessel
	return nil
}

func main() {
	vessels:=[]*pb.Vessel{
		&pb.Vessel{Id:"vessel001", Name: "Boaty McBoatface", MaxWeight: 200000, Capacity: 500},}
		repo := &VesselRepository{vessels}
		service:=micro.NewService(
			micro.Name("shippy.service.vessel"),
		)
		service.Init()

		//Регистрация нашей реализации
		if err:=pb.RegisterVesselServiceHandler(service.Server(),&vesselService{repo}); err != nil {
			log.Panic(err)
		}
		if err:=service.Run(); err != nil {
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