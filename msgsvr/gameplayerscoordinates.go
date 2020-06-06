// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type GamePlayersCoordinates struct{}

func (m GamePlayersCoordinates) ProtocolId() d1proto.MsgSvrId {
	return d1proto.GamePlayersCoordinates
}

func (m GamePlayersCoordinates) Serialized() (string, error) {
	return "", nil
}

func (m *GamePlayersCoordinates) Deserialize(extra string) error {
	return nil
}
