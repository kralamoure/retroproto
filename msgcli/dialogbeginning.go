// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type DialogBeginning struct{}

func (m DialogBeginning) ProtocolId() d1proto.MsgCliId {
	return d1proto.DialogBeginning
}

func (m DialogBeginning) Serialized() (string, error) {
	return "", nil
}

func (m *DialogBeginning) Deserialize(extra string) error {
	return nil
}
