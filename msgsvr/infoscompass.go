// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type InfosCompass struct{}

func (m InfosCompass) MessageId() retroproto.MsgSvrId {
	return retroproto.InfosCompass
}

func (m InfosCompass) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *InfosCompass) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
