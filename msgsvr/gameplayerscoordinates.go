// generated (unrevised)

package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type GamePlayersCoordinates struct{}

func (m GamePlayersCoordinates) MessageId() retroproto.MsgSvrId {
	return retroproto.GamePlayersCoordinates
}

func (m GamePlayersCoordinates) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *GamePlayersCoordinates) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
