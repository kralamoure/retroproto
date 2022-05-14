// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type QuestsStep struct{}

func NewQuestsStep(extra string) (QuestsStep, error) {
	var m QuestsStep

	if err := m.Deserialize(extra); err != nil {
		return QuestsStep{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m QuestsStep) MessageId() retroproto.MsgSvrId {
	return retroproto.QuestsStep
}

func (m QuestsStep) MessageName() string {
	return "QuestsStep"
}

func (m QuestsStep) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestsStep) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
