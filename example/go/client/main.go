package main

import (
	"client/pb"
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {

	// 构建请求头
	header := map[string]string{
		"rule":   "/rpc/test",
		"method": "POST",
	}

	// 构建请求体
	body, err := json.Marshal(map[string]string{
		"b":  "B",
		"b1": "B1",
	})
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	// 连接到 gRPC 服务端
	conn, err := grpc.Dial("127.0.0.1:60001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	// 设置超时上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// 创建客户端
	c := pb.NewServiceClient(conn)

	// 发起 RPC 调用
	r, err := c.Call(ctx, &pb.Request{
		Header: header,
		Body:   body,
	})
	if err != nil {
		log.Fatalf("RPC call failed: %v", err)
	}

	// 获取响应并转换为字符串
	b := r.GetResponse()
	res := string(b)

	// 打印响应数据
	fmt.Println("====>↓响应数据")
	fmt.Println(res)
}
