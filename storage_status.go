package boostly

import (
	"context"
	"fmt"
	"time"

	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/filecoin-project/go-state-types/crypto"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	host "github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
)

//go:generate go run github.com/hannahhoward/cbor-gen-for@latest --map-encoding DealStatusRequest DealStatusResponse DealStatus

const FilStorageStatusProtocol_1_2_0 = "/fil/storage/status/1.2.0"

// DealStatusRequest is sent to get the current state of a deal from a
// storage provider
type DealStatusRequest struct {
	DealUUID  uuid.UUID
	Signature crypto.Signature
}

// DealStatusResponse is the current state of a deal
type DealStatusResponse struct {
	DealUUID uuid.UUID
	// Error is non-empty if there is an error getting the deal status
	// (eg invalid request signature)
	Error          string
	DealStatus     *DealStatus
	IsOffline      bool
	TransferSize   uint64
	NBytesReceived uint64
}

type DealStatus struct {
	// Error is non-empty if the deal is in the error state
	Error string
	// Status is a string corresponding to a deal checkpoint
	Status string
	// SealingStatus is the sealing status reported by lotus miner
	SealingStatus string
	// Proposal is the deal proposal
	Proposal market.DealProposal
	// SignedProposalCid is the cid of the client deal proposal + signature
	SignedProposalCid cid.Cid
	// PublishCid is the cid of the Publish message sent on chain, if the deal
	// has reached the publish stage
	PublishCid *cid.Cid
	// ChainDealID is the id of the deal in chain state
	ChainDealID abi.DealID
}

func GetDealStatus(ctx context.Context, h host.Host, id peer.ID, dealUUID uuid.UUID, signer func(context.Context, []byte) (*crypto.Signature, error)) (*DealStatusResponse, error) {
	uuidBytes, err := dealUUID.MarshalBinary()
	if err != nil {
		return nil, fmt.Errorf("getting uuid bytes: %w", err)
	}

	sig, err := signer(ctx, uuidBytes)
	if err != nil {
		return nil, fmt.Errorf("signing uuid bytes: %w", err)
	}

	// Create a libp2p stream to the provider
	s, err := h.NewStream(ctx, id, FilStorageStatusProtocol_1_2_0)
	if err != nil {
		return nil, err
	}
	defer s.Close()

	// Set a deadline on writing to the stream so it doesn't hang
	_ = s.SetWriteDeadline(time.Now().Add(10 * time.Second)) // TODO parameterise
	defer s.SetWriteDeadline(time.Time{})                    // nolint

	// Write the deal status request to the stream
	req := DealStatusRequest{DealUUID: dealUUID, Signature: *sig}
	if err = cborutil.WriteCborRPC(s, &req); err != nil {
		return nil, fmt.Errorf("sending deal status req: %w", err)
	}

	// Set a deadline on reading from the stream so it doesn't hang
	_ = s.SetReadDeadline(time.Now().Add(10 * time.Second)) // TODO parameterise
	defer s.SetReadDeadline(time.Time{})

	// Read the response from the stream
	var resp DealStatusResponse
	if err := resp.UnmarshalCBOR(s); err != nil {
		return nil, fmt.Errorf("reading deal status response: %w", err)
	}
	return &resp, nil
}
