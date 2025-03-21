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

package blob

import (
	"context"
	"time"

	ctypes "github.com/berachain/beacon-kit/consensus-types/types"
	"github.com/berachain/beacon-kit/da/kzg"
	dastore "github.com/berachain/beacon-kit/da/store"
	datypes "github.com/berachain/beacon-kit/da/types"
	"github.com/berachain/beacon-kit/log"
	"github.com/berachain/beacon-kit/primitives/common"
	"github.com/berachain/beacon-kit/primitives/eip4844"
	"github.com/berachain/beacon-kit/primitives/math"
)

// Processor is the blob processor that handles the processing and verification
// of blob sidecars.
type Processor struct {
	// logger is used to log information and errors.
	logger log.Logger
	// verifier is responsible for verifying the blobs.
	verifier *verifier
	// metrics is used to collect and report processor metrics.
	metrics *processorMetrics
}

// NewProcessor creates a new blob processor.
func NewProcessor(
	logger log.Logger,
	proofVerifier kzg.BlobProofVerifier,
	telemetrySink TelemetrySink,
) *Processor {
	verifier := newVerifier(proofVerifier, telemetrySink)

	return &Processor{
		logger:   logger,
		verifier: verifier,
		metrics:  newProcessorMetrics(telemetrySink),
	}
}

// VerifySidecars verifies the blobs and ensures they match the local state.
func (sp *Processor) VerifySidecars(
	ctx context.Context,
	sidecars datypes.BlobSidecars,
	blkHeader *ctypes.BeaconBlockHeader,
	kzgCommitments eip4844.KZGCommitments[common.ExecutionHash],
) error {
	defer sp.metrics.measureVerifySidecarsDuration(
		time.Now(), math.U64(len(sidecars)),
	)

	// Abort if there are no blobs to store.
	if len(sidecars) == 0 {
		return nil
	}

	// Verify the blobs and ensure they match the local state.
	return sp.verifier.verifySidecars(
		ctx,
		sidecars,
		blkHeader,
		kzgCommitments,
	)
}

// ProcessSidecars processes the blobs and ensures they match the local state.
func (sp *Processor) ProcessSidecars(
	avs *dastore.Store,
	sidecars datypes.BlobSidecars,
) error {
	defer sp.metrics.measureProcessSidecarsDuration(
		time.Now(), math.U64(len(sidecars)),
	)

	// Abort if there are no blobs to store.
	if len(sidecars) == 0 {
		return nil
	}

	// If we have reached this point, we can safely assume that the blobs are
	// valid and can be persisted, as well as that index 0 is filled.
	return avs.Persist(sidecars)
}
