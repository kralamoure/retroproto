// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type GameFightOption struct{}

func NewGameFightOption(extra string) (GameFightOption, error) {
	var m GameFightOption

	if err := m.Deserialize(extra); err != nil {
		return GameFightOption{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m GameFightOption) MessageId() retroproto.MsgSvrId {
	return retroproto.GameFightOption
}

func (m GameFightOption) MessageName() string {
	return "GameFightOption"
}

func (m GameFightOption) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GameFightOption) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
