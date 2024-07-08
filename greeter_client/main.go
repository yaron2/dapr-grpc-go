/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "google.golang.org/grpc/examples/helloworld/helloworld"
	"google.golang.org/grpc/metadata"
)

var (
	addr = flag.String("addr", "localhost:50001", "the address to connect to")
)

func main() {
	flag.Parse()

	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx := context.TODO()
	ctx = metadata.AppendToOutgoingContext(ctx, "dapr-app-id", "server")

	i := 0

	for {
		i++

		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: fmt.Sprintf("%v", i)})
		if err != nil {
			log.Printf("could not greet: %v", err)
			return
		}

		t := time.Now()
		log.Printf("Greeting: %s, %s", r.GetMessage(), t.Format("20060102150405"))
		time.Sleep(time.Second * 1)
	}
}
