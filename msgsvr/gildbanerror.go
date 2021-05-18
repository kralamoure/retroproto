// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildBanError struct{}

func (m GildBanError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.GildBanError
}

func (m GildBanError) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildBanError) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
