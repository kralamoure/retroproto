package msgcli

import (
	"fmt"
	"strconv"

	"github.com/kralamoure/retroproto"
)

type SpellsForget struct {
	Id int
}

func (m SpellsForget) MessageId() retroproto.MsgCliId {
	return retroproto.SpellsForget
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
