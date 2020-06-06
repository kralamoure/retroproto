// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type HousesLockedProperty struct{}

func (m HousesLockedProperty) ProtocolId() d1proto.MsgSvrId {
	return d1proto.HousesLockedProperty
}

func (m HousesLockedProperty) Serialized() (string, error) {
	return "", nil
}

func (m *HousesLockedProperty) Deserialize(extra string) error {
	return nil
}
