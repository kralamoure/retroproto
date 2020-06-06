// generated (unrevised)
package msgsvr

import (
	"github.com/kralamoure/d1proto"
)

type KeyLeave struct{}

func (m KeyLeave) ProtocolId() d1proto.MsgSvrId {
	return d1proto.KeyLeave
}

func (m KeyLeave) Serialized() (string, error) {
	return "", nil
}

func (m *KeyLeave) Deserialize(extra string) error {
	return nil
}
