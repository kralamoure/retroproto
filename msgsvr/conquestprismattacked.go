// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismAttacked struct{}

func (m ConquestPrismAttacked) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismAttacked
}

func (m ConquestPrismAttacked) MessageName() string {
	return "ConquestPrismAttacked"
}

func (m ConquestPrismAttacked) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismAttacked) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
