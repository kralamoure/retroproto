// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type GamePlayersCoordinates struct{}

func (m GamePlayersCoordinates) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.GamePlayersCoordinates
}

func (m GamePlayersCoordinates) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *GamePlayersCoordinates) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
