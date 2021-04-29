// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismAttacked struct{}

func (m ConquestPrismAttacked) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismAttacked
}

func (m ConquestPrismAttacked) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismAttacked) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
