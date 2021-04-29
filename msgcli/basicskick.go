// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsKick struct{}

func (m BasicsKick) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsKick
}

func (m BasicsKick) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsKick) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
