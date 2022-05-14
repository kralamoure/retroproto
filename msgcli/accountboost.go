package msgcli

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retro/retrotyp"

	"github.com/kralamoure/retroproto"
)

type AccountBoost struct {
	CharacteristicId retrotyp.CharacteristicId
	Amount           int
}

func NewAccountBoost(extra string) (AccountBoost, error) {
	var m AccountBoost

	if err := m.Deserialize(extra); err != nil {
		return AccountBoost{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m AccountBoost) MessageId() retroproto.MsgCliId {
	return retroproto.AccountBoost
}

func (m AccountBoost) MessageName() string {
	return "AccountBoost"
}

func (m AccountBoost) Serialized() (string, error) {
	return fmt.Sprintf("%d|%d", m.CharacteristicId, m.Amount), nil
}

func (m *AccountBoost) Deserialize(extra string) error {
	sli := strings.Split(extra, "|")

	characteristicId, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.CharacteristicId = retrotyp.CharacteristicId(characteristicId)

	if len(sli) >= 2 {
		amount, err := strconv.ParseInt(sli[1], 10, 32)
		if err != nil {
			return err
		}
		m.Amount = int(amount)
	} else {
		m.Amount = 1
	}

	return nil
}
