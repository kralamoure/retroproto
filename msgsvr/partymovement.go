// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyMovement struct{}

func (m PartyMovement) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyMovement
}

func (m PartyMovement) MessageName() string {
	return "PartyMovement"
}

func (m PartyMovement) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyMovement) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
