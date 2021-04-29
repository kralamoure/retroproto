// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildBanError struct{}

func (m GildBanError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildBanError
}

func (m GildBanError) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildBanError) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
