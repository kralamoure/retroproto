// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsKick struct{}

func (m BasicsKick) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsKick
}

func (m BasicsKick) MessageName() string {
	return "BasicsKick"
}

func (m BasicsKick) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsKick) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
