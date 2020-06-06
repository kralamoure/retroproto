// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SpellsBoost struct{}

func (m SpellsBoost) ProtocolId() d1proto.MsgCliId {
	return d1proto.SpellsBoost
}

func (m SpellsBoost) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsBoost) Deserialize(extra string) error {
	return nil
}
