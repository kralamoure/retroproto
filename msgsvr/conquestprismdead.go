// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type ConquestPrismDead struct{}

func (m ConquestPrismDead) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.ConquestPrismDead
}

func (m ConquestPrismDead) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *ConquestPrismDead) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
