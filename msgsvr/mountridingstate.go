// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountRidingState struct{}

func (m MountRidingState) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountRidingState
}

func (m MountRidingState) Serialized() (string, error) {
	return "", nil
}

func (m *MountRidingState) Deserialize(extra string) error {
	return nil
}
