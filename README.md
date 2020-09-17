# NRPC

[![Build Status](https://travis-ci.org/teamlint/nrpc.svg?branch=master)](https://travis-ci.org/teamlint/nrpc)

NRPC 是一个基于 [NATS](https://nats.io/) 的 RPC 框架

通过 .proto 文件生成 NRPC 的客户端和服务端, 同时生 NATS 的消息处理器 [MsgHandler](https://godoc.org/github.com/nats-io/nats.go#MsgHandler).

## 为什么使用 NATS?

基于 NATS 的 [Request-Reply](https://docs.nats.io/nats-concepts/reqreply) 模型可以获得一些优于 gRPC 的特性: 

- **最小限度服务发现机制**: 客户端和服务端仅需要知道 NATS 集群连接结节点, 客户端不需要做其它服务发现
- **简单的负载均衡**: 每个微服务连接相同的 NATS 集群, 使用 NATS [Queue Groups](https://docs.nats.io/nats-concepts/queue) 可以随机处理传入请求, 不需要额外的负载均衡器

当然在大规模请求下 NATS 集群本身可能成为一个瓶颈,  类似 gRPC 的流处理及高级认证等功能没有提供支持, 但是 NRPC 简化了操作复杂性, 如果您的应用规模适中仍不失为一个优秀的选择.

NRPC 已成功的在 [OpsDash](https://www.opsdash.com) Saas 产品中应用并取得不错的进展.

## 概览

NRPC 附带了一个 protobuf 编译器插件 protoc-gen-nrpc，它可以
从 .proto 文件生成代码

给定一个  .proto 文件, 例如:  [helloworld.proto](https://github.com/grpc/grpc-go/blob/master/examples/helloworld/helloworld/helloworld.proto), 则可以这样使用:

```shell
$ ls
helloworld.proto
$ protoc --go_out=. --nrpc_out=. helloworld.proto
$ ls
helloworld.nrpc.go	helloworld.pb.go	helloworld.proto
```

生成的 `.pb.go` 文件由标准的 Go 插件生成, 包含了消息类的定义,  `.nrpc.go` 文件由 nrpc 插件生成, 包含了服务接口定义以及服务端及客户端调用代码, NATS 消息 Handler.

代码结构:

- 服务定义 [helloworld.proto](https://github.com/teamlint/nrpc/tree/master/examples/helloworld/helloworld/helloworld.proto)
- NRPC  [helloworld.nrpc.go](https://github.com/teamlint/nrpc/tree/master/examples/helloworld/helloworld/helloworld.nrpc.go)
- 服务端示例 [greeter_server/main.go](https://github.com/teamlint/nrpc/tree/master/examples/helloworld/greeter_server/main.go)
- 客户端示例 [greeter_client/main.go](https://github.com/teamlint/nrpc/tree/master/examples/helloworld/greeter_client/main.go)

## RPC 模型

### 调用模式

#### Request/Reply 请求回复模式

##### **标准模式**

服务端订阅消息, 客户端发送Request消息获取Reply消息

##### **流回复模式**

该模式下, 客户端与服务端建立连接后, 服务端持续向客户端发送消息, 客户端方法签名生成回复消息回调函数用于循环读取服务端消息,直至服务端消息发送结束

**示例:** 

**Proto 定义:**

```protobuf
option (nrpc.packageSubjectParams) = "instance"; // 包主题参数

service SvcCustomSubject {
	option (nrpc.serviceSubject) = 'custom_subject'; // 自定义服务主题
	rpc MtStreamedReply(StringArg) returns (SimpleStringReply) {
  	option (nrpc.streamedReply) = true;
	}
}
```

**服务端实现:**

```go
func (s BasicServerImpl) MtStreamedReply(ctx context.Context, req *StringArg, send func(rep *SimpleStringReply)) error {
  if req.GetArg1() == "please fail" {
    panic("Failing")
  }
  if req.GetArg1() == "very long call" {
    select {
    case <-ctx.Done():
      return ctx.Err()
    case <-time.After(time.Minute):
      time.Sleep(time.Minute)
   return nil
    }
  }
  time.Sleep(time.Second)
  send(&SimpleStringReply{Reply: "msg1"})
  time.Sleep(250 * time.Millisecond)
  send(&SimpleStringReply{Reply: "msg2"})
  time.Sleep(250 * time.Millisecond)
  send(&SimpleStringReply{Reply: "msg3"})
  time.Sleep(250 * time.Millisecond)
 return nil
}
```

**客户端调用:**

```go
client := NewSvcCustomSubjectClient(c, "default") // default 为包主题参数值
err := client.MtStreamedReply(context.Background(), &StringArg{Arg1: "arg"}, func(ctx context.Context, rep *SimpleStringReply) {
   fmt.Println("received", rep)
   resChan <- rep.GetReply()
})
```

 **结果:**
```shell
 received reply:"msg1"
 received reply:"msg2"
 received reply:"msg3"
```



#### NoReply 无回复模式

客户端发布消息, 服务端需要实现消息处理逻辑

**Proto 定义**:

```protobuf
option (nrpc.packageSubjectParams) = "instance"; // 包主题参数

service SvcCustomSubject {
	option (nrpc.serviceSubject) = 'custom_subject'; // 自定义服务主题
  rpc MtRequestNoReply(StringArg) returns (nrpc.NoReply) {}
}
```

**服务端实现**:

```go
func (s BasicServerImpl) MtRequestNoReply(ctx context.Context, req *StringArg) {
	s.t.Log("Will publish to MtRequestNoReply")
	s.t.Logf("client publish msg = %v\n", *req)
}
/*
func (s BasicServerImpl) MtNoReply(ctx context.Context) {
	s.t.Log("Will publish to MtNoRequest")
	s.handler.MtNoRequestPublish("default", &SimpleStringReply{Reply: "Hi there"})
	s.handler2.MtNoRequestWParamsPublish("default", "me", "mtvalue", &SimpleStringReply{Reply: "Hi there"})
}*/
```



**客户端调用**:

```go
c2 := NewSvcSubjectParamsClient(conn, "default", "me") // me 为服务主题参数值
sub, err := c2.MtNoRequestWParamsSubscribeSync( // 同步订阅
  "mtvalue",
)
if err != nil {
  t.Fatal(err)
}
defer sub.Unsubscribe()
c2.MtNoReply()
reply, err := sub.Next(time.Second)
if err != nil {
  t.Fatal(err)
}
if reply.GetReply() != "Hi there" {
  t.Errorf("Expected 'Hi there', got %s", reply.GetReply())
}

```





**结果:**

```shell

```





#### NoRequest 无请求模式

服务端发布消息, 客户端订阅消息

- 服务端生成

  服务端==不生成方法接口==, 转而生成 `方法名+Publish` 方法, 用于消息发布

- 客户端生成 

  - `方法名+Subject` 方法用于获取消息主题, 
  - `方法名+Subscribe`  用于异步订阅主题
  - `方法名+SubscribeSync`  用于同步订阅主题
  - `方法名+SubscribeChan`  用于管道订阅主题

**Proto 定义**:

```protobuf
option (nrpc.packageSubjectParams) = "instance"; // 包主题参数

service SvcCustomSubject {
	rpc MtNoRequest(nrpc.NoRequest) returns (SimpleStringReply) {}
}
```



**服务端实现**:



**客户端调用**:



**结果:**



### proto 定义

- nrpc.Void 无参数, 不改变调用模式, 服务端返回值带error

- nrpc.NoReply 改变调用模式
  客户端使用PUB模式, 服务端不返回任何值(不带error)

  **服务端签名:**
  
  ```go
   MtRequestNoReply(ctx context.Context, req *StringArg)
  ```
  
  
  
- nprc.NoRequest 改变调用模式, 不生成普通方法
  服务端使用PUB模式发布消息, 方法签名 `方法名+Publish`
  客户端订阅消息, 方法签名 `方法名+Subscribe` `方法名+Subject`

### 运行机制 

The .proto file defines messages (like HelloRequest and HelloReply in the
example) and services (Greeter) that have methods (SayHello).

The messages are generated as Go structs by the regular Go protobuf compiler
plugin and gets written out to \*.pb.go files.

For the rest, nRPC generates three logical pieces.

The first is a Go interface type (GreeterServer) which your actual
microservice code should implement:

```go
// This is what is contained in the .proto file
service Greeter {
    rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// This is the generated interface which you've to implement
type GreeterService interface {
    SayHello(ctx context.Context, req HelloRequest) (resp HelloReply, err error)
}
```

The second is a client (GreeterClient struct). This struct has
methods with appropriate types, that correspond to the service definition. The
client code will marshal and wrap the request object (HelloRequest) and do a
NATS `Request`.

```go
// The client is associated with a NATS connection.
func NewGreeterClient(nc *nats.Conn) *GreeterClient {...}

// And has properly typed methods that will marshal and perform a NATS request.
func (c *GreeterClient) SayHello(req HelloRequest) (resp HelloReply, err error) {...}
```

The third and final piece is the handler (GreeterHandler). Given a NATS
connection and a server implementation, it can accept NATS requests in the
format sent by the client above. It should be installed as a message handler for
a particular NATS subject (defaults to the name of the service) using the
NATS Subscribe() or QueueSubscribe() methods. It will invoke the appropriate
method of the GreeterServer interface upon receiving the appropriate request.

```go
// A handler is associated with a NATS connection and a server implementation.
func NewGreeterHandler(ctx context.Context, nc *nats.Conn, s GreeterService) *GreeterHandler {...}

// It has a method that can (should) be used as a NATS message handler.
func (h *GreeterHandler) MsgHandler(msg *nats.Msg) {...}
```

Standing up a microservice involves:

- writing the .proto service definition file
- generating the \*.pb.go and \*.nrpc.go files
- implementing the server interface
- writing a main app that will connect to NATS and start the handler ([see
  example](https://github.com/teamlint/nrpc/blob/master/examples/helloworld/greeter_server/main.go))

To call the service:

- import the package that contains the generated *.nrpc.go files
- in the client code, connect to NATS
- create a Caller object and call the methods as necessary ([see example](https://github.com/teamlint/nrpc/blob/master/examples/helloworld/greeter_client/main.go))

## 特性

查看链接获取详细信息:

- [负载均衡](https://github.com//nrpc/wiki/Load-Balancing)
- [度量统计](https://github.com/teamlint/nrpc/wiki/Metrics-Instrumentation)
  使用 [Prometheus](https://github.com/prometheus/prometheus)

## 安装

nRPC needs Go 1.7 or higher. $GOPATH/bin needs to be in $PATH for the protoc
invocation to work. To generate code, you need the protobuf compiler (which
you can install from [here](https://github.com/google/protobuf/releases))
and the nRPC protoc plugin.

安装  protoc NRPC 插件:

```shell
$ go get github.com/teamlint/nrpc/protoc-gen-nrpc
```

编译运行示例服务端 greeter_server:

```shell
$ go get github.com/teamlint/nrpc/examples/helloworld/greeter_server
$ greeter_server
server is running, ^C quits.
```

编译运行示例客户端  greeter_client:

```shell
$ go get github.com/teamlint/nrpc/examples/helloworld/greeter_client
$ greeter_client
Greeting: Hello world
$
```

## 参考文档

学习如何使用 .proto 文件描述 gRPC 服务, [看这里](https://grpc.io/docs/guides/concepts.html).
了解更多 NATS, 访问 [NATS 官网](https://nats.io/). 

## 当前状态
当前仅提供对 Go 语言的支持

原始版本由 RapidLoop 构建,  [Teamlint](http://www.teamlint.com) 进行扩展完善.

## TODO

- NoReply 请求方法签名返回值加error
- NRPC client request timeout or CallOption or Context
- user [protogen](google.golang.org/protobuf/compiler/protogen) refactor protoc-gen-nrpc



