// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type MountData struct{}

func (m MountData) ProtocolId() d1proto.MsgSvrId {
	return d1proto.MountData
}

func (m MountData) Serialized() (string, error) {
	return "", nil
}

func (m *MountData) Deserialize(extra string) error {
	return nil
}
