// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyAcceptInvitation struct{}

func (m PartyAcceptInvitation) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyAcceptInvitation
}

func (m PartyAcceptInvitation) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyAcceptInvitation) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
