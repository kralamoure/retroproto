// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedLine struct{}

func NewBasicsAuthorizedLine(extra string) (BasicsAuthorizedLine, error) {
	var m BasicsAuthorizedLine

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedLine{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedLine) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedLine
}

func (m BasicsAuthorizedLine) MessageName() string {
	return "BasicsAuthorizedLine"
}

func (m BasicsAuthorizedLine) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedLine) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
