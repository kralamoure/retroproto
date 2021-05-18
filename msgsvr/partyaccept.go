// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyAccept struct{}

func (m PartyAccept) ProtocolId() retroproto.MsgSvrId {
	return retroproto.PartyAccept
}

func (m PartyAccept) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyAccept) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
