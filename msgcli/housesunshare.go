// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesUnShare struct{}

func (m HousesUnShare) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesUnShare
}

func (m HousesUnShare) Serialized() (string, error) {
	return "", nil
}

func (m *HousesUnShare) Deserialize(extra string) error {
	return nil
}
