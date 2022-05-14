// generated (unrevised)

package msgsvr

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type PartyCreateSuccess struct{}

func NewPartyCreateSuccess(extra string) (PartyCreateSuccess, error) {
	var m PartyCreateSuccess

	if err := m.Deserialize(extra); err != nil {
		return PartyCreateSuccess{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m PartyCreateSuccess) MessageId() retroproto.MsgSvrId {
	return retroproto.PartyCreateSuccess
}

func (m PartyCreateSuccess) MessageName() string {
	return "PartyCreateSuccess"
}

func (m PartyCreateSuccess) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *PartyCreateSuccess) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
