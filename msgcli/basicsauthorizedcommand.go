// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAuthorizedCommand struct{}

func (m BasicsAuthorizedCommand) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsAuthorizedCommand
}

func (m BasicsAuthorizedCommand) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAuthorizedCommand) Deserialize(extra string) error {
	return nil
}
