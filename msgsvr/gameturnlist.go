// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameTurnList struct{}

func NewGameTurnList(extra string) (GameTurnList, error) {
	var m GameTurnList

	if err := m.Deserialize(extra); err != nil {
		return GameTurnList{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameTurnList) MessageId() retroproto.MsgSvrId {
	return retroproto.GameTurnList
}

func (m GameTurnList) MessageName() string {
	return "GameTurnList"
}

func (m GameTurnList) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameTurnList) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
