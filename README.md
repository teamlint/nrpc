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

服务端订阅消息, 客户端发送 Request 消息获取 Reply 消息

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

客户端发布消息, 服务端需要实现消息处理逻辑, 无任何返回值

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
  // TODO: 处理客户端发送过来的消息, 对于长时处理任务, 可以创建个NoRequest API, 然后在这里给客户端发送消息,客户端订阅消息进行后续处理 如:
  s.handler.MtNoRequestPublish("default", &SimpleStringReply{Reply: "Hi there"})
}
```

**客户端调用**:

```go
// 客户端订阅消息
c1 := NewSvcCustomSubjectClient(conn, "default")
arg := "[client.sync]req-noreply -> [server]process -> [server]publish to client"
// 创建同步消息订阅器
sub, err := c1.MtNoRequestSubscribeSync() // NoRequest 模式生成的订阅方法
if err != nil {
  t.Fatal(err)
}
defer sub.Unsubscribe()
err = c1.MtRequestNoReply(&StringArg{Arg1: arg}) // NoReply 客户端方法
if err != nil {
  t.Fatal(err)
}
// 获取消息
reply, err := sub.Next(10 * time.Second) 
if err != nil {
  t.Fatal(err)
}
```

#### NoRequest 无请求模式

服务端发布消息, 客户端订阅消息

- 服务端生成

  服务端==不生成方法接口==, 转而生成 `RPC方法名+Publish` 方法, 用于消息发布

- 客户端生成 

  - `RPC方法名+Subject` 方法用于获取消息主题, 
  - `RPC方法名+Subscribe`  用于异步订阅主题
  - `RPC方法名+SubscribeSync`  用于同步订阅主题
  - `RPC方法名+SubscribeChan`  用于管道订阅主题

**Proto 定义**:

```protobuf
option (nrpc.packageSubjectParams) = "instance"; // 包主题参数

service SvcCustomSubject {
	rpc MtNoRequest(nrpc.NoRequest) returns (SimpleStringReply) {}
}
```

**服务端实现**:

服务端==不生成方法接口==, 不需要具体实现, 转而生成 `RPC方法名+Publish` 方法, 服务端其它方法实现进行相关调用

**客户端调用**:

客户端直接订阅消息主题进行处理, 支持同步订阅, 异步订阅, 管道订阅

```go
// 客户端订阅消息
c1 := NewSvcCustomSubjectClient(conn, "default")
repChan := make(chan string)
// 创建异步消息订阅器
sub, err := c1.MtNoRequestSubscribe(func(reply *SimpleStringReply) {
  defer close(repChan)
  repChan <- reply.GetReply()
})
if err != nil {
  t.Fatal(err)
}
defer sub.Unsubscribe()
// 获取消息
for rep := range repChan {
  t.Log(rep)
}
```

### Protobuf 参数

- nrpc.Void 空参数

  不改变调用模式, 服务端返回值带 error

- nrpc.NoReply 改变调用模式
  客户端使用PUB模式, 服务端不返回任何值(不带error)

- nprc.NoRequest 改变调用模式, 不生成方法接口
  服务端使用PUB模式发布消息, 方法签名 `RPC方法名+Publish`
  客户端订阅消息,
  
  - 订阅方法签名 `RPC方法名+Subscribe[Sync|Chan]` 
  - 获取消息主题方法签名`RPC方法名+Subject`

## 特性

查看链接获取详细信息:

- [负载均衡](https://github.com//nrpc/wiki/Load-Balancing)
- [度量统计](https://github.com/teamlint/nrpc/wiki/Metrics-Instrumentation)
  使用 [Prometheus](https://github.com/prometheus/prometheus)

## 安装

### 安装 protoc 工具

[下载](https://github.com/protocolbuffers/protobuf) 并安装 protocol buffer 编译工具

### 安装  protoc GO 插件

```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go
```

### 安装  protoc NRPC 插件

```shell
$ go get github.com/teamlint/nrpc/protoc-gen-nrpc
```

## 示例

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
```

## 参考文档

学习如何使用 .proto 文件描述 gRPC 服务, [看这里](https://grpc.io/docs/guides/concepts.html).
了解更多 NATS, 访问 [NATS 官网](https://nats.io/). 

## 当前状态
当前仅提供对 Go 语言的支持

原始版本由 RapidLoop 构建,  [Teamlint](http://www.teamlint.com) 进行扩展完善.

## TODO

- NRPC client request timeout or CallOption or Context
- user [protogen](google.golang.org/protobuf/compiler/protogen) refactor protoc-gen-nrpc
- Hub 动态调用不同客户端, pkg实例参数?



