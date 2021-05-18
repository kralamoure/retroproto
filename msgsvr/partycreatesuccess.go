// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyCreateSuccess struct{}

func (m PartyCreateSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.PartyCreateSuccess
}

func (m PartyCreateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyCreateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
