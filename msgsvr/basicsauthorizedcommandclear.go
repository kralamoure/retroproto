// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandClear struct{}

func NewBasicsAuthorizedCommandClear(extra string) (BasicsAuthorizedCommandClear, error) {
	var m BasicsAuthorizedCommandClear

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedCommandClear{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedCommandClear) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandClear
}

func (m BasicsAuthorizedCommandClear) MessageName() string {
	return "BasicsAuthorizedCommandClear"
}

func (m BasicsAuthorizedCommandClear) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandClear) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
