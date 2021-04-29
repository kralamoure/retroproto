package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsUpgradeSpellError struct{}

func (m SpellsUpgradeSpellError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsUpgradeSpellError
}

func (m SpellsUpgradeSpellError) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsUpgradeSpellError) Deserialize(extra string) error {
	return nil
}
