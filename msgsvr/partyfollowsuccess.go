// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyFollowSuccess struct{}

func NewPartyFollowSuccess(extra string) (PartyFollowSuccess, error) {
	var m PartyFollowSuccess

	if err := m.Deserialize(extra); err != nil {
		return PartyFollowSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyFollowSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyFollowSuccess
}

func (m PartyFollowSuccess) MessageName() string {
	return "PartyFollowSuccess"
}

func (m PartyFollowSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyFollowSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
