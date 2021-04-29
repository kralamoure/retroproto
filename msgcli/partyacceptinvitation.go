// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyAcceptInvitation struct{}

func (m PartyAcceptInvitation) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyAcceptInvitation
}

func (m PartyAcceptInvitation) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyAcceptInvitation) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
