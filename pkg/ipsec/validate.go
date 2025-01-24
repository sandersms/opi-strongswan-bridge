// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2024 Dell Inc, or its subsidiaries.

package ipsec

import (
	"go.einride.tech/aip/fieldbehavior"
	"go.einride.tech/aip/resourceid"

	pb "github.com/opiproject/opi-api/security/v1alpha1/gen/go"
)

func (s *Server) validateCreateIpsecSaRequest(in *pb.CreateIpsecSaRequest) error {
	// check required fields
	if err := fieldbehavior.ValidateRequiredFields(in); err != nil {
		return err
	}

	// see https://google.aip.dev/133#user-specified-ids
	if in.IpsecSaId != "" {
		if err := resourceid.ValidateUserSettable(in.IpsecSaId); err != nil {
			return err
		}
	}
	return nil
}
