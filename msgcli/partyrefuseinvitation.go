// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyRefuseInvitation struct{}

func (m PartyRefuseInvitation) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyRefuseInvitation
}

func (m PartyRefuseInvitation) Serialized() (string, error) {
	return "", nil
}

func (m *PartyRefuseInvitation) Deserialize(extra string) error {
	return nil
}
