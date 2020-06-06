// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type HousesRequestLeave struct{}

func (m HousesRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.HousesRequestLeave
}

func (m HousesRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *HousesRequestLeave) Deserialize(extra string) error {
	return nil
}
