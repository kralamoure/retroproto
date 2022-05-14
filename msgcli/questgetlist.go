// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type QuestGetList struct{}

func NewQuestGetList(extra string) (QuestGetList, error) {
	var m QuestGetList

	if err := m.Deserialize(extra); err != nil {
		return QuestGetList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m QuestGetList) MessageId() retroproto.MsgCliId {
	return retroproto.QuestGetList
}

func (m QuestGetList) MessageName() string {
	return "QuestGetList"
}

func (m QuestGetList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestGetList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
