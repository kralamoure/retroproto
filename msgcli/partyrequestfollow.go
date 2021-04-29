// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyRequestFollow struct{}

func (m PartyRequestFollow) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyRequestFollow
}

func (m PartyRequestFollow) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyRequestFollow) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
