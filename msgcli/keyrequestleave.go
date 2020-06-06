// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type KeyRequestLeave struct{}

func (m KeyRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.KeyRequestLeave
}

func (m KeyRequestLeave) Serialized() (string, error) {
	return "", nil
}

func (m *KeyRequestLeave) Deserialize(extra string) error {
	return nil
}
