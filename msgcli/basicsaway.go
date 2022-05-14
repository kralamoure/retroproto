// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAway struct{}

func (m BasicsAway) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAway
}

func (m BasicsAway) MessageName() string {
	return "BasicsAway"
}

func (m BasicsAway) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAway) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
