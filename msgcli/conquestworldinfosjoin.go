// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestWorldInfosJoin struct{}

func NewConquestWorldInfosJoin(extra string) (ConquestWorldInfosJoin, error) {
	var m ConquestWorldInfosJoin

	if err := m.Deserialize(extra); err != nil {
		return ConquestWorldInfosJoin{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestWorldInfosJoin) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestWorldInfosJoin
}

func (m ConquestWorldInfosJoin) MessageName() string {
	return "ConquestWorldInfosJoin"
}

func (m ConquestWorldInfosJoin) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestWorldInfosJoin) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
