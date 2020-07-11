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