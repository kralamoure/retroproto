// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GildBanSuccess struct{}

func (m GildBanSuccess) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GildBanSuccess
}

func (m GildBanSuccess) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GildBanSuccess) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
