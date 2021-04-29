package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1/d1typ"

	"github.com/kralamoure/d1proto"
)

type AccountBoost struct {
	CharacteristicId d1typ.CharacteristicId
}

func (m AccountBoost) ProtocolId() d1proto.MsgCliId {
	return d1proto.AccountBoost
}

func (m AccountBoost) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.CharacteristicId), nil
}

func (m *AccountBoost) Deserialize(extra string) error {
	characteristicId, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.CharacteristicId = d1typ.CharacteristicId(characteristicId)

	return nil
}
