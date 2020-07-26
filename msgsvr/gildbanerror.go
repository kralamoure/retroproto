// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GildBanError struct{}

func (m GildBanError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GildBanError
}

func (m GildBanError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *GildBanError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
