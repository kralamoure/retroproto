package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/d1encoding"
)

type SpellsForget struct {
	Id int
}

func (m SpellsForget) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.SpellsForget
}

func (m SpellsForget) Serialized() (string, error) {
	return fmt.Sprintf("%d", m.Id), nil
}

func (m *SpellsForget) Deserialize(extra string) error {
	id, err := strconv.ParseInt(extra, 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	return nil
}
