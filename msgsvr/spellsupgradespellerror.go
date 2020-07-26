// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsUpgradeSpellError struct{}

func (m SpellsUpgradeSpellError) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsUpgradeSpellError
}

func (m SpellsUpgradeSpellError) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SpellsUpgradeSpellError) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
