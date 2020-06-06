// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SpellsForget struct{}

func (m SpellsForget) ProtocolId() d1proto.MsgCliId {
	return d1proto.SpellsForget
}

func (m SpellsForget) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsForget) Deserialize(extra string) error {
	return nil
}
