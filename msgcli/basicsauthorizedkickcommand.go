// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedKickCommand struct{}

func NewBasicsAuthorizedKickCommand(extra string) (BasicsAuthorizedKickCommand, error) {
	var m BasicsAuthorizedKickCommand

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedKickCommand{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedKickCommand) MessageId() retroproto.MsgCliId {
	return retroproto.BasicsAuthorizedKickCommand
}

func (m BasicsAuthorizedKickCommand) MessageName() string {
	return "BasicsAuthorizedKickCommand"
}

func (m BasicsAuthorizedKickCommand) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedKickCommand) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
