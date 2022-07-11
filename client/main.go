package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

	fmt.Println("==================== reqs=10000 arrLen=10")
	start := time.Now()
	benchmark(client, 10000, 10)
	fmt.Println("======== RT total: ", time.Now().Sub(start))

	fmt.Println("==================== reqs=10000 arrLen=100")
	start = time.Now()
	benchmark(client, 10000, 100)
	fmt.Println("======== RT total: ", time.Now().Sub(start))

}

func benchmark(client echo.Client, reqs int, arrLen int) {

	for i := 0; i < reqs; i++ {
		arr := make([]string, arrLen)
		for j := 0; j < arrLen; j++ {
			arr[j] = randStr(5)
		}
		req := &api.Request{Messages: arr}
		client.Echo(context.Background(), req)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//log.Println(resp)
		//time.Sleep(time.Second)
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
