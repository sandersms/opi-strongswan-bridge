// SPDX-License-Identifier: Apache-2.0
// Copyright (c) 2024 Dell Inc, or its subsidiaries.

package ipsec

import "go.einride.tech/aip/resourcename"

func resourceIDToFullName(resourceID string) string {
	return resourcename.Join(
		"//security.opiproject.org/",
		"sa", resourceID,
	)
}