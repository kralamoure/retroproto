// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type QuestsList struct{}

func NewQuestsList(extra string) (QuestsList, error) {
	var m QuestsList

	if err := m.Deserialize(extra); err != nil {
		return QuestsList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m QuestsList) MessageId() retroproto.MsgSvrId {
	return retroproto.QuestsList
}

func (m QuestsList) MessageName() string {
	return "QuestsList"
}

func (m QuestsList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *QuestsList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
