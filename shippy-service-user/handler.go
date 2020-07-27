package main

import(
	"context"
	"errors"

	pb "github.com/EvgenyiK/shippy/shippy-service-user/proto/user"
	"golang.org/x/crypto/bcrypt"
)

type authable interface{
	Decode(token string)(*CustomClaims, error)
	Encode(user *pb.User)(string,error)
}

type handler struct{
	repository Repository
	tokenService authable
}

func (s *handler)Get(ctx context.Context, req *pb.User, res *pb.Response) error{
	result,err:= s.repository.Get(ctx,req.Id)
	if err != nil {
		return  err
	}

	user:= UnmarshalUser(result)
	res.User = user

	return nil
}

func (s *handler)GetAll(ctx context.Context, req *pb.User, res *pb.Response) error{
	results,err:= s.repository.GetAll(ctx)
	if err != nil {
		return  err
	}

	users:= UnmarshalUserCollection(results)
	res.User = user

	return nil
}