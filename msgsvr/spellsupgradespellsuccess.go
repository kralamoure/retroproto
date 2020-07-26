// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsUpgradeSpellSuccess struct{}

func (m SpellsUpgradeSpellSuccess) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsUpgradeSpellSuccess
}

func (m SpellsUpgradeSpellSuccess) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *SpellsUpgradeSpellSuccess) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
