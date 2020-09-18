package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nats-io/nats.go"

	// This is the package containing the generated *.pb.go and *.nrpc.go
	// files.
	"github.com/teamlint/nrpc/examples/helloworld/helloworld"
)

// greeterService implements the helloworld.GreeterService interface.
type greeterService struct{}

// SayHello is an implementation of the SayHello method from the definition of
// the Greeter service.
func (s *greeterService) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	resp := new(helloworld.HelloReply)
	resp.Message = "Hello " + req.Name
	return resp, nil
}

func main() {
	var natsURL = nats.DefaultURL
	if len(os.Args) == 2 {
		natsURL = os.Args[1]
	}
	// Connect to the NATS server.
	nc, err := nats.Connect(natsURL, nats.Timeout(5*time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Our server implementation.
	s := &greeterService{}

	// The NATS handler from the helloworld.nrpc.proto file.
	h := helloworld.NewGreeterHandler(context.TODO(), nc, s)

	// Start a NATS subscription using the handler. You can also use the
	// QueueSubscribe() method for a load-balanced set of servers.
	sub, err := nc.Subscribe(h.Subject(), h.MsgHandler)
	if err != nil {
		log.Fatal(err)
	}
	defer sub.Unsubscribe()

	// Keep running until ^C.
	fmt.Println("server is running, ^C quits.")
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	close(c)
}
