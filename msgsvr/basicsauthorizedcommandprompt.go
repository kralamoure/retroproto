// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type BasicsAuthorizedCommandPrompt struct{}

func NewBasicsAuthorizedCommandPrompt(extra string) (BasicsAuthorizedCommandPrompt, error) {
	var m BasicsAuthorizedCommandPrompt

	if err := m.Deserialize(extra); err != nil {
		return BasicsAuthorizedCommandPrompt{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m BasicsAuthorizedCommandPrompt) MessageId() retroproto.MsgSvrId {
	return retroproto.BasicsAuthorizedCommandPrompt
}

func (m BasicsAuthorizedCommandPrompt) MessageName() string {
	return "BasicsAuthorizedCommandPrompt"
}

func (m BasicsAuthorizedCommandPrompt) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *BasicsAuthorizedCommandPrompt) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
