package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type AksHelloGame struct{}

func NewAksHelloGame(extra string) (AksHelloGame, error) {
	var m AksHelloGame

	if err := m.Deserialize(extra); err != nil {
		return AksHelloGame{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AksHelloGame) MessageId() retroproto.MsgSvrId {
	return retroproto.AksHelloGame
}

func (m AksHelloGame) MessageName() string {
	return "AksHelloGame"
}

func (m AksHelloGame) Serialized() (string, error) {
	return "", nil
}

func (m *AksHelloGame) Deserialize(extra string) error {
	return nil
}
