// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type BasicsSanctionMe struct{}

func (m BasicsSanctionMe) ProtocolId() d1proto.MsgCliId {
	return d1proto.BasicsSanctionMe
}

func (m BasicsSanctionMe) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *BasicsSanctionMe) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
