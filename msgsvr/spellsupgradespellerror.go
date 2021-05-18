package msgsvr

import (
	"github.com/kralamoure/retroproto"
)

type SpellsUpgradeSpellError struct{}

func (m SpellsUpgradeSpellError) ProtocolId() retroproto.MsgSvrId {
	return retroproto.SpellsUpgradeSpellError
}

func (m SpellsUpgradeSpellError) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsUpgradeSpellError) Deserialize(extra string) error {
	return nil
}
