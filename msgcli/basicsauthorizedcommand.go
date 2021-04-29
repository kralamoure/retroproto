// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type BasicsAuthorizedCommand struct{}

func (m BasicsAuthorizedCommand) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.BasicsAuthorizedCommand
}

func (m BasicsAuthorizedCommand) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *BasicsAuthorizedCommand) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
