package main

import (
	"context"
	"log"

	"github.com/daidai21/kitex_pr_issue433/kitex_gen/api"
	"github.com/daidai21/kitex_pr_issue433/kitex_gen/api/echo"
)

var _ api.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	//klog.Info("echo called")
	return &api.Response{Message: "123"}, nil
}

func main() {
	svr := echo.NewServer(new(EchoImpl))
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
