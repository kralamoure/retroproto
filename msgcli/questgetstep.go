// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type QuestGetStep struct{}

func NewQuestGetStep(extra string) (QuestGetStep, error) {
	var m QuestGetStep

	if err := m.Deserialize(extra); err != nil {
		return QuestGetStep{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m QuestGetStep) MessageId() retroproto.MsgCliId {
	return retroproto.QuestGetStep
}

func (m QuestGetStep) MessageName() string {
	return "QuestGetStep"
}

func (m QuestGetStep) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestGetStep) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
