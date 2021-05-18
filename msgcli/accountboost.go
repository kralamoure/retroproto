package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type AccountBoost struct {
	CharacteristicId retrotyp.CharacteristicId
}

func (m AccountBoost) ProtocolId() retroproto.MsgCliId {
	return retroproto.AccountBoost
}

func (m AccountBoost) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.CharacteristicId), nil
}

func (m *AccountBoost) Deserialize(extra string) error {
	characteristicId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.CharacteristicId = retrotyp.CharacteristicId(characteristicId)

	return nil
}
