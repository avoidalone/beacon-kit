// SPDX-License-Identifier: BUSL-1.1
//
// Copyright (C) 2025, Berachain Foundation. All rights reserved.
// Use of this software is governed by the Business Source License included
// in the LICENSE file of this repository and at www.mariadb.com/bsl11.
//
// ANY USE OF THE LICENSED WORK IN VIOLATION OF THIS LICENSE WILL AUTOMATICALLY
// TERMINATE YOUR RIGHTS UNDER THIS LICENSE FOR THE CURRENT AND ALL OTHER
// VERSIONS OF THE LICENSED WORK.
//
// THIS LICENSE DOES NOT GRANT YOU ANY RIGHT IN ANY TRADEMARK OR LOGO OF
// LICENSOR OR ITS AFFILIATES (PROVIDED THAT YOU MAY USE A TRADEMARK OR LOGO OF
// LICENSOR AS EXPRESSLY REQUIRED BY THIS LICENSE).
//
// TO THE EXTENT PERMITTED BY APPLICABLE LAW, THE LICENSED WORK IS PROVIDED ON
// AN “AS IS” BASIS. LICENSOR HEREBY DISCLAIMS ALL WARRANTIES AND CONDITIONS,
// EXPRESS OR IMPLIED, INCLUDING (WITHOUT LIMITATION) WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, NON-INFRINGEMENT, AND
// TITLE.

package chain_test

import (
	"testing"

	"github.com/berachain/beacon-kit/chain"
	"github.com/berachain/beacon-kit/primitives/common"
	"github.com/berachain/beacon-kit/primitives/math"
	"github.com/berachain/beacon-kit/primitives/version"
	"github.com/stretchr/testify/require"
)

// TODO: Add setupValid, setupInvalid functions and use in each test.

// Create an instance of chainSpec with test data.
var spec, _ = chain.NewSpec(
	&chain.SpecData{
		Deneb1ForkEpoch:                  9,
		ElectraForkEpoch:                 10,
		SlotsPerEpoch:                    32,
		MinEpochsForBlobsSidecarsRequest: 5,
		MaxWithdrawalsPerPayload:         2,
	},
)

// TestActiveForkVersionForEpoch tests the ActiveForkVersionForEpoch method.
func TestActiveForkVersionForEpoch(t *testing.T) {
	t.Parallel()
	// Define test cases
	tests := []struct {
		name     string
		epoch    math.Epoch
		expected common.Version
	}{
		{name: "Before Electra Fork", epoch: 9, expected: version.Deneb1()},
		{name: "At Electra Fork", epoch: 10, expected: version.Electra()},
		{name: "After Electra Fork", epoch: 11, expected: version.Electra()},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := spec.ActiveForkVersionForEpoch(tt.epoch)
			require.Equal(t, tt.expected, result, "Test case : %s", tt.name)
		})
	}
}

// TestSlotToEpoch tests the SlotToEpoch method.
func TestSlotToEpoch(t *testing.T) {
	t.Parallel()
	// Define test cases
	tests := []struct {
		name     string
		slot     math.Slot
		expected math.Epoch
	}{
		{name: "Epoch 0, Slot 0", slot: 0, expected: 0},
		{name: "Epoch 0, Slot 31", slot: 31, expected: 0},
		{name: "Epoch 1, Slot 32", slot: 32, expected: 1},
		{name: "Epoch 1, Slot 63", slot: 63, expected: 1},
		{name: "Epoch 2, Slot 64", slot: 64, expected: 2},
		{name: "Epoch 2, Slot 95", slot: 95, expected: 2},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := spec.SlotToEpoch(tt.slot)
			require.Equal(t, tt.expected, result, "Test case : %s", tt.name)
		})
	}
}

// TestActiveForkVersionForSlot tests the ActiveForkVersionForSlot method.
func TestActiveForkVersionForSlot(t *testing.T) {
	t.Parallel()
	// Define test cases
	tests := []struct {
		name     string
		slot     math.Slot
		expected common.Version
	}{
		{name: "Before Electra Fork", slot: 0, expected: version.Deneb()},
		{
			name:     "Just Before Electra Fork",
			slot:     319,
			expected: version.Deneb1(),
		},
		{name: "At Electra Fork", slot: 320, expected: version.Electra()},
		{name: "After Electra Fork", slot: 640, expected: version.Electra()},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := spec.ActiveForkVersionForSlot(tt.slot)
			require.Equal(t, tt.expected, result, "Test case : %s", tt.name)
		})
	}
}

// TestWithinDAPeriod tests the WithinDAPeriod method.
func TestWithinDAPeriod(t *testing.T) {
	t.Parallel()
	// Define test cases
	tests := []struct {
		name     string
		block    math.Slot
		current  math.Slot
		expected bool
	}{
		// Block is within DA period (5 epochs).
		{name: "Within DA Period", block: 0, current: 160, expected: true},
		// Block is outside DA period (>5 epochs).
		{name: "Outside DA Period", block: 0, current: 192, expected: false},
		// Block is within DA period.
		{name: "Within DA Period 2", block: 160, current: 320, expected: true},
		// Block is outside DA period.
		{
			name:     "Outside DA Period 2",
			block:    160,
			current:  352,
			expected: false,
		},
	}

	// Run test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := spec.WithinDAPeriod(tt.block, tt.current)
			require.Equal(t, tt.expected, result, "Test case : %s", tt.name)
		})
	}
}
