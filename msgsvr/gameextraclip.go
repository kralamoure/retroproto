// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameExtraClip struct{}

func NewGameExtraClip(extra string) (GameExtraClip, error) {
	var m GameExtraClip

	if err := m.Deserialize(extra); err != nil {
		return GameExtraClip{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameExtraClip) MessageId() retroproto.MsgSvrId {
	return retroproto.GameExtraClip
}

func (m GameExtraClip) MessageName() string {
	return "GameExtraClip"
}

func (m GameExtraClip) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameExtraClip) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
