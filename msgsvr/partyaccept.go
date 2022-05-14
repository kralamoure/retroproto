// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyAccept struct{}

func (m PartyAccept) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyAccept
}

func (m PartyAccept) MessageName() string {
	return "PartyAccept"
}

func (m PartyAccept) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyAccept) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
