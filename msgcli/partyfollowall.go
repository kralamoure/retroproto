// generated (unrevised)

package msgcli

import (
	"github.com/kralamoure/retroproto"
)

type PartyFollowAll struct{}

func (m PartyFollowAll) ProtocolId() retroproto.MsgCliId {
	return retroproto.PartyFollowAll
}

func (m PartyFollowAll) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyFollowAll) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
