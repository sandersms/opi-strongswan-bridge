// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.

// Package main is the main package of the application
package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/opiproject/opi-api/security/v1alpha1/gen/go"
	"github.com/opiproject/opi-strongswan-bridge/pkg/ipsec"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var port int
	flag.IntVar(&port, "server_port", 50051, "The server port")

	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ipsecServ := ipsec.NewServer()

	pb.RegisterIkePeerServiceServer(s, ipsecServ)
	pb.RegisterIkeConnectionServiceServer(s, ipsecServ)
	pb.RegisterIpsecSaServiceServer(s, ipsecServ)
	pb.RegisterIpsecPolicyServiceServer(s, ipsecServ)

	reflection.Register(s)

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
