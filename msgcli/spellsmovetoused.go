// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type SpellsMoveToUsed struct{}

func (m SpellsMoveToUsed) ProtocolId() d1proto.MsgCliId {
	return d1proto.SpellsMoveToUsed
}

func (m SpellsMoveToUsed) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsMoveToUsed) Deserialize(extra string) error {
	return nil
}
