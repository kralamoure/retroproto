package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type GameGetExtraInformations struct{}

func (m GameGetExtraInformations) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.GameGetExtraInformations
}

func (m GameGetExtraInformations) Serialized() (string, error) {
	return "", nil
}

func (m *GameGetExtraInformations) Deserialize(extra string) error {
	return nil
}
