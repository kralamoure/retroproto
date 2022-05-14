// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyLeave struct{}

func (m PartyLeave) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyLeave
}

func (m PartyLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
