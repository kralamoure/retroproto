// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type TutorialGameBegin struct{}

func NewTutorialGameBegin(extra string) (TutorialGameBegin, error) {
	var m TutorialGameBegin

	if err := m.Deserialize(extra); err != nil {
		return TutorialGameBegin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m TutorialGameBegin) MessageId() retroproto.MsgSvrId {
	return retroproto.TutorialGameBegin
}

func (m TutorialGameBegin) MessageName() string {
	return "TutorialGameBegin"
}

func (m TutorialGameBegin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *TutorialGameBegin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
