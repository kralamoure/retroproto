// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GildBanSuccess struct{}

func (m GildBanSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.GildBanSuccess
}

func (m GildBanSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GildBanSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
