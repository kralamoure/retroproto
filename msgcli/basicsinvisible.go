// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsInvisible struct{}

func (m BasicsInvisible) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsInvisible
}

func (m BasicsInvisible) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsInvisible) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
