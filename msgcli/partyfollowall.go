// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1encoding"
)

type PartyFollowAll struct{}

func (m PartyFollowAll) ProtocolId() d1encoding.MsgCliId {
	return d1encoding.PartyFollowAll
}

func (m PartyFollowAll) Serialized() (string, error) {
	return "", d1encoding.ErrNotImplemented
}

func (m *PartyFollowAll) Deserialize(extra string) error {
	return d1encoding.ErrNotImplemented
}
