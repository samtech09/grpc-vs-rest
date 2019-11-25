package handlers

import (
	"fmt"

	"github.com/samtech09/grpc-vs-rest/grpc/model"
	"github.com/samtech09/grpc-vs-rest/grpc/service"
	"github.com/samtech09/grpc-vs-rest/server/reposit"

	"golang.org/x/net/context"
)

type Server struct {
}

// func NewServer() *Server {
// 	return &Server{}
// }

func (s *Server) GetDetail(ctx context.Context, in *model.Filter) (*service.ServiceResp, error) {
	ret := service.ServiceResp{}

	if in.From == "" {
		return &ret, fmt.Errorf("Invalid value for FROM")
	}
	if in.To == "" {
		return &ret, fmt.Errorf("Invalid value for TO")
	}
	dd, err := reposit.GetData(in.From, in.To)
	if err != nil {
		return &ret, err
	}

	ret.Count = int64(len(dd))
	ret.Data = dd
	return &ret, nil
}
