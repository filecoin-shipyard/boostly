package boostly

import (
	"context"
	"fmt"

	"github.com/ipld/go-ipld-prime/codec/dagcbor"
	"github.com/ipld/go-ipld-prime/node/bindnode"
	"github.com/ipld/go-ipld-prime/node/bindnode/registry"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
)

const (
	FilRetrievalTransportsProtocol_1_0_0 = protocol.ID("/fil/retrieval/transports/1.0.0")

	transporsIpldSchema = `
type Multiaddr bytes
type Protocol struct {
  Name String
  Addresses [Multiaddr]
}
type TransportsQueryResponse struct {
  Protocols [Protocol]
}`
)

type TransportsQueryResponse struct {
	Protocols []struct {
		Name      string                `json:"name,omitempty"`
		Addresses []multiaddr.Multiaddr `json:"addresses,omitempty"`
	} `json:"protocols,omitempty"`
}

var reg = registry.NewRegistry()

func init() {
	if err := reg.RegisterType(
		(*TransportsQueryResponse)(nil),
		transporsIpldSchema,
		"TransportsQueryResponse",
		bindnode.TypedBytesConverter((*multiaddr.Multiaddr)(nil), func(b []byte) (any, error) {
			switch ma, err := multiaddr.NewMultiaddrBytes(b); {
			case err != nil:
				return nil, err
			default:
				return &ma, err
			}
		}, func(v any) ([]byte, error) {
			switch ma, ok := v.(*multiaddr.Multiaddr); {
			case !ok:
				return nil, fmt.Errorf("expected *Multiaddr value")
			default:
				return (*ma).Bytes(), nil
			}
		}),
	); err != nil {
		panic(err)
	}
}

func QueryTransports(ctx context.Context, h host.Host, id peer.ID) (*TransportsQueryResponse, error) {
	stream, err := h.NewStream(ctx, id, FilRetrievalTransportsProtocol_1_0_0)
	if err != nil {
		return nil, err
	}
	defer func() { _ = stream.Close() }()
	if resp, err := reg.TypeFromReader(stream, (*TransportsQueryResponse)(nil), dagcbor.Decode); err != nil {
		return nil, err
	} else if qResp, ok := resp.(*TransportsQueryResponse); !ok {
		return nil, fmt.Errorf("expected TransportsQueryResponse but got %v", resp)
	} else {
		return qResp, nil
	}
}
