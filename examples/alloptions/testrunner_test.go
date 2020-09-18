package main

import (
	"log"
	"os"
	"strconv"
	"testing"

	natsd "github.com/nats-io/nats-server/v2/server"
)

var NatsURL string

func TestMain(m *testing.M) {
	opts := natsd.Options{
		Port: 4222,
	}
	s, err := natsd.NewServer(&opts)
	if err != nil {
		log.Printf("NATS server run err=%v\n", err)
		panic(err)
	}
	NatsURL = "nats://localhost:" + strconv.Itoa(opts.Port)
	go natsd.Run(s)
	defer s.Shutdown()
	log.Printf("NATS server(%v) running... \n", "default")
	os.Exit(m.Run())
}
