// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommand struct{}

func (m BasicsAuthorizedCommand) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAuthorizedCommand
}

func (m BasicsAuthorizedCommand) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommand) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
