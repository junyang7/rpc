package main

import (
	"client/pb"
	"context"
	"fmt"
	"github.com/junyang7/go-common/src/_json"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

func main() {

	header := map[string]string{
		"rule":   "/rpc/test",
		"method": "POST",
	}

	body := _json.Encode(map[string]string{
		"b":  "B",
		"b1": "B1",
	})

	conn, err := grpc.Dial("10.222.96.105:60001", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if nil != err {
		panic(err)
	}

	defer conn.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	c := pb.NewServiceClient(conn)

	r, err := c.Call(ctx, &pb.Request{
		Header: header,
		Body:   body,
	})

	if nil != err {
		panic(err)
	}

	b := r.GetResponse()
	res := string(b)

	fmt.Println("====>↓响应数据")
	fmt.Println(res)

}
