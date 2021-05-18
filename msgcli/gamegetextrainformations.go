package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type GameGetExtraInformations struct{}

func (m GameGetExtraInformations) ProtocolId() retroproto.MsgCliId {
	return retroproto.GameGetExtraInformations
}

func (m GameGetExtraInformations) Serialized() (string, error) {
	return "", nil
}

func (m *GameGetExtraInformations) Deserialize(extra string) error {
	return nil
}
