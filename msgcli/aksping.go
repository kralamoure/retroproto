// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type AksPing struct{}

func (m AksPing) ProtocolId() d1proto.MsgCliId {
	return d1proto.AksPing
}

func (m AksPing) Serialized() (string, error) {
	return "", nil
}

func (m *AksPing) Deserialize(extra string) error {
	return nil
}
