// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type SpellsList struct{}

func (m SpellsList) ProtocolId() d1proto.MsgSvrId {
	return d1proto.SpellsList
}

func (m SpellsList) Serialized() (string, error) {
	return "", nil
}

func (m *SpellsList) Deserialize(extra string) error {
	return nil
}
