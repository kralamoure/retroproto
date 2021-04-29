// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyRequestFollow struct{}

func (m PartyRequestFollow) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyRequestFollow
}

func (m PartyRequestFollow) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyRequestFollow) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
