package boostly

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/builtin/v9/market"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
)

//go:generate go run github.com/hannahhoward/cbor-gen-for@latest --map-encoding StorageAsk DealStatusRequest DealStatusResponse DealStatus DealParams Transfer DealResponse

const (
	FilStorageMarketProtocol_1_2_0 = "/fil/storage/mk/1.2.0"
)

// StorageAsk defines the parameters by which a miner will choose to accept or
// reject a deal. Note: making a storage deal proposal which matches the miner's
// ask is a precondition, but not sufficient to ensure the deal is accepted (the
// storage provider may run its own decision logic).
type StorageAsk struct {
	// Price per GiB / Epoch
	Price         abi.TokenAmount
	VerifiedPrice abi.TokenAmount

	MinPieceSize abi.PaddedPieceSize
	MaxPieceSize abi.PaddedPieceSize
	Miner        address.Address
}

type DealParams struct {
	DealUUID           uuid.UUID
	IsOffline          bool
	ClientDealProposal market.ClientDealProposal
	DealDataRoot       cid.Cid
	Transfer           Transfer // Transfer params will be the zero value if this is an offline deal
	RemoveUnsealedCopy bool
	SkipIPNIAnnounce   bool
}

// Transfer has the parameters for a data transfer
type Transfer struct {
	// The type of transfer eg "http"
	Type string
	// An optional ID that can be supplied by the client to identify the deal
	ClientID string
	// A byte array containing marshalled data specific to the transfer type
	// eg a JSON encoded struct { URL: "<url>", Headers: {...} }
	Params []byte
	// The size of the data transferred in bytes
	Size uint64
}

type DealResponse struct {
	Accepted bool
	// Message is the reason the deal proposal was rejected. It is empty if
	// the deal was accepted.
	Message string
}
