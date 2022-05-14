// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyRequestLeave struct{}

func (m PartyRequestLeave) MessageId() retroproto.MsgCliId {
	return retroproto.PartyRequestLeave
}

func (m PartyRequestLeave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyRequestLeave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
