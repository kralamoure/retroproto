package msgsvr

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/kralamoure/retroproto"
)

type SpellsUpgradeSpellSuccess struct {
	Id    int
	Level int
}

func (m SpellsUpgradeSpellSuccess) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SpellsUpgradeSpellSuccess
}

func (m SpellsUpgradeSpellSuccess) Serialized() (string, error) {
	return fmt.Sprintf("%d~%d", m.Id, m.Level), nil
}

func (m *SpellsUpgradeSpellSuccess) Deserialize(extra string) error {
	sli := strings.Split(extra, "~")
	if len(sli) != 2 {
		return retroproto.ErrInvalidMsg
	}

	id, err := strconv.ParseInt(sli[0], 10, 32)
	if err != nil {
		return err
	}
	m.Id = int(id)

	level, err := strconv.ParseInt(sli[1], 10, 32)
	if err != nil {
		return err
	}
	m.Level = int(level)

	return nil
}
