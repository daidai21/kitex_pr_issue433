package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/daidai21/kitex_pr_issue433/kitex_gen/api"
	"github.com/daidai21/kitex_pr_issue433/kitex_gen/api/echo"
)

func main() {
	client, err := echo.NewClient("echo", client.WithHostPorts("[::1]:8888"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Messages: []string{"my request", "a123"}}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp)
		time.Sleep(time.Second)
	}
}
