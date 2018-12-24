package main

import (
	"client/helloworld"
	"context"
	"fmt"
	"log"
	"time"

	profile "github.com/akhenakh/statgo"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func profileSystem(exit chan bool) {
	ticker := time.NewTicker(time.Second * 10).C
	stats := profile.NewStat()

	cpu := stats.CPUStats()
	fmt.Println("1 = ", cpu.LoadMin1, " 5 = ", cpu.LoadMin5, " 15 = ", cpu.LoadMin15)
	memory := []int{}

	for {
		select {
		case <-ticker:
			mem := stats.MemStats()
			cpu := stats.CPUStats()
			fmt.Println("1 = ", cpu.LoadMin1, " 5 = ", cpu.LoadMin5, " 15 = ", cpu.LoadMin15)
			memory = append(memory, mem.Used)
		case <-exit:
			cpu := stats.CPUStats()
			var totalMem int
			for _, value := range memory {
				totalMem += value
			}
			fmt.Println("1 = ", cpu.LoadMin1, " 5 = ", cpu.LoadMin5, " 15 = ", cpu.LoadMin15)
			fmt.Println("average system memory = ", totalMem/len(memory))
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
	for i := 0; i < 10000; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		r, err := c.SayHello(ctx, &helloworld.HelloRequest{Name: name})
		_ = r
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}

		// log.Printf("Greeting: %s", r.A)
		time.Sleep(100 * time.Millisecond)
	}

	exit <- true
	time.Sleep(3 * time.Second)
}
