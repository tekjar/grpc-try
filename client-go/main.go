package main

import (
	"client/helloworld"
	"context"
	"fmt"
	"log"
	"time"

	load "github.com/shirou/gopsutil/load"
	mem "github.com/shirou/gopsutil/mem"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func profileSystem(exit chan bool) {
	ticker := time.NewTicker(time.Second * 30).C

	cpu, err := load.Avg()
	if err != nil {
		log.Fatalln("Unable to get load avg. Error = ", err)
	}

	sysMem, err := mem.VirtualMemory()
	if err != nil {
		log.Fatalln("Unable to get memory. Error = ", err)
	}

	fmt.Println("1 = ", cpu.Load1, " 5 = ", cpu.Load5, " 15 = ", cpu.Load15)
	fmt.Println("Pre memory usage = ", sysMem.Used)
	memory := []uint64{}

	for {
		select {
		case <-ticker:
			sysMem, err := mem.VirtualMemory()
			if err != nil {
				log.Fatalln("Unable to get memory. Error = ", err)
			}

			memory = append(memory, sysMem.Used)
		case <-exit:
			cpu, err := load.Avg()
			if err != nil {
				log.Fatalln("Unable to get load avg. Error = ", err)
			}

			var totalMem uint64
			for _, value := range memory {
				totalMem += value
			}

			fmt.Println("1 = ", cpu.Load1, " 5 = ", cpu.Load5, " 15 = ", cpu.Load15)
			fmt.Println("Post average memory usage = ", totalMem/uint64(len(memory)))
			return
		}
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := helloworld.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	exit := make(chan bool, 1)

	go profileSystem(exit)
	for i := 0; i < 3000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})

		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		log.Printf("Greeting: %s", r.A)
		time.Sleep(100 * time.Millisecond)
	}

	exit <- true
	time.Sleep(3 * time.Second)
}
