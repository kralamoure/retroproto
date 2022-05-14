// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type ConquestPrismDead struct{}

func (m ConquestPrismDead) MessageId() retroproto.MsgSvrId {
	return retroproto.ConquestPrismDead
}

func (m ConquestPrismDead) MessageName() string {
	return "ConquestPrismDead"
}

func (m ConquestPrismDead) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestPrismDead) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
