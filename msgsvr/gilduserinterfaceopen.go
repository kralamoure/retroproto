// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildUserInterfaceOpen struct{}

func (m GildUserInterfaceOpen) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildUserInterfaceOpen
}

func (m GildUserInterfaceOpen) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildUserInterfaceOpen) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
