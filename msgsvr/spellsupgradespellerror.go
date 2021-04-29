package msgsvr

import (
	"github.com/kralamoure/d1encoding"
)

type SpellsUpgradeSpellError struct{}

func (m SpellsUpgradeSpellError) ProtocolId() d1encoding.MsgSvrId {
	return d1encoding.SpellsUpgradeSpellError
}

func (m SpellsUpgradeSpellError) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsUpgradeSpellError) Deserialize(extra string) error {
	return nil
}
