// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyRequestLeave struct{}

func (m PartyRequestLeave) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyRequestLeave
}

func (m PartyRequestLeave) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyRequestLeave) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
