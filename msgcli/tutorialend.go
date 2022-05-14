// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type TutorialEnd struct{}

func NewTutorialEnd(extra string) (TutorialEnd, error) {
	var m TutorialEnd

	if err := m.Deserialize(extra); err != nil {
		return TutorialEnd{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m TutorialEnd) MessageId() retroproto.MsgCliId {
	return retroproto.TutorialEnd
}

func (m TutorialEnd) MessageName() string {
	return "TutorialEnd"
}

func (m TutorialEnd) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialEnd) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
