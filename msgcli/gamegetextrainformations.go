package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type GameGetExtraInformations struct{}

func (m GameGetExtraInformations) ProtocolId() d1proto.MsgCliId {
	return d1proto.GameGetExtraInformations
}

func (m GameGetExtraInformations) Serialized() (string, error) {
	return "", nil
}

func (m *GameGetExtraInformations) Deserialize(extra string) error {
	return nil
}
