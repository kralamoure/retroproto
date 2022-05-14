// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameActionsFinish struct{}

func NewGameActionsFinish(extra string) (GameActionsFinish, error) {
	var m GameActionsFinish

	if err := m.Deserialize(extra); err != nil {
		return GameActionsFinish{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameActionsFinish) MessageId() retroproto.MsgSvrId {
	return retroproto.GameActionsFinish
}

func (m GameActionsFinish) MessageName() string {
	return "GameActionsFinish"
}

func (m GameActionsFinish) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameActionsFinish) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
