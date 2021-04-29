// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type ConquestPrismAttacked struct{}

func (m ConquestPrismAttacked) ProtocolId() d1proto.MsgSvrId {
	return d1proto.ConquestPrismAttacked
}

func (m ConquestPrismAttacked) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *ConquestPrismAttacked) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
