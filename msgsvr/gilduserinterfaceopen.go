// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildUserInterfaceOpen struct{}

func (m GildUserInterfaceOpen) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildUserInterfaceOpen
}

func (m GildUserInterfaceOpen) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GildUserInterfaceOpen) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
