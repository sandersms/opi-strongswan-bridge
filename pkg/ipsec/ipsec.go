// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2022-2023 Intel Corporation, or its subsidiaries.

// Package ipsec is the main package of the application
package ipsec

import (
	"context"
	"log"

	pb "github.com/opiproject/opi-api/security/v1alpha1/gen/go"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Server represents the Server object
type Server struct {
	pb.UnimplementedIkePeerServiceServer
	pb.UnimplementedIkeConnectionServiceServer
	pb.UnimplementedIpsecSaServiceServer
	pb.UnimplementedIpsecPolicyServiceServer
}

// TODO: Refactor IpSec implementation with new CRUD model Protobufs

// Ike Peer Service
// Create Ike Peer
func (s *Server) CreateIkePeer(_ context.Context, in *pb.CreateIkePeerRequest) (*pb.IkePeer, error) {
	log.Printf("CreateIkePeer: received from %v", in)
	return &pb.IkePeer{}, nil
}

// Delete Ike Peer
func (s *Server) DeleteIkePeer(_ context.Context, in *pb.DeleteIkePeerRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteIkePeer: received from %v", in)
	return &emptypb.Empty{}, nil
}

// Update Ike Peer
func (s *Server) UpdateIkePeer(_ context.Context, in *pb.UpdateIkePeerRequest) (*pb.IkePeer, error) {
	log.Printf("UpdateIkePeer: received from %v", in)
	return &pb.IkePeer{}, nil
}

// Get Ike Peer
func (s *Server) GetIkePeer(_ context.Context, in *pb.GetIkePeerRequest) (*pb.IkePeer, error) {
	log.Printf("GetIkePeer: received from %v", in)
	return &pb.IkePeer{}, nil
}

// List Ike Peer
func (s *Server) ListIkePeers(_ context.Context, in *pb.ListIkePeersRequest) (*pb.ListIkePeersResponse, error) {
	log.Printf("ListIkePeers: received from %v", in)
	return &pb.ListIkePeersResponse{}, nil
}

// Ike Connection Service
// Create Ike Connection
func (s *Server) CreateIkeConnection(_ context.Context, in *pb.CreateIkeConnectionRequest) (*pb.IkeConnection, error) {
	log.Printf("CreateIkeConnection: received from %v", in)
	return &pb.IkeConnection{}, nil
}

// Delete Ike Connection
func (s *Server) DeleteIkeConnection(_ context.Context, in *pb.DeleteIkeConnectionRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteIkeConnection: received from %v", in)
	return &emptypb.Empty{}, nil
}

// Update Ike Connection
func (s *Server) UpdateIkeConnection(_ context.Context, in *pb.UpdateIkeConnectionRequest) (*pb.IkeConnection, error) {
	log.Printf("UpdateIkeConnection: received from %v", in)
	return &pb.IkeConnection{}, nil
}

// Get Ike Connection
func (s *Server) GetIkeConnection(_ context.Context, in *pb.GetIkeConnectionRequest) (*pb.IkeConnection, error) {
	log.Printf("GetIkeConnection: received from %v", in)
	return &pb.IkeConnection{}, nil
}

// List Ike Connections
func (s *Server) ListIkeConnections(_ context.Context, in *pb.ListIkeConnectionsRequest) (*pb.ListIkeConnectionsResponse, error) {
	log.Printf("ListIkeConnections: received from %v", in)
	return &pb.ListIkeConnectionsResponse{}, nil
}

// Get Ike Connection Statistics
func (s *Server) StatsIkeConnections(_ context.Context, in *pb.StatsIkeConnectionsRequest) (*pb.StatsIkeConnectionsResponse, error) {
	log.Printf("StatsIkeConnections: received from %v", in)

	stats, err := ipsecStats()
	if err != nil {
		log.Printf("IPsecStats: Failed %v", err)
		return nil, err
	}

	return stats, nil

}

// Ipsec Sa Service
// Create Ipsec Security Association
func (s *Server) CreateIpsecSa(_ context.Context, in *pb.CreateIpsecSaRequest) (*pb.IpsecSa, error) {
	log.Printf("CreateIpsecSa: received from %v", in)
	return &pb.IpsecSa{}, nil
}

// Delete Ipsec Security Association
func (s *Server) DeleteIpsecSa(_ context.Context, in *pb.DeleteIpsecSaRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteIpsecSa: received from %v", in)
	return &emptypb.Empty{}, nil
}

// Update Ipsec Security Association
func (s *Server) UpdateIpsecSa(_ context.Context, in *pb.UpdateIpsecSaRequest) (*pb.IpsecSa, error) {
	log.Printf("UpdateIpsecSa: received from %v", in)
	return &pb.IpsecSa{}, nil
}

// Get Ipsec Security Association
func (s *Server) GetIpsecSa(_ context.Context, in *pb.GetIpsecSaRequest) (*pb.IpsecSa, error) {
	log.Printf("GetIpsecSa: received from %v", in)
	return &pb.IpsecSa{}, nil
}

// List Ipsec Security Associations
func (s *Server) ListIpsecSas(_ context.Context, in *pb.ListIpsecSasRequest) (*pb.ListIpsecSasResponse, error) {
	log.Printf("ListIpsecSas: received from %v", in)
	return &pb.ListIpsecSasResponse{}, nil
}

// Ipsec Policy Service
// Create IPsec Policy
func (s *Server) CreateIpsecPolicy(_ context.Context, in *pb.CreateIpsecPolicyRequest) (*pb.IpsecPolicy, error) {
	log.Printf("CreateIpsecPolicy: received from %v", in)
	return &pb.IpsecPolicy{}, nil
}

// Delete IPsec Policy
func (s *Server) DeleteIpsecPolicy(_ context.Context, in *pb.DeleteIpsecPolicyRequest) (*emptypb.Empty, error) {
	log.Printf("DeleteIpsecPolicy: received from %v", in)
	return &emptypb.Empty{}, nil
}

// Update IPsec Policy
func (s *Server) UpdateIpsecPolicy(_ context.Context, in *pb.UpdateIpsecPolicyRequest) (*pb.IpsecPolicy, error) {
	log.Printf("UpdateIpsecPolicy: received from %v", in)
	return &pb.IpsecPolicy{}, nil
}

// Get IPsec Policy
func (s *Server) GetIpsecPolicy(_ context.Context, in *pb.GetIpsecPolicyRequest) (*pb.IpsecPolicy, error) {
	log.Printf("GetIpsecPolicy: received from %v", in)
	return &pb.IpsecPolicy{}, nil
}

// List IPsec Policies
func (s *Server) ListIpsecPolicies(_ context.Context, in *pb.ListIpsecPoliciesRequest) (*pb.ListIpsecPoliciesResponse, error) {
	log.Printf("ListIIpsecPolicies: received from %v", in)
	return &pb.ListIpsecPoliciesResponse{}, nil
}

/*
// IPsecVersion executes the ipsecVersion
func (s *Server) IPsecVersion(_ context.Context, _ *pb.IPsecVersionRequest) (*pb.IPsecVersionResponse, error) {
	ver, err := ipsecVersion()
	if err != nil {
		log.Printf("IPsecVersion: Failed %v", err)
		return nil, err
	}

	return ver, nil
}

// IPsecStats executes the ipsecStats
func (s *Server) IPsecStats(_ context.Context, _ *pb.IPsecStatsRequest) (*pb.IPsecStatsResponse, error) {
	stats, err := ipsecStats()
	if err != nil {
		log.Printf("IPsecStats: Failed %v", err)
		return nil, err
	}

	return stats, nil
}

// IPsecInitiate executes the initiateConn
func (s *Server) IPsecInitiate(_ context.Context, in *pb.IPsecInitiateRequest) (*pb.IPsecInitiateResponse, error) {
	log.Printf("IPsecInitiate: Received: %v", in)

	err := initiateConn(in)
	if err != nil {
		log.Printf("IPsecInitiate: Failed %v", err)
		return nil, err
	}

	ret := pb.IPsecInitiateResponse{}

	return &ret, nil
}

// IPsecTerminate executes the terminateConn
func (s *Server) IPsecTerminate(_ context.Context, in *pb.IPsecTerminateRequest) (*pb.IPsecTerminateResponse, error) {
	log.Printf("IPsecTerminate: Received: %v", in)

	matches, err := terminateConn(in)
	if err != nil {
		log.Printf("IPsecTerminate: Failed %v", err)
		return nil, err
	}

	ret := pb.IPsecTerminateResponse{
		Success: "Yes",
		Matches: matches,
	}

	return &ret, nil
}

// IPsecRekey executes the rekeyConn
func (s *Server) IPsecRekey(_ context.Context, in *pb.IPsecRekeyRequest) (*pb.IPsecRekeyResponse, error) {
	log.Printf("IPsecRekey: Received: %v", in)

	success, matches, err := rekeyConn(in)
	if err != nil {
		log.Printf("IPsecRekey: Failed: %v", err)
		return nil, err
	}

	ret := pb.IPsecRekeyResponse{
		Success: success,
		Matches: matches,
	}

	return &ret, nil
}

// IPsecListSas executes the listSas
func (s *Server) IPsecListSas(_ context.Context, in *pb.IPsecListSasRequest) (*pb.IPsecListSasResponse, error) {
	log.Printf("IPsecListSas: Received %v", in)

	ret, err := listSas(in)
	if err != nil {
		log.Printf("IPsecListSas: Failed: %v", err)
		return nil, err
	}

	return ret, nil
}

// IPsecListConns executes the listConns
func (s *Server) IPsecListConns(_ context.Context, in *pb.IPsecListConnsRequest) (*pb.IPsecListConnsResponse, error) {
	log.Printf("IPsecListConns: Received: %v", in)

	ret, err := listConns(in)
	if err != nil {
		log.Printf("IPsecListConns: Failed: %v", err)
		return nil, err
	}

	return ret, nil
}

// IPsecListCerts executes the listCerts
func (s *Server) IPsecListCerts(_ context.Context, in *pb.IPsecListCertsRequest) (*pb.IPsecListCertsResponse, error) {
	log.Printf("IPsecListCerts: Received: %v", in)

	ret, err := listCerts(in)
	if err != nil {
		log.Printf("IPsecListConns: Failed: %v", err)
		return nil, err
	}

	return ret, nil
}

// IPsecLoadConn executes the loadConn
func (s *Server) IPsecLoadConn(_ context.Context, in *pb.IPsecLoadConnRequest) (*pb.IPsecLoadConnResponse, error) {
	log.Printf("IPsecLoadConn: Received: %v", in.GetConnection())

	err := loadConn(in)
	if err != nil {
		log.Printf("IPsecLoadConn: Failed %v", err)
		return nil, err
	}

	ret := pb.IPsecLoadConnResponse{
		Success: "Yes",
	}

	return &ret, nil
}

// IPsecUnloadConn executes the unloadConn
func (s *Server) IPsecUnloadConn(_ context.Context, in *pb.IPsecUnloadConnRequest) (*pb.IPsecUnloadConnResponse, error) {
	log.Printf("IPsecUnloadConn: Received: %v", in.GetName())

	err := unloadConn(in)
	if err != nil {
		log.Printf("IPsecUnloadConn: Failed %v", err)
		return nil, err
	}

	ret := pb.IPsecUnloadConnResponse{
		Success: "Yes",
	}

	return &ret, nil
}

*/
