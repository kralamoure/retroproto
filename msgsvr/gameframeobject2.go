// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFrameObject2 struct{}

func NewGameFrameObject2(extra string) (GameFrameObject2, error) {
	var m GameFrameObject2

	if err := m.Deserialize(extra); err != nil {
		return GameFrameObject2{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFrameObject2) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFrameObject2
}

func (m GameFrameObject2) MessageName() string {
	return "GameFrameObject2"
}

func (m GameFrameObject2) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFrameObject2) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
