package main

import (
	"log"
	"strconv"
	"testing"

	natsd "github.com/nats-io/nats-server/v2/server"
)

var natsURL string

func TestMain(m *testing.M) {
	// gnatsd := natsServer.New(&natsServer.Options{Port: natsServer.RANDOM_PORT})
	// gnatsd.SetLogger(
	// 	logger.NewStdLogger(false, false, false, false, false),
	// 	false, false)
	// go gnatsd.Start()
	// defer gnatsd.Shutdown()

	// if !gnatsd.ReadyForConnections(time.Second) {
	// 	log.Fatal("Cannot start the gnatsd server")
	// }
	// natsURL = "nats://" + gnatsd.Addr().String()

	// os.Exit(m.Run())

	opts := natsd.Options{
		Port: 4222,
	}
	s, err := natsd.NewServer(&opts)
	if err != nil {
		log.Printf("NATS server run err=%v\n", err)
		panic(err)
	}
	// log.Printf("NATS server =%v\n", s)
	natsURL = "nats://" + strconv.Itoa(opts.Port)
	go natsd.Run(s) // or
	// s.Start()
	// defer s.Shutdown()
	log.Printf("NATS server(%v) running... \n", "default")
}
