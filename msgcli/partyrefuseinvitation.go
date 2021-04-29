// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyRefuseInvitation struct{}

func (m PartyRefuseInvitation) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyRefuseInvitation
}

func (m PartyRefuseInvitation) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyRefuseInvitation) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
