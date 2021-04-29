// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyRequestLeave struct{}

func (m PartyRequestLeave) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyRequestLeave
}

func (m PartyRequestLeave) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyRequestLeave) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
