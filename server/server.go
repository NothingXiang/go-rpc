package server

import (
	context "context"
	"fmt"

	"github.com/NothingXiang/go-rpc/api"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, message *api.StringMessage) (*api.Person, error) {
	//panic("implement me")

	fmt.Println("request: ", message)

	return message.P, nil
}
