package server

import (
	context "context"

	"github.com/NothingXiang/go-rpc/api"
)

type Server struct {
}

func (s *Server) Echo(ctx context.Context, message *api.StringMessage) (*api.Person, error) {
	panic("implement me")
}
