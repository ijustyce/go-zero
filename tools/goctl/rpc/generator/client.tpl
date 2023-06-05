package rpc

import (
	"time"

	"github.com/zeromicro/go-zero/zrpc"
)

func TestCall() {
	config := zrpc.RpcClientConf{NonBlock: true, Timeout: 10_000, KeepaliveTime: 10 * time.Second, Target: "127.0.0.1:8080"}
	zrpcClient, err := zrpc.NewClient(config)
	////	下面只是一个示例，请按实际情况调整代码
	//sayHelloClient := sayhello.NewSayHello(zrpcClient)
	//resp, err := sayHelloClient.Greet(context.Background(), &sayhello.HelloReq{Name: "chun"})
	//if err != nil {
	//	log.Fatalln(err)
	//} else {
	//	log.Default().Println(resp.Message)
	//}
}