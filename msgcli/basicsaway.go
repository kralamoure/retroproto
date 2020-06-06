// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsAway struct{}

func (m BasicsAway) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsAway
}

func (m BasicsAway) Serialized() (string, error) {
	return "", nil
}

func (m *BasicsAway) Deserialize(extra string) error {
	return nil
}
