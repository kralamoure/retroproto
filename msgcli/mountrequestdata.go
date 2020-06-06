// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type MountRequestData struct{}

func (m MountRequestData) ProtocolId() d1proto.MsgCliId {
	return d1proto.MountRequestData
}

func (m MountRequestData) Serialized() (string, error) {
	return "", nil
}

func (m *MountRequestData) Deserialize(extra string) error {
	return nil
}
