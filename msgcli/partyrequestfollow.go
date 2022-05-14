// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyRequestFollow struct{}

func (m PartyRequestFollow) MessageId() retroproto.MsgCliId {
	return retroproto.PartyRequestFollow
}

func (m PartyRequestFollow) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRequestFollow) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
