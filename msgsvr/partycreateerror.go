// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type PartyCreateError struct{}

func (m PartyCreateError) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyCreateError
}

func (m PartyCreateError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyCreateError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
