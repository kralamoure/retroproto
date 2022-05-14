// generated (unrevised)

package msgcli

import (
	"fmt"

	"github.com/kralamoure/retroproto"
)

type ConquestWorldInfosLave struct{}

func NewConquestWorldInfosLave(extra string) (ConquestWorldInfosLave, error) {
	var m ConquestWorldInfosLave

	if err := m.Deserialize(extra); err != nil {
		return ConquestWorldInfosLave{}, fmt.Errorf("could not deserialize message: %w", err)
	}

	return m, nil
}

func (m ConquestWorldInfosLave) MessageId() retroproto.MsgCliId {
	return retroproto.ConquestWorldInfosLave
}

func (m ConquestWorldInfosLave) MessageName() string {
	return "ConquestWorldInfosLave"
}

func (m ConquestWorldInfosLave) Serialized() (string, error) {
	return "", retroproto.ErrNotImplemented
}

func (m *ConquestWorldInfosLave) Deserialize(extra string) error {
	return retroproto.ErrNotImplemented
}
