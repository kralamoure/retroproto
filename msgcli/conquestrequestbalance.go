// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type ConquestRequestBalance struct{}

func (m ConquestRequestBalance) ProtocolId() d1proto.MsgCliId {
	return d1proto.ConquestRequestBalance
}

func (m ConquestRequestBalance) Serialized() (string, error) {
	return "", nil
}

func (m *ConquestRequestBalance) Deserialize(extra string) error {
	return nil
}
