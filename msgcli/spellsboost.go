package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1proto"
)

type SpellsBoost struct {
	Id int
}

func (m SpellsBoost) ProtocolId() d1proto.MsgCliId {
	return d1proto.SpellsBoost
}

func (m SpellsBoost) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *SpellsBoost) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
