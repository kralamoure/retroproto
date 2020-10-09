// generated (unrevised)
package msgcli

import (
	"github.com/kralamoure/d1proto"
)

type PartyFollowAll struct{}

func (m PartyFollowAll) ProtocolId() d1proto.MsgCliId {
	return d1proto.PartyFollowAll
}

func (m PartyFollowAll) Serialized() (string, error) {
	return "", d1proto.ErrNotImplemented
}

func (m *PartyFollowAll) Deserialize(extra string) error {
	return d1proto.ErrNotImplemented
}
