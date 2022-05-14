// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFrameObjectExternal struct{}

func NewGameFrameObjectExternal(extra string) (GameFrameObjectExternal, error) {
	var m GameFrameObjectExternal

	if err := m.Deserialize(extra); err != nil {
		return GameFrameObjectExternal{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFrameObjectExternal) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFrameObjectExternal
}

func (m GameFrameObjectExternal) MessageName() string {
	return "GameFrameObjectExternal"
}

func (m GameFrameObjectExternal) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFrameObjectExternal) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
