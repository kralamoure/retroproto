// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type TutorialCreate struct{}

func NewTutorialCreate(extra string) (TutorialCreate, error) {
	var m TutorialCreate

	if err := m.Deserialize(extra); err != nil {
		return TutorialCreate{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m TutorialCreate) MessageId() retroproto.MsgSvrId {
	return retroproto.TutorialCreate
}

func (m TutorialCreate) MessageName() string {
	return "TutorialCreate"
}

func (m TutorialCreate) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialCreate) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
