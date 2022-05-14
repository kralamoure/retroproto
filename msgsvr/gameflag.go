// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFlag struct{}

func NewGameFlag(extra string) (GameFlag, error) {
	var m GameFlag

	if err := m.Deserialize(extra); err != nil {
		return GameFlag{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFlag) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFlag
}

func (m GameFlag) MessageName() string {
	return "GameFlag"
}

func (m GameFlag) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFlag) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
