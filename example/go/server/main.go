package main

import (
	"context"
	"fmt"
	"github.com/junyang7/go-common/src/_interceptor"
	"github.com/junyang7/go-common/src/_json"
	"github.com/junyang7/go-common/src/_response"
	"google.golang.org/grpc"
	"net"
	"server/pb"
)

type rpcCall struct {
	pb.UnimplementedServiceServer
}

func (this *rpcCall) Call(c context.Context, r *pb.Request) (oRes *pb.Response, oErr error) {

	// 接受请求数据
	fmt.Println("↑请求数据")
	fmt.Println(r.Header)
	fmt.Println(r.Body)
	var a map[string]interface{}
	_json.Decode(r.Body, &a)
	fmt.Println(a)

	res := _response.New()
	defer func() {
		if err := recover(); nil != err {
			res.Code = -1
			res.Message = fmt.Sprintf("%v", err)
			oRes = &pb.Response{Response: _json.Encode(res)}
		}
	}()

	//
	// 处理业务逻辑
	//

	// 业务逻辑返回数据
	res.Code = 0
	res.Message = "success"
	res.Data = map[string]string{"test": "Hello World!"}
	oRes = &pb.Response{Response: _json.Encode(res)}

	return oRes, oErr

}

func main() {

	l, err := net.Listen("tcp", ":60001")
	if nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}
	s := grpc.NewServer()
	pb.RegisterServiceServer(s, &rpcCall{})
	if err := s.Serve(l); nil != err {
		_interceptor.Insure(false).Message(err).Do()
	}

}
